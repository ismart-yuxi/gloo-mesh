// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_cert_signer is a generated GoMock package.
package mock_cert_signer

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1"
	types0 "github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1/types"
	cert_secrets "github.com/solo-io/service-mesh-hub/pkg/security/secrets"
)

// MockVirtualMeshCertClient is a mock of VirtualMeshCertClient interface.
type MockVirtualMeshCertClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshCertClientMockRecorder
}

// MockVirtualMeshCertClientMockRecorder is the mock recorder for MockVirtualMeshCertClient.
type MockVirtualMeshCertClientMockRecorder struct {
	mock *MockVirtualMeshCertClient
}

// NewMockVirtualMeshCertClient creates a new mock instance.
func NewMockVirtualMeshCertClient(ctrl *gomock.Controller) *MockVirtualMeshCertClient {
	mock := &MockVirtualMeshCertClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshCertClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshCertClient) EXPECT() *MockVirtualMeshCertClientMockRecorder {
	return m.recorder
}

// GetRootCaBundle mocks base method.
func (m *MockVirtualMeshCertClient) GetRootCaBundle(ctx context.Context, meshRef *types.ResourceRef) (*cert_secrets.RootCAData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRootCaBundle", ctx, meshRef)
	ret0, _ := ret[0].(*cert_secrets.RootCAData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRootCaBundle indicates an expected call of GetRootCaBundle.
func (mr *MockVirtualMeshCertClientMockRecorder) GetRootCaBundle(ctx, meshRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRootCaBundle", reflect.TypeOf((*MockVirtualMeshCertClient)(nil).GetRootCaBundle), ctx, meshRef)
}

// MockVirtualMeshCSRSigner is a mock of VirtualMeshCSRSigner interface.
type MockVirtualMeshCSRSigner struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshCSRSignerMockRecorder
}

// MockVirtualMeshCSRSignerMockRecorder is the mock recorder for MockVirtualMeshCSRSigner.
type MockVirtualMeshCSRSignerMockRecorder struct {
	mock *MockVirtualMeshCSRSigner
}

// NewMockVirtualMeshCSRSigner creates a new mock instance.
func NewMockVirtualMeshCSRSigner(ctrl *gomock.Controller) *MockVirtualMeshCSRSigner {
	mock := &MockVirtualMeshCSRSigner{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshCSRSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshCSRSigner) EXPECT() *MockVirtualMeshCSRSignerMockRecorder {
	return m.recorder
}

// Sign mocks base method.
func (m *MockVirtualMeshCSRSigner) Sign(ctx context.Context, obj *v1alpha1.VirtualMeshCertificateSigningRequest) *types0.VirtualMeshCertificateSigningRequestStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", ctx, obj)
	ret0, _ := ret[0].(*types0.VirtualMeshCertificateSigningRequestStatus)
	return ret0
}

// Sign indicates an expected call of Sign.
func (mr *MockVirtualMeshCSRSignerMockRecorder) Sign(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockVirtualMeshCSRSigner)(nil).Sign), ctx, obj)
}