// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1"
	controller "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockSettingsReconciler is a mock of SettingsReconciler interface.
type MockSettingsReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsReconcilerMockRecorder
}

// MockSettingsReconcilerMockRecorder is the mock recorder for MockSettingsReconciler.
type MockSettingsReconcilerMockRecorder struct {
	mock *MockSettingsReconciler
}

// NewMockSettingsReconciler creates a new mock instance.
func NewMockSettingsReconciler(ctrl *gomock.Controller) *MockSettingsReconciler {
	mock := &MockSettingsReconciler{ctrl: ctrl}
	mock.recorder = &MockSettingsReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingsReconciler) EXPECT() *MockSettingsReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSettings mocks base method.
func (m *MockSettingsReconciler) ReconcileSettings(obj *v1.Settings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettings", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSettings indicates an expected call of ReconcileSettings.
func (mr *MockSettingsReconcilerMockRecorder) ReconcileSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettings", reflect.TypeOf((*MockSettingsReconciler)(nil).ReconcileSettings), obj)
}

// MockSettingsDeletionReconciler is a mock of SettingsDeletionReconciler interface.
type MockSettingsDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsDeletionReconcilerMockRecorder
}

// MockSettingsDeletionReconcilerMockRecorder is the mock recorder for MockSettingsDeletionReconciler.
type MockSettingsDeletionReconcilerMockRecorder struct {
	mock *MockSettingsDeletionReconciler
}

// NewMockSettingsDeletionReconciler creates a new mock instance.
func NewMockSettingsDeletionReconciler(ctrl *gomock.Controller) *MockSettingsDeletionReconciler {
	mock := &MockSettingsDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockSettingsDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingsDeletionReconciler) EXPECT() *MockSettingsDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSettingsDeletion mocks base method.
func (m *MockSettingsDeletionReconciler) ReconcileSettingsDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettingsDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileSettingsDeletion indicates an expected call of ReconcileSettingsDeletion.
func (mr *MockSettingsDeletionReconcilerMockRecorder) ReconcileSettingsDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettingsDeletion", reflect.TypeOf((*MockSettingsDeletionReconciler)(nil).ReconcileSettingsDeletion), req)
}

// MockSettingsFinalizer is a mock of SettingsFinalizer interface.
type MockSettingsFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsFinalizerMockRecorder
}

// MockSettingsFinalizerMockRecorder is the mock recorder for MockSettingsFinalizer.
type MockSettingsFinalizerMockRecorder struct {
	mock *MockSettingsFinalizer
}

// NewMockSettingsFinalizer creates a new mock instance.
func NewMockSettingsFinalizer(ctrl *gomock.Controller) *MockSettingsFinalizer {
	mock := &MockSettingsFinalizer{ctrl: ctrl}
	mock.recorder = &MockSettingsFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingsFinalizer) EXPECT() *MockSettingsFinalizerMockRecorder {
	return m.recorder
}

// FinalizeSettings mocks base method.
func (m *MockSettingsFinalizer) FinalizeSettings(obj *v1.Settings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeSettings", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeSettings indicates an expected call of FinalizeSettings.
func (mr *MockSettingsFinalizerMockRecorder) FinalizeSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeSettings", reflect.TypeOf((*MockSettingsFinalizer)(nil).FinalizeSettings), obj)
}

// ReconcileSettings mocks base method.
func (m *MockSettingsFinalizer) ReconcileSettings(obj *v1.Settings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettings", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSettings indicates an expected call of ReconcileSettings.
func (mr *MockSettingsFinalizerMockRecorder) ReconcileSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettings", reflect.TypeOf((*MockSettingsFinalizer)(nil).ReconcileSettings), obj)
}

// SettingsFinalizerName mocks base method.
func (m *MockSettingsFinalizer) SettingsFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SettingsFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// SettingsFinalizerName indicates an expected call of SettingsFinalizerName.
func (mr *MockSettingsFinalizerMockRecorder) SettingsFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SettingsFinalizerName", reflect.TypeOf((*MockSettingsFinalizer)(nil).SettingsFinalizerName))
}

// MockSettingsReconcileLoop is a mock of SettingsReconcileLoop interface.
type MockSettingsReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsReconcileLoopMockRecorder
}

// MockSettingsReconcileLoopMockRecorder is the mock recorder for MockSettingsReconcileLoop.
type MockSettingsReconcileLoopMockRecorder struct {
	mock *MockSettingsReconcileLoop
}

// NewMockSettingsReconcileLoop creates a new mock instance.
func NewMockSettingsReconcileLoop(ctrl *gomock.Controller) *MockSettingsReconcileLoop {
	mock := &MockSettingsReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockSettingsReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingsReconcileLoop) EXPECT() *MockSettingsReconcileLoopMockRecorder {
	return m.recorder
}

// RunSettingsReconciler mocks base method.
func (m *MockSettingsReconcileLoop) RunSettingsReconciler(ctx context.Context, rec controller.SettingsReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunSettingsReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunSettingsReconciler indicates an expected call of RunSettingsReconciler.
func (mr *MockSettingsReconcileLoopMockRecorder) RunSettingsReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSettingsReconciler", reflect.TypeOf((*MockSettingsReconcileLoop)(nil).RunSettingsReconciler), varargs...)
}

// MockDashboardReconciler is a mock of DashboardReconciler interface.
type MockDashboardReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockDashboardReconcilerMockRecorder
}

// MockDashboardReconcilerMockRecorder is the mock recorder for MockDashboardReconciler.
type MockDashboardReconcilerMockRecorder struct {
	mock *MockDashboardReconciler
}

// NewMockDashboardReconciler creates a new mock instance.
func NewMockDashboardReconciler(ctrl *gomock.Controller) *MockDashboardReconciler {
	mock := &MockDashboardReconciler{ctrl: ctrl}
	mock.recorder = &MockDashboardReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDashboardReconciler) EXPECT() *MockDashboardReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDashboard mocks base method.
func (m *MockDashboardReconciler) ReconcileDashboard(obj *v1.Dashboard) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDashboard", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDashboard indicates an expected call of ReconcileDashboard.
func (mr *MockDashboardReconcilerMockRecorder) ReconcileDashboard(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDashboard", reflect.TypeOf((*MockDashboardReconciler)(nil).ReconcileDashboard), obj)
}

// MockDashboardDeletionReconciler is a mock of DashboardDeletionReconciler interface.
type MockDashboardDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockDashboardDeletionReconcilerMockRecorder
}

// MockDashboardDeletionReconcilerMockRecorder is the mock recorder for MockDashboardDeletionReconciler.
type MockDashboardDeletionReconcilerMockRecorder struct {
	mock *MockDashboardDeletionReconciler
}

// NewMockDashboardDeletionReconciler creates a new mock instance.
func NewMockDashboardDeletionReconciler(ctrl *gomock.Controller) *MockDashboardDeletionReconciler {
	mock := &MockDashboardDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockDashboardDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDashboardDeletionReconciler) EXPECT() *MockDashboardDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDashboardDeletion mocks base method.
func (m *MockDashboardDeletionReconciler) ReconcileDashboardDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDashboardDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileDashboardDeletion indicates an expected call of ReconcileDashboardDeletion.
func (mr *MockDashboardDeletionReconcilerMockRecorder) ReconcileDashboardDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDashboardDeletion", reflect.TypeOf((*MockDashboardDeletionReconciler)(nil).ReconcileDashboardDeletion), req)
}

// MockDashboardFinalizer is a mock of DashboardFinalizer interface.
type MockDashboardFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockDashboardFinalizerMockRecorder
}

// MockDashboardFinalizerMockRecorder is the mock recorder for MockDashboardFinalizer.
type MockDashboardFinalizerMockRecorder struct {
	mock *MockDashboardFinalizer
}

// NewMockDashboardFinalizer creates a new mock instance.
func NewMockDashboardFinalizer(ctrl *gomock.Controller) *MockDashboardFinalizer {
	mock := &MockDashboardFinalizer{ctrl: ctrl}
	mock.recorder = &MockDashboardFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDashboardFinalizer) EXPECT() *MockDashboardFinalizerMockRecorder {
	return m.recorder
}

// DashboardFinalizerName mocks base method.
func (m *MockDashboardFinalizer) DashboardFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DashboardFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DashboardFinalizerName indicates an expected call of DashboardFinalizerName.
func (mr *MockDashboardFinalizerMockRecorder) DashboardFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DashboardFinalizerName", reflect.TypeOf((*MockDashboardFinalizer)(nil).DashboardFinalizerName))
}

// FinalizeDashboard mocks base method.
func (m *MockDashboardFinalizer) FinalizeDashboard(obj *v1.Dashboard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeDashboard", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeDashboard indicates an expected call of FinalizeDashboard.
func (mr *MockDashboardFinalizerMockRecorder) FinalizeDashboard(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeDashboard", reflect.TypeOf((*MockDashboardFinalizer)(nil).FinalizeDashboard), obj)
}

// ReconcileDashboard mocks base method.
func (m *MockDashboardFinalizer) ReconcileDashboard(obj *v1.Dashboard) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDashboard", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDashboard indicates an expected call of ReconcileDashboard.
func (mr *MockDashboardFinalizerMockRecorder) ReconcileDashboard(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDashboard", reflect.TypeOf((*MockDashboardFinalizer)(nil).ReconcileDashboard), obj)
}

// MockDashboardReconcileLoop is a mock of DashboardReconcileLoop interface.
type MockDashboardReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockDashboardReconcileLoopMockRecorder
}

// MockDashboardReconcileLoopMockRecorder is the mock recorder for MockDashboardReconcileLoop.
type MockDashboardReconcileLoopMockRecorder struct {
	mock *MockDashboardReconcileLoop
}

// NewMockDashboardReconcileLoop creates a new mock instance.
func NewMockDashboardReconcileLoop(ctrl *gomock.Controller) *MockDashboardReconcileLoop {
	mock := &MockDashboardReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockDashboardReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDashboardReconcileLoop) EXPECT() *MockDashboardReconcileLoopMockRecorder {
	return m.recorder
}

// RunDashboardReconciler mocks base method.
func (m *MockDashboardReconcileLoop) RunDashboardReconciler(ctx context.Context, rec controller.DashboardReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunDashboardReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunDashboardReconciler indicates an expected call of RunDashboardReconciler.
func (mr *MockDashboardReconcileLoopMockRecorder) RunDashboardReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunDashboardReconciler", reflect.TypeOf((*MockDashboardReconcileLoop)(nil).RunDashboardReconciler), varargs...)
}
