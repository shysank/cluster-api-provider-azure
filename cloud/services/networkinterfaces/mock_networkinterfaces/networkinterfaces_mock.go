/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../networkinterfaces.go

// Package mock_networkinterfaces is a generated GoMock package.
package mock_networkinterfaces

import (
	autorest "github.com/Azure/go-autorest/autorest"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	v1alpha3 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
	v1alpha30 "sigs.k8s.io/cluster-api/api/v1alpha3"
)

// MockNICScope is a mock of NICScope interface.
type MockNICScope struct {
	ctrl     *gomock.Controller
	recorder *MockNICScopeMockRecorder
}

// MockNICScopeMockRecorder is the mock recorder for MockNICScope.
type MockNICScopeMockRecorder struct {
	mock *MockNICScope
}

// NewMockNICScope creates a new mock instance.
func NewMockNICScope(ctrl *gomock.Controller) *MockNICScope {
	mock := &MockNICScope{ctrl: ctrl}
	mock.recorder = &MockNICScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNICScope) EXPECT() *MockNICScopeMockRecorder {
	return m.recorder
}

// Info mocks base method.
func (m *MockNICScope) Info(msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockNICScopeMockRecorder) Info(msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockNICScope)(nil).Info), varargs...)
}

// Enabled mocks base method.
func (m *MockNICScope) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockNICScopeMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockNICScope)(nil).Enabled))
}

// Error mocks base method.
func (m *MockNICScope) Error(err error, msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{err, msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockNICScopeMockRecorder) Error(err, msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{err, msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockNICScope)(nil).Error), varargs...)
}

// V mocks base method.
func (m *MockNICScope) V(level int) logr.InfoLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V", level)
	ret0, _ := ret[0].(logr.InfoLogger)
	return ret0
}

// V indicates an expected call of V.
func (mr *MockNICScopeMockRecorder) V(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V", reflect.TypeOf((*MockNICScope)(nil).V), level)
}

// WithValues mocks base method.
func (m *MockNICScope) WithValues(keysAndValues ...interface{}) logr.Logger {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithValues", varargs...)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithValues indicates an expected call of WithValues.
func (mr *MockNICScopeMockRecorder) WithValues(keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithValues", reflect.TypeOf((*MockNICScope)(nil).WithValues), keysAndValues...)
}

// WithName mocks base method.
func (m *MockNICScope) WithName(name string) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithName", name)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithName indicates an expected call of WithName.
func (mr *MockNICScopeMockRecorder) WithName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithName", reflect.TypeOf((*MockNICScope)(nil).WithName), name)
}

// SubscriptionID mocks base method.
func (m *MockNICScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockNICScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockNICScope)(nil).SubscriptionID))
}

// ClientID mocks base method.
func (m *MockNICScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockNICScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockNICScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockNICScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockNICScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockNICScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockNICScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockNICScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockNICScope)(nil).CloudEnvironment))
}

// TenantID mocks base method.
func (m *MockNICScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockNICScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockNICScope)(nil).TenantID))
}

// BaseURI mocks base method.
func (m *MockNICScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockNICScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockNICScope)(nil).BaseURI))
}

// Authorizer mocks base method.
func (m *MockNICScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockNICScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockNICScope)(nil).Authorizer))
}

// ResourceGroup mocks base method.
func (m *MockNICScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockNICScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockNICScope)(nil).ResourceGroup))
}

// ClusterName mocks base method.
func (m *MockNICScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockNICScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockNICScope)(nil).ClusterName))
}

// Location mocks base method.
func (m *MockNICScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockNICScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockNICScope)(nil).Location))
}

// AdditionalTags mocks base method.
func (m *MockNICScope) AdditionalTags() v1alpha3.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1alpha3.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockNICScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockNICScope)(nil).AdditionalTags))
}

// FailureDomains mocks base method.
func (m *MockNICScope) FailureDomains() v1alpha30.FailureDomains {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailureDomains")
	ret0, _ := ret[0].(v1alpha30.FailureDomains)
	return ret0
}

// FailureDomains indicates an expected call of FailureDomains.
func (mr *MockNICScopeMockRecorder) FailureDomains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureDomains", reflect.TypeOf((*MockNICScope)(nil).FailureDomains))
}

// NICSpecs mocks base method.
func (m *MockNICScope) NICSpecs() []azure.NICSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NICSpecs")
	ret0, _ := ret[0].([]azure.NICSpec)
	return ret0
}

// NICSpecs indicates an expected call of NICSpecs.
func (mr *MockNICScopeMockRecorder) NICSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NICSpecs", reflect.TypeOf((*MockNICScope)(nil).NICSpecs))
}
