// Code generated by MockGen. DO NOT EDIT.
// Source: ./snapshot.go

// Package mock_input is a generated GoMock package.
package mock_input

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1alpha2sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2/sets"
	input "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	v1alpha2sets0 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2/sets"
	v1alpha2sets1 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2/sets"
	v1alpha1sets "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/sets"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockSnapshot is a mock of Snapshot interface
type MockSnapshot struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotMockRecorder
}

// MockSnapshotMockRecorder is the mock recorder for MockSnapshot
type MockSnapshotMockRecorder struct {
	mock *MockSnapshot
}

// NewMockSnapshot creates a new mock instance
func NewMockSnapshot(ctrl *gomock.Controller) *MockSnapshot {
	mock := &MockSnapshot{ctrl: ctrl}
	mock.recorder = &MockSnapshotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSnapshot) EXPECT() *MockSnapshotMockRecorder {
	return m.recorder
}

// Settings mocks base method
func (m *MockSnapshot) Settings() v1alpha2sets1.SettingsSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Settings")
	ret0, _ := ret[0].(v1alpha2sets1.SettingsSet)
	return ret0
}

// Settings indicates an expected call of Settings
func (mr *MockSnapshotMockRecorder) Settings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Settings", reflect.TypeOf((*MockSnapshot)(nil).Settings))
}

// TrafficTargets mocks base method
func (m *MockSnapshot) TrafficTargets() v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficTargets")
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// TrafficTargets indicates an expected call of TrafficTargets
func (mr *MockSnapshotMockRecorder) TrafficTargets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficTargets", reflect.TypeOf((*MockSnapshot)(nil).TrafficTargets))
}

// Workloads mocks base method
func (m *MockSnapshot) Workloads() v1alpha2sets.WorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Workloads")
	ret0, _ := ret[0].(v1alpha2sets.WorkloadSet)
	return ret0
}

// Workloads indicates an expected call of Workloads
func (mr *MockSnapshotMockRecorder) Workloads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Workloads", reflect.TypeOf((*MockSnapshot)(nil).Workloads))
}

// Meshes mocks base method
func (m *MockSnapshot) Meshes() v1alpha2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meshes")
	ret0, _ := ret[0].(v1alpha2sets.MeshSet)
	return ret0
}

// Meshes indicates an expected call of Meshes
func (mr *MockSnapshotMockRecorder) Meshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meshes", reflect.TypeOf((*MockSnapshot)(nil).Meshes))
}

// TrafficPolicies mocks base method
func (m *MockSnapshot) TrafficPolicies() v1alpha2sets0.TrafficPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficPolicies")
	ret0, _ := ret[0].(v1alpha2sets0.TrafficPolicySet)
	return ret0
}

// TrafficPolicies indicates an expected call of TrafficPolicies
func (mr *MockSnapshotMockRecorder) TrafficPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficPolicies", reflect.TypeOf((*MockSnapshot)(nil).TrafficPolicies))
}

// AccessPolicies mocks base method
func (m *MockSnapshot) AccessPolicies() v1alpha2sets0.AccessPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessPolicies")
	ret0, _ := ret[0].(v1alpha2sets0.AccessPolicySet)
	return ret0
}

// AccessPolicies indicates an expected call of AccessPolicies
func (mr *MockSnapshotMockRecorder) AccessPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessPolicies", reflect.TypeOf((*MockSnapshot)(nil).AccessPolicies))
}

// VirtualMeshes mocks base method
func (m *MockSnapshot) VirtualMeshes() v1alpha2sets0.VirtualMeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualMeshes")
	ret0, _ := ret[0].(v1alpha2sets0.VirtualMeshSet)
	return ret0
}

// VirtualMeshes indicates an expected call of VirtualMeshes
func (mr *MockSnapshotMockRecorder) VirtualMeshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualMeshes", reflect.TypeOf((*MockSnapshot)(nil).VirtualMeshes))
}

// FailoverServices mocks base method
func (m *MockSnapshot) FailoverServices() v1alpha2sets0.FailoverServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailoverServices")
	ret0, _ := ret[0].(v1alpha2sets0.FailoverServiceSet)
	return ret0
}

// FailoverServices indicates an expected call of FailoverServices
func (mr *MockSnapshotMockRecorder) FailoverServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailoverServices", reflect.TypeOf((*MockSnapshot)(nil).FailoverServices))
}

// Secrets mocks base method
func (m *MockSnapshot) Secrets() v1sets.SecretSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Secrets")
	ret0, _ := ret[0].(v1sets.SecretSet)
	return ret0
}

// Secrets indicates an expected call of Secrets
func (mr *MockSnapshotMockRecorder) Secrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Secrets", reflect.TypeOf((*MockSnapshot)(nil).Secrets))
}

// KubernetesClusters mocks base method
func (m *MockSnapshot) KubernetesClusters() v1alpha1sets.KubernetesClusterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubernetesClusters")
	ret0, _ := ret[0].(v1alpha1sets.KubernetesClusterSet)
	return ret0
}

// KubernetesClusters indicates an expected call of KubernetesClusters
func (mr *MockSnapshotMockRecorder) KubernetesClusters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubernetesClusters", reflect.TypeOf((*MockSnapshot)(nil).KubernetesClusters))
}

// SyncStatuses mocks base method
func (m *MockSnapshot) SyncStatuses(ctx context.Context, c client.Client, opts input.SyncStatusOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatuses", ctx, c, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatuses indicates an expected call of SyncStatuses
func (mr *MockSnapshotMockRecorder) SyncStatuses(ctx, c, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatuses", reflect.TypeOf((*MockSnapshot)(nil).SyncStatuses), ctx, c, opts)
}

// MarshalJSON mocks base method
func (m *MockSnapshot) MarshalJSON() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (mr *MockSnapshotMockRecorder) MarshalJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockSnapshot)(nil).MarshalJSON))
}

// MockBuilder is a mock of Builder interface
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// BuildSnapshot mocks base method
func (m *MockBuilder) BuildSnapshot(ctx context.Context, name string, opts input.BuildOptions) (input.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSnapshot", ctx, name, opts)
	ret0, _ := ret[0].(input.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSnapshot indicates an expected call of BuildSnapshot
func (mr *MockBuilderMockRecorder) BuildSnapshot(ctx, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildSnapshot), ctx, name, opts)
}
