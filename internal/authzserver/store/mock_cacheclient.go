// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/marmotedu/api/proto/apiserver/v1 (interfaces: CacheClient)

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/marmotedu/api/proto/apiserver/v1"
	grpc "google.golang.org/grpc"
)

// MockCacheClient is a mock of CacheClient interface.
type MockCacheClient struct {
	ctrl     *gomock.Controller
	recorder *MockCacheClientMockRecorder
}

// MockCacheClientMockRecorder is the mock recorder for MockCacheClient.
type MockCacheClientMockRecorder struct {
	mock *MockCacheClient
}

// NewMockCacheClient creates a new mock instance.
func NewMockCacheClient(ctrl *gomock.Controller) *MockCacheClient {
	mock := &MockCacheClient{ctrl: ctrl}
	mock.recorder = &MockCacheClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheClient) EXPECT() *MockCacheClientMockRecorder {
	return m.recorder
}

// ListPolicies mocks base method.
func (m *MockCacheClient) ListPolicies(arg0 context.Context, arg1 *v1.ListPoliciesRequest, arg2 ...grpc.CallOption) (*v1.ListPoliciesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPolicies", varargs...)
	ret0, _ := ret[0].(*v1.ListPoliciesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPolicies indicates an expected call of ListPolicies.
func (mr *MockCacheClientMockRecorder) ListPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicies", reflect.TypeOf((*MockCacheClient)(nil).ListPolicies), varargs...)
}

// ListSecrets mocks base method.
func (m *MockCacheClient) ListSecrets(arg0 context.Context, arg1 *v1.ListSecretsRequest, arg2 ...grpc.CallOption) (*v1.ListSecretsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSecrets", varargs...)
	ret0, _ := ret[0].(*v1.ListSecretsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockCacheClientMockRecorder) ListSecrets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockCacheClient)(nil).ListSecrets), varargs...)
}
