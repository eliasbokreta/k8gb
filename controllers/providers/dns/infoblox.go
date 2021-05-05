/*
Copyright 2021 The k8gb Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Generated by GoLic, for more details see: https://github.com/AbsaOSS/golic
*/
package dns

import (
	"fmt"
	"time"

	externaldns "sigs.k8s.io/external-dns/endpoint"

	"github.com/AbsaOSS/k8gb/controllers/providers/assistant"

	k8gbv1beta1 "github.com/AbsaOSS/k8gb/api/v1beta1"
	"github.com/AbsaOSS/k8gb/controllers/depresolver"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

type InfobloxProvider struct {
	assistant assistant.IAssistant
	config    depresolver.Config
}

func NewInfobloxDNS(config depresolver.Config, assistant assistant.IAssistant) *InfobloxProvider {
	return &InfobloxProvider{
		assistant: assistant,
		config:    config,
	}
}

func (p *InfobloxProvider) CreateZoneDelegationForExternalDNS(gslb *k8gbv1beta1.Gslb) error {
	objMgr, err := p.infobloxConnection()
	if err != nil {
		return err
	}
	addresses, err := p.assistant.GslbIngressExposedIPs(gslb)
	if err != nil {
		return err
	}
	var delegateTo []ibclient.NameServer

	for _, address := range addresses {
		nameServer := ibclient.NameServer{Address: address, Name: nsServerName(p.config)}
		delegateTo = append(delegateTo, nameServer)
	}

	findZone, err := objMgr.GetZoneDelegated(p.config.DNSZone)
	if err != nil {
		return err
	}

	if !p.config.SplitBrainCheck {
		log.Info().Msg("Split-brain handling is disabled")
	}

	if findZone != nil {
		err = p.checkZoneDelegated(findZone)
		if err != nil {
			return err
		}

		if len(findZone.Ref) > 0 {

			// Drop own records for straight away update
			existingDelegateTo := p.filterOutDelegateTo(findZone.DelegateTo, nsServerName(p.config))
			existingDelegateTo = append(existingDelegateTo, delegateTo...)

			// Drop external records if they are stale
			if p.config.SplitBrainCheck {
				for _, extCluster := range p.config.ExtClustersGeoTags {
					nsServerNameExt := getNSServerName(extCluster, p.config.DNSZone, p.config.EdgeDNSZone)
					err = p.assistant.InspectTXTThreshold(
						getExternalClusterHeartbeatFQDN(gslb, extCluster, p.config.EdgeDNSZone),
						p.config.Override.FakeDNSEnabled,
						time.Second*time.Duration(gslb.Spec.Strategy.SplitBrainThresholdSeconds))
					if err != nil {
						log.Err(err).Msgf("Got the error from TXT based checkAlive. External cluster (%s) doesn't "+
							"look alive, filtering it out from delegated zone configuration...", nsServerNameExt)
						existingDelegateTo = p.filterOutDelegateTo(existingDelegateTo, nsServerNameExt)
					}
				}
			}

			log.Info().Msgf("Updating delegated zone(%s) with the server list(%v)", p.config.DNSZone, existingDelegateTo)
			_, err = objMgr.UpdateZoneDelegated(findZone.Ref, existingDelegateTo)
			if err != nil {
				return err
			}
		}
	} else {
		log.Info().Msgf("Creating delegated zone(%s)...", p.config.DNSZone)
		_, err = objMgr.CreateZoneDelegated(p.config.DNSZone, delegateTo)
		if err != nil {
			return err
		}
	}
	if p.config.SplitBrainCheck {
		return p.saveHeartbeatTXTRecord(objMgr, gslb)
	}
	return nil
}

func (p *InfobloxProvider) Finalize(gslb *k8gbv1beta1.Gslb) error {
	objMgr, err := p.infobloxConnection()
	if err != nil {
		return err
	}
	findZone, err := objMgr.GetZoneDelegated(p.config.DNSZone)
	if err != nil {
		return err
	}

	if findZone != nil {
		err = p.checkZoneDelegated(findZone)
		if err != nil {
			return err
		}
		if len(findZone.Ref) > 0 {
			log.Info().Msgf("Deleting delegated zone(%s)...", p.config.DNSZone)
			_, err := objMgr.DeleteZoneDelegated(findZone.Ref)
			if err != nil {
				return err
			}
		}
	}

	heartbeatTXTName := fmt.Sprintf("%s-heartbeat-%s.%s", gslb.Name, p.config.ClusterGeoTag, p.config.EdgeDNSZone)
	findTXT, err := objMgr.GetTXTRecord(heartbeatTXTName)
	if err != nil {
		return err
	}

	if findTXT != nil {
		if len(findTXT.Ref) > 0 {
			log.Info().Msgf("Deleting split brain TXT record(%s)...", heartbeatTXTName)
			_, err := objMgr.DeleteTXTRecord(findTXT.Ref)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *InfobloxProvider) GetExternalTargets(host string) (targets []string) {
	return p.assistant.GetExternalTargets(host, p.config.Override.FakeDNSEnabled, nsServerNameExt(p.config))
}

func (p *InfobloxProvider) GslbIngressExposedIPs(gslb *k8gbv1beta1.Gslb) ([]string, error) {
	return p.assistant.GslbIngressExposedIPs(gslb)
}

func (p *InfobloxProvider) SaveDNSEndpoint(gslb *k8gbv1beta1.Gslb, i *externaldns.DNSEndpoint) error {
	return p.assistant.SaveDNSEndpoint(gslb.Namespace, i)
}

func (p *InfobloxProvider) String() string {
	return "Infoblox"
}

func (p *InfobloxProvider) saveHeartbeatTXTRecord(objMgr *ibclient.ObjectManager, gslb *k8gbv1beta1.Gslb) (err error) {
	var heartbeatTXTRecord *ibclient.RecordTXT
	edgeTimestamp := fmt.Sprint(time.Now().UTC().Format("2006-01-02T15:04:05"))
	heartbeatTXTName := getExternalClusterHeartbeatFQDN(gslb, p.config.ClusterGeoTag, p.config.EdgeDNSZone)
	heartbeatTXTRecord, err = objMgr.GetTXTRecord(heartbeatTXTName)
	if err != nil {
		return
	}
	if heartbeatTXTRecord == nil {
		log.Info().Str("HeartbeatTXTName", heartbeatTXTName).Msg("Creating split brain TXT record")
		_, err = objMgr.CreateTXTRecord(heartbeatTXTName, edgeTimestamp, gslb.Spec.Strategy.DNSTtlSeconds, "default")
		if err != nil {
			return
		}
	} else {
		log.Info().Str("HeartbeatTXTName", heartbeatTXTName).Msg("Updating split brain TXT record")
		_, err = objMgr.UpdateTXTRecord(heartbeatTXTName, edgeTimestamp)
		if err != nil {
			return
		}
	}
	return
}
