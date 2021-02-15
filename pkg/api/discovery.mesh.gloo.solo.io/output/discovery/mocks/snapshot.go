// Code generated by MockGen. DO NOT EDIT.
// Source: ./snapshot.go

// Package mock_discovery is a generated GoMock package.
package mock_discovery

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	discovery "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/output/discovery"
	v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	v1alpha2sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2/sets"
	output "github.com/solo-io/skv2/contrib/pkg/output"
	multicluster "github.com/solo-io/skv2/pkg/multicluster"
	resource "github.com/solo-io/skv2/pkg/resource"
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

// TrafficTargets mocks base method
func (m *MockSnapshot) TrafficTargets() []discovery.LabeledTrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficTargets")
	ret0, _ := ret[0].([]discovery.LabeledTrafficTargetSet)
	return ret0
}

// TrafficTargets indicates an expected call of TrafficTargets
func (mr *MockSnapshotMockRecorder) TrafficTargets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficTargets", reflect.TypeOf((*MockSnapshot)(nil).TrafficTargets))
}

// Workloads mocks base method
func (m *MockSnapshot) Workloads() []discovery.LabeledWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Workloads")
	ret0, _ := ret[0].([]discovery.LabeledWorkloadSet)
	return ret0
}

// Workloads indicates an expected call of Workloads
func (mr *MockSnapshotMockRecorder) Workloads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Workloads", reflect.TypeOf((*MockSnapshot)(nil).Workloads))
}

// Meshes mocks base method
func (m *MockSnapshot) Meshes() []discovery.LabeledMeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meshes")
	ret0, _ := ret[0].([]discovery.LabeledMeshSet)
	return ret0
}

// Meshes indicates an expected call of Meshes
func (mr *MockSnapshotMockRecorder) Meshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meshes", reflect.TypeOf((*MockSnapshot)(nil).Meshes))
}

// ApplyLocalCluster mocks base method
func (m *MockSnapshot) ApplyLocalCluster(ctx context.Context, clusterClient client.Client, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyLocalCluster", ctx, clusterClient, errHandler)
}

// ApplyLocalCluster indicates an expected call of ApplyLocalCluster
func (mr *MockSnapshotMockRecorder) ApplyLocalCluster(ctx, clusterClient, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyLocalCluster", reflect.TypeOf((*MockSnapshot)(nil).ApplyLocalCluster), ctx, clusterClient, errHandler)
}

// ApplyMultiCluster mocks base method
func (m *MockSnapshot) ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyMultiCluster", ctx, multiClusterClient, errHandler)
}

// ApplyMultiCluster indicates an expected call of ApplyMultiCluster
func (mr *MockSnapshotMockRecorder) ApplyMultiCluster(ctx, multiClusterClient, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyMultiCluster", reflect.TypeOf((*MockSnapshot)(nil).ApplyMultiCluster), ctx, multiClusterClient, errHandler)
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

// MockLabeledTrafficTargetSet is a mock of LabeledTrafficTargetSet interface
type MockLabeledTrafficTargetSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledTrafficTargetSetMockRecorder
}

// MockLabeledTrafficTargetSetMockRecorder is the mock recorder for MockLabeledTrafficTargetSet
type MockLabeledTrafficTargetSetMockRecorder struct {
	mock *MockLabeledTrafficTargetSet
}

// NewMockLabeledTrafficTargetSet creates a new mock instance
func NewMockLabeledTrafficTargetSet(ctrl *gomock.Controller) *MockLabeledTrafficTargetSet {
	mock := &MockLabeledTrafficTargetSet{ctrl: ctrl}
	mock.recorder = &MockLabeledTrafficTargetSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLabeledTrafficTargetSet) EXPECT() *MockLabeledTrafficTargetSetMockRecorder {
	return m.recorder
}

// Labels mocks base method
func (m *MockLabeledTrafficTargetSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockLabeledTrafficTargetSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledTrafficTargetSet)(nil).Labels))
}

// Set mocks base method
func (m *MockLabeledTrafficTargetSet) Set() v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLabeledTrafficTargetSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledTrafficTargetSet)(nil).Set))
}

// Generic mocks base method
func (m *MockLabeledTrafficTargetSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockLabeledTrafficTargetSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledTrafficTargetSet)(nil).Generic))
}

// MockLabeledWorkloadSet is a mock of LabeledWorkloadSet interface
type MockLabeledWorkloadSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledWorkloadSetMockRecorder
}

// MockLabeledWorkloadSetMockRecorder is the mock recorder for MockLabeledWorkloadSet
type MockLabeledWorkloadSetMockRecorder struct {
	mock *MockLabeledWorkloadSet
}

// NewMockLabeledWorkloadSet creates a new mock instance
func NewMockLabeledWorkloadSet(ctrl *gomock.Controller) *MockLabeledWorkloadSet {
	mock := &MockLabeledWorkloadSet{ctrl: ctrl}
	mock.recorder = &MockLabeledWorkloadSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLabeledWorkloadSet) EXPECT() *MockLabeledWorkloadSetMockRecorder {
	return m.recorder
}

// Labels mocks base method
func (m *MockLabeledWorkloadSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockLabeledWorkloadSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledWorkloadSet)(nil).Labels))
}

// Set mocks base method
func (m *MockLabeledWorkloadSet) Set() v1alpha2sets.WorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1alpha2sets.WorkloadSet)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLabeledWorkloadSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledWorkloadSet)(nil).Set))
}

// Generic mocks base method
func (m *MockLabeledWorkloadSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockLabeledWorkloadSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledWorkloadSet)(nil).Generic))
}

// MockLabeledMeshSet is a mock of LabeledMeshSet interface
type MockLabeledMeshSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledMeshSetMockRecorder
}

// MockLabeledMeshSetMockRecorder is the mock recorder for MockLabeledMeshSet
type MockLabeledMeshSetMockRecorder struct {
	mock *MockLabeledMeshSet
}

// NewMockLabeledMeshSet creates a new mock instance
func NewMockLabeledMeshSet(ctrl *gomock.Controller) *MockLabeledMeshSet {
	mock := &MockLabeledMeshSet{ctrl: ctrl}
	mock.recorder = &MockLabeledMeshSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLabeledMeshSet) EXPECT() *MockLabeledMeshSetMockRecorder {
	return m.recorder
}

// Labels mocks base method
func (m *MockLabeledMeshSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockLabeledMeshSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledMeshSet)(nil).Labels))
}

// Set mocks base method
func (m *MockLabeledMeshSet) Set() v1alpha2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1alpha2sets.MeshSet)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLabeledMeshSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledMeshSet)(nil).Set))
}

// Generic mocks base method
func (m *MockLabeledMeshSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockLabeledMeshSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledMeshSet)(nil).Generic))
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

// AddTrafficTargets mocks base method
func (m *MockBuilder) AddTrafficTargets(trafficTargets ...*v1alpha2.TrafficTarget) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range trafficTargets {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddTrafficTargets", varargs...)
}

// AddTrafficTargets indicates an expected call of AddTrafficTargets
func (mr *MockBuilderMockRecorder) AddTrafficTargets(trafficTargets ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrafficTargets", reflect.TypeOf((*MockBuilder)(nil).AddTrafficTargets), trafficTargets...)
}

// GetTrafficTargets mocks base method
func (m *MockBuilder) GetTrafficTargets() v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrafficTargets")
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// GetTrafficTargets indicates an expected call of GetTrafficTargets
func (mr *MockBuilderMockRecorder) GetTrafficTargets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrafficTargets", reflect.TypeOf((*MockBuilder)(nil).GetTrafficTargets))
}

// AddWorkloads mocks base method
func (m *MockBuilder) AddWorkloads(workloads ...*v1alpha2.Workload) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range workloads {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddWorkloads", varargs...)
}

// AddWorkloads indicates an expected call of AddWorkloads
func (mr *MockBuilderMockRecorder) AddWorkloads(workloads ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkloads", reflect.TypeOf((*MockBuilder)(nil).AddWorkloads), workloads...)
}

// GetWorkloads mocks base method
func (m *MockBuilder) GetWorkloads() v1alpha2sets.WorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkloads")
	ret0, _ := ret[0].(v1alpha2sets.WorkloadSet)
	return ret0
}

// GetWorkloads indicates an expected call of GetWorkloads
func (mr *MockBuilderMockRecorder) GetWorkloads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkloads", reflect.TypeOf((*MockBuilder)(nil).GetWorkloads))
}

// AddMeshes mocks base method
func (m *MockBuilder) AddMeshes(meshes ...*v1alpha2.Mesh) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range meshes {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMeshes", varargs...)
}

// AddMeshes indicates an expected call of AddMeshes
func (mr *MockBuilderMockRecorder) AddMeshes(meshes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMeshes", reflect.TypeOf((*MockBuilder)(nil).AddMeshes), meshes...)
}

// GetMeshes mocks base method
func (m *MockBuilder) GetMeshes() v1alpha2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshes")
	ret0, _ := ret[0].(v1alpha2sets.MeshSet)
	return ret0
}

// GetMeshes indicates an expected call of GetMeshes
func (mr *MockBuilderMockRecorder) GetMeshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshes", reflect.TypeOf((*MockBuilder)(nil).GetMeshes))
}

// BuildLabelPartitionedSnapshot mocks base method
func (m *MockBuilder) BuildLabelPartitionedSnapshot(labelKey string) (discovery.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildLabelPartitionedSnapshot", labelKey)
	ret0, _ := ret[0].(discovery.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildLabelPartitionedSnapshot indicates an expected call of BuildLabelPartitionedSnapshot
func (mr *MockBuilderMockRecorder) BuildLabelPartitionedSnapshot(labelKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildLabelPartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildLabelPartitionedSnapshot), labelKey)
}

// BuildSinglePartitionedSnapshot mocks base method
func (m *MockBuilder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (discovery.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSinglePartitionedSnapshot", snapshotLabels)
	ret0, _ := ret[0].(discovery.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSinglePartitionedSnapshot indicates an expected call of BuildSinglePartitionedSnapshot
func (mr *MockBuilderMockRecorder) BuildSinglePartitionedSnapshot(snapshotLabels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSinglePartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildSinglePartitionedSnapshot), snapshotLabels)
}

// AddCluster mocks base method
func (m *MockBuilder) AddCluster(cluster string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddCluster", cluster)
}

// AddCluster indicates an expected call of AddCluster
func (mr *MockBuilderMockRecorder) AddCluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCluster", reflect.TypeOf((*MockBuilder)(nil).AddCluster), cluster)
}

// Clusters mocks base method
func (m *MockBuilder) Clusters() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clusters")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Clusters indicates an expected call of Clusters
func (mr *MockBuilderMockRecorder) Clusters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clusters", reflect.TypeOf((*MockBuilder)(nil).Clusters))
}

// Merge mocks base method
func (m *MockBuilder) Merge(other discovery.Builder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Merge", other)
}

// Merge indicates an expected call of Merge
func (mr *MockBuilderMockRecorder) Merge(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merge", reflect.TypeOf((*MockBuilder)(nil).Merge), other)
}

// Clone mocks base method
func (m *MockBuilder) Clone() discovery.Builder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(discovery.Builder)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockBuilderMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockBuilder)(nil).Clone))
}

// Delta mocks base method
func (m *MockBuilder) Delta(newSnap discovery.Builder) output.SnapshotDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSnap)
	ret0, _ := ret[0].(output.SnapshotDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockBuilderMockRecorder) Delta(newSnap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockBuilder)(nil).Delta), newSnap)
}

// Generic mocks base method
func (m *MockBuilder) Generic() resource.ClusterSnapshot {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(resource.ClusterSnapshot)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockBuilderMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockBuilder)(nil).Generic))
}
