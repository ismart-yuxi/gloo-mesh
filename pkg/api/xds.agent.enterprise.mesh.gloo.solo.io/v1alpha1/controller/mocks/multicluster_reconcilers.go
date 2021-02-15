// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1alpha1"
	controller "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	reflect "reflect"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterXdsConfigReconciler is a mock of MulticlusterXdsConfigReconciler interface
type MockMulticlusterXdsConfigReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterXdsConfigReconcilerMockRecorder
}

// MockMulticlusterXdsConfigReconcilerMockRecorder is the mock recorder for MockMulticlusterXdsConfigReconciler
type MockMulticlusterXdsConfigReconcilerMockRecorder struct {
	mock *MockMulticlusterXdsConfigReconciler
}

// NewMockMulticlusterXdsConfigReconciler creates a new mock instance
func NewMockMulticlusterXdsConfigReconciler(ctrl *gomock.Controller) *MockMulticlusterXdsConfigReconciler {
	mock := &MockMulticlusterXdsConfigReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterXdsConfigReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterXdsConfigReconciler) EXPECT() *MockMulticlusterXdsConfigReconcilerMockRecorder {
	return m.recorder
}

// ReconcileXdsConfig mocks base method
func (m *MockMulticlusterXdsConfigReconciler) ReconcileXdsConfig(clusterName string, obj *v1alpha1.XdsConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileXdsConfig", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileXdsConfig indicates an expected call of ReconcileXdsConfig
func (mr *MockMulticlusterXdsConfigReconcilerMockRecorder) ReconcileXdsConfig(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileXdsConfig", reflect.TypeOf((*MockMulticlusterXdsConfigReconciler)(nil).ReconcileXdsConfig), clusterName, obj)
}

// MockMulticlusterXdsConfigDeletionReconciler is a mock of MulticlusterXdsConfigDeletionReconciler interface
type MockMulticlusterXdsConfigDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterXdsConfigDeletionReconcilerMockRecorder
}

// MockMulticlusterXdsConfigDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterXdsConfigDeletionReconciler
type MockMulticlusterXdsConfigDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterXdsConfigDeletionReconciler
}

// NewMockMulticlusterXdsConfigDeletionReconciler creates a new mock instance
func NewMockMulticlusterXdsConfigDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterXdsConfigDeletionReconciler {
	mock := &MockMulticlusterXdsConfigDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterXdsConfigDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterXdsConfigDeletionReconciler) EXPECT() *MockMulticlusterXdsConfigDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileXdsConfigDeletion mocks base method
func (m *MockMulticlusterXdsConfigDeletionReconciler) ReconcileXdsConfigDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileXdsConfigDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileXdsConfigDeletion indicates an expected call of ReconcileXdsConfigDeletion
func (mr *MockMulticlusterXdsConfigDeletionReconcilerMockRecorder) ReconcileXdsConfigDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileXdsConfigDeletion", reflect.TypeOf((*MockMulticlusterXdsConfigDeletionReconciler)(nil).ReconcileXdsConfigDeletion), clusterName, req)
}

// MockMulticlusterXdsConfigReconcileLoop is a mock of MulticlusterXdsConfigReconcileLoop interface
type MockMulticlusterXdsConfigReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterXdsConfigReconcileLoopMockRecorder
}

// MockMulticlusterXdsConfigReconcileLoopMockRecorder is the mock recorder for MockMulticlusterXdsConfigReconcileLoop
type MockMulticlusterXdsConfigReconcileLoopMockRecorder struct {
	mock *MockMulticlusterXdsConfigReconcileLoop
}

// NewMockMulticlusterXdsConfigReconcileLoop creates a new mock instance
func NewMockMulticlusterXdsConfigReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterXdsConfigReconcileLoop {
	mock := &MockMulticlusterXdsConfigReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterXdsConfigReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterXdsConfigReconcileLoop) EXPECT() *MockMulticlusterXdsConfigReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterXdsConfigReconciler mocks base method
func (m *MockMulticlusterXdsConfigReconcileLoop) AddMulticlusterXdsConfigReconciler(ctx context.Context, rec controller.MulticlusterXdsConfigReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterXdsConfigReconciler", varargs...)
}

// AddMulticlusterXdsConfigReconciler indicates an expected call of AddMulticlusterXdsConfigReconciler
func (mr *MockMulticlusterXdsConfigReconcileLoopMockRecorder) AddMulticlusterXdsConfigReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterXdsConfigReconciler", reflect.TypeOf((*MockMulticlusterXdsConfigReconcileLoop)(nil).AddMulticlusterXdsConfigReconciler), varargs...)
}
