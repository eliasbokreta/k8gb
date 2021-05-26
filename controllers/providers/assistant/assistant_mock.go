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
// Code generated by MockGen. DO NOT EDIT.
// Source: controllers/providers/assistant/assistant.go

// Package assistant is a generated GoMock package.
package assistant

import (
	reflect "reflect"
	time "time"

	v1beta1 "github.com/AbsaOSS/k8gb/api/v1beta1"
	gomock "github.com/golang/mock/gomock"
	endpoint "sigs.k8s.io/external-dns/endpoint"
)

// MockAssistant is a mock of Assistant interface.
type MockAssistant struct {
	ctrl     *gomock.Controller
	recorder *MockAssistantMockRecorder
}

// MockAssistantMockRecorder is the mock recorder for MockAssistant.
type MockAssistantMockRecorder struct {
	mock *MockAssistant
}

// NewMockAssistant creates a new mock instance.
func NewMockAssistant(ctrl *gomock.Controller) *MockAssistant {
	mock := &MockAssistant{ctrl: ctrl}
	mock.recorder = &MockAssistantMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssistant) EXPECT() *MockAssistantMockRecorder {
	return m.recorder
}

// CoreDNSExposedIPs mocks base method.
func (m *MockAssistant) CoreDNSExposedIPs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoreDNSExposedIPs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CoreDNSExposedIPs indicates an expected call of CoreDNSExposedIPs.
func (mr *MockAssistantMockRecorder) CoreDNSExposedIPs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoreDNSExposedIPs", reflect.TypeOf((*MockAssistant)(nil).CoreDNSExposedIPs))
}

// GetExternalTargets mocks base method.
func (m *MockAssistant) GetExternalTargets(host string, fakeDNSEnabled bool, extClusterNsNames map[string]string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalTargets", host, fakeDNSEnabled, extClusterNsNames)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetExternalTargets indicates an expected call of GetExternalTargets.
func (mr *MockAssistantMockRecorder) GetExternalTargets(host, fakeDNSEnabled, extClusterNsNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalTargets", reflect.TypeOf((*MockAssistant)(nil).GetExternalTargets), host, fakeDNSEnabled, extClusterNsNames)
}

// GslbIngressExposedIPs mocks base method.
func (m *MockAssistant) GslbIngressExposedIPs(gslb *v1beta1.Gslb) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GslbIngressExposedIPs", gslb)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GslbIngressExposedIPs indicates an expected call of GslbIngressExposedIPs.
func (mr *MockAssistantMockRecorder) GslbIngressExposedIPs(gslb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GslbIngressExposedIPs", reflect.TypeOf((*MockAssistant)(nil).GslbIngressExposedIPs), gslb)
}

// InspectTXTThreshold mocks base method.
func (m *MockAssistant) InspectTXTThreshold(fqdn string, fakeDNSEnabled bool, splitBrainThreshold time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectTXTThreshold", fqdn, fakeDNSEnabled, splitBrainThreshold)
	ret0, _ := ret[0].(error)
	return ret0
}

// InspectTXTThreshold indicates an expected call of InspectTXTThreshold.
func (mr *MockAssistantMockRecorder) InspectTXTThreshold(fqdn, fakeDNSEnabled, splitBrainThreshold interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectTXTThreshold", reflect.TypeOf((*MockAssistant)(nil).InspectTXTThreshold), fqdn, fakeDNSEnabled, splitBrainThreshold)
}

// RemoveEndpoint mocks base method.
func (m *MockAssistant) RemoveEndpoint(endpointName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveEndpoint", endpointName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEndpoint indicates an expected call of RemoveEndpoint.
func (mr *MockAssistantMockRecorder) RemoveEndpoint(endpointName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEndpoint", reflect.TypeOf((*MockAssistant)(nil).RemoveEndpoint), endpointName)
}

// SaveDNSEndpoint mocks base method.
func (m *MockAssistant) SaveDNSEndpoint(namespace string, i *endpoint.DNSEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDNSEndpoint", namespace, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveDNSEndpoint indicates an expected call of SaveDNSEndpoint.
func (mr *MockAssistantMockRecorder) SaveDNSEndpoint(namespace, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDNSEndpoint", reflect.TypeOf((*MockAssistant)(nil).SaveDNSEndpoint), namespace, i)
}
