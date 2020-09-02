// Code generated by MockGen. DO NOT EDIT.
// Source: ./appmesh_output_syncer.go

// Package mock_appmesh is a generated GoMock package.
package mock_appmesh

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	appmesh "github.com/solo-io/service-mesh-hub/pkg/api/external/appmesh"
	output "github.com/solo-io/skv2/contrib/pkg/output"
)

// MockOutputSyncer is a mock of OutputSyncer interface
type MockOutputSyncer struct {
	ctrl     *gomock.Controller
	recorder *MockOutputSyncerMockRecorder
}

// MockOutputSyncerMockRecorder is the mock recorder for MockOutputSyncer
type MockOutputSyncerMockRecorder struct {
	mock *MockOutputSyncer
}

// NewMockOutputSyncer creates a new mock instance
func NewMockOutputSyncer(ctrl *gomock.Controller) *MockOutputSyncer {
	mock := &MockOutputSyncer{ctrl: ctrl}
	mock.recorder = &MockOutputSyncerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOutputSyncer) EXPECT() *MockOutputSyncerMockRecorder {
	return m.recorder
}

// Apply mocks base method
func (m *MockOutputSyncer) Apply(ctx context.Context, outputs appmesh.Snapshot, errHandler output.ErrorHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", ctx, outputs, errHandler)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply
func (mr *MockOutputSyncerMockRecorder) Apply(ctx, outputs, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockOutputSyncer)(nil).Apply), ctx, outputs, errHandler)
}
