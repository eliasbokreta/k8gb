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
package test

import (
	"k8gbterratest/utils"

	"github.com/AbsaOSS/gopkg/env"
)

var (
	settings utils.TestSettings
)

func init() {
	p1, _ := env.GetEnvAsIntOrFallback("DNS_SERVER1_PORT", 5053)
	p2, _ := env.GetEnvAsIntOrFallback("DNS_SERVER2_PORT", 5054)
	p3, _ := env.GetEnvAsIntOrFallback("DNS_SERVER3_PORT", 5055)
	settings = utils.TestSettings{
		DNSZone:         env.GetEnvAsStringOrFallback("GSLB_DOMAIN", "cloud.example.com"),
		PrimaryGeoTag:   env.GetEnvAsStringOrFallback("PRIMARY_GEO_TAG", "eu"),
		SecondaryGeoTag: env.GetEnvAsStringOrFallback("SECONDARY_GEO_TAG", "us"),
		DNSServer1:      env.GetEnvAsStringOrFallback("DNS_SERVER1", "localhost"),
		Port1:           p1,
		DNSServer2:      env.GetEnvAsStringOrFallback("DNS_SERVER2", "localhost"),
		Port2:           p2,
		DNSServer3:      env.GetEnvAsStringOrFallback("DNS_SERVER3", "localhost"),
		Port3:           p3,
		Cluster1:        env.GetEnvAsStringOrFallback("K8GB_CLUSTER1", "k3d-test-gslb1"),
		Cluster2:        env.GetEnvAsStringOrFallback("K8GB_CLUSTER2", "k3d-test-gslb2"),
		Cluster3:        env.GetEnvAsStringOrFallback("K8GB_CLUSTER3", "k3d-test-gslb3"),
		PodinfoImage:    env.GetEnvAsStringOrFallback("PODINFO_IMAGE_REPO", "ghcr.io/stefanprodan/podinfo"),
	}
}
