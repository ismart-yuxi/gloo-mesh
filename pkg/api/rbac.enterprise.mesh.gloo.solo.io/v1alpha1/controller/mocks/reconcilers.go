// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/rbac.enterprise.mesh.gloo.solo.io/v1alpha1"
	controller "github.com/solo-io/gloo-mesh/pkg/api/rbac.enterprise.mesh.gloo.solo.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	reflect "reflect"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockRoleReconciler is a mock of RoleReconciler interface
type MockRoleReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockRoleReconcilerMockRecorder
}

// MockRoleReconcilerMockRecorder is the mock recorder for MockRoleReconciler
type MockRoleReconcilerMockRecorder struct {
	mock *MockRoleReconciler
}

// NewMockRoleReconciler creates a new mock instance
func NewMockRoleReconciler(ctrl *gomock.Controller) *MockRoleReconciler {
	mock := &MockRoleReconciler{ctrl: ctrl}
	mock.recorder = &MockRoleReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleReconciler) EXPECT() *MockRoleReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRole mocks base method
func (m *MockRoleReconciler) ReconcileRole(obj *v1alpha1.Role) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRole", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRole indicates an expected call of ReconcileRole
func (mr *MockRoleReconcilerMockRecorder) ReconcileRole(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRole", reflect.TypeOf((*MockRoleReconciler)(nil).ReconcileRole), obj)
}

// MockRoleDeletionReconciler is a mock of RoleDeletionReconciler interface
type MockRoleDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockRoleDeletionReconcilerMockRecorder
}

// MockRoleDeletionReconcilerMockRecorder is the mock recorder for MockRoleDeletionReconciler
type MockRoleDeletionReconcilerMockRecorder struct {
	mock *MockRoleDeletionReconciler
}

// NewMockRoleDeletionReconciler creates a new mock instance
func NewMockRoleDeletionReconciler(ctrl *gomock.Controller) *MockRoleDeletionReconciler {
	mock := &MockRoleDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockRoleDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleDeletionReconciler) EXPECT() *MockRoleDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRoleDeletion mocks base method
func (m *MockRoleDeletionReconciler) ReconcileRoleDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRoleDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRoleDeletion indicates an expected call of ReconcileRoleDeletion
func (mr *MockRoleDeletionReconcilerMockRecorder) ReconcileRoleDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRoleDeletion", reflect.TypeOf((*MockRoleDeletionReconciler)(nil).ReconcileRoleDeletion), req)
}

// MockRoleFinalizer is a mock of RoleFinalizer interface
type MockRoleFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockRoleFinalizerMockRecorder
}

// MockRoleFinalizerMockRecorder is the mock recorder for MockRoleFinalizer
type MockRoleFinalizerMockRecorder struct {
	mock *MockRoleFinalizer
}

// NewMockRoleFinalizer creates a new mock instance
func NewMockRoleFinalizer(ctrl *gomock.Controller) *MockRoleFinalizer {
	mock := &MockRoleFinalizer{ctrl: ctrl}
	mock.recorder = &MockRoleFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleFinalizer) EXPECT() *MockRoleFinalizerMockRecorder {
	return m.recorder
}

// ReconcileRole mocks base method
func (m *MockRoleFinalizer) ReconcileRole(obj *v1alpha1.Role) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRole", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRole indicates an expected call of ReconcileRole
func (mr *MockRoleFinalizerMockRecorder) ReconcileRole(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRole", reflect.TypeOf((*MockRoleFinalizer)(nil).ReconcileRole), obj)
}

// RoleFinalizerName mocks base method
func (m *MockRoleFinalizer) RoleFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// RoleFinalizerName indicates an expected call of RoleFinalizerName
func (mr *MockRoleFinalizerMockRecorder) RoleFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleFinalizerName", reflect.TypeOf((*MockRoleFinalizer)(nil).RoleFinalizerName))
}

// FinalizeRole mocks base method
func (m *MockRoleFinalizer) FinalizeRole(obj *v1alpha1.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeRole", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeRole indicates an expected call of FinalizeRole
func (mr *MockRoleFinalizerMockRecorder) FinalizeRole(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeRole", reflect.TypeOf((*MockRoleFinalizer)(nil).FinalizeRole), obj)
}

// MockRoleReconcileLoop is a mock of RoleReconcileLoop interface
type MockRoleReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockRoleReconcileLoopMockRecorder
}

// MockRoleReconcileLoopMockRecorder is the mock recorder for MockRoleReconcileLoop
type MockRoleReconcileLoopMockRecorder struct {
	mock *MockRoleReconcileLoop
}

// NewMockRoleReconcileLoop creates a new mock instance
func NewMockRoleReconcileLoop(ctrl *gomock.Controller) *MockRoleReconcileLoop {
	mock := &MockRoleReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockRoleReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleReconcileLoop) EXPECT() *MockRoleReconcileLoopMockRecorder {
	return m.recorder
}

// RunRoleReconciler mocks base method
func (m *MockRoleReconcileLoop) RunRoleReconciler(ctx context.Context, rec controller.RoleReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunRoleReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunRoleReconciler indicates an expected call of RunRoleReconciler
func (mr *MockRoleReconcileLoopMockRecorder) RunRoleReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunRoleReconciler", reflect.TypeOf((*MockRoleReconcileLoop)(nil).RunRoleReconciler), varargs...)
}

// MockRoleBindingReconciler is a mock of RoleBindingReconciler interface
type MockRoleBindingReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockRoleBindingReconcilerMockRecorder
}

// MockRoleBindingReconcilerMockRecorder is the mock recorder for MockRoleBindingReconciler
type MockRoleBindingReconcilerMockRecorder struct {
	mock *MockRoleBindingReconciler
}

// NewMockRoleBindingReconciler creates a new mock instance
func NewMockRoleBindingReconciler(ctrl *gomock.Controller) *MockRoleBindingReconciler {
	mock := &MockRoleBindingReconciler{ctrl: ctrl}
	mock.recorder = &MockRoleBindingReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleBindingReconciler) EXPECT() *MockRoleBindingReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRoleBinding mocks base method
func (m *MockRoleBindingReconciler) ReconcileRoleBinding(obj *v1alpha1.RoleBinding) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRoleBinding", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRoleBinding indicates an expected call of ReconcileRoleBinding
func (mr *MockRoleBindingReconcilerMockRecorder) ReconcileRoleBinding(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRoleBinding", reflect.TypeOf((*MockRoleBindingReconciler)(nil).ReconcileRoleBinding), obj)
}

// MockRoleBindingDeletionReconciler is a mock of RoleBindingDeletionReconciler interface
type MockRoleBindingDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockRoleBindingDeletionReconcilerMockRecorder
}

// MockRoleBindingDeletionReconcilerMockRecorder is the mock recorder for MockRoleBindingDeletionReconciler
type MockRoleBindingDeletionReconcilerMockRecorder struct {
	mock *MockRoleBindingDeletionReconciler
}

// NewMockRoleBindingDeletionReconciler creates a new mock instance
func NewMockRoleBindingDeletionReconciler(ctrl *gomock.Controller) *MockRoleBindingDeletionReconciler {
	mock := &MockRoleBindingDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockRoleBindingDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleBindingDeletionReconciler) EXPECT() *MockRoleBindingDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRoleBindingDeletion mocks base method
func (m *MockRoleBindingDeletionReconciler) ReconcileRoleBindingDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRoleBindingDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRoleBindingDeletion indicates an expected call of ReconcileRoleBindingDeletion
func (mr *MockRoleBindingDeletionReconcilerMockRecorder) ReconcileRoleBindingDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRoleBindingDeletion", reflect.TypeOf((*MockRoleBindingDeletionReconciler)(nil).ReconcileRoleBindingDeletion), req)
}

// MockRoleBindingFinalizer is a mock of RoleBindingFinalizer interface
type MockRoleBindingFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockRoleBindingFinalizerMockRecorder
}

// MockRoleBindingFinalizerMockRecorder is the mock recorder for MockRoleBindingFinalizer
type MockRoleBindingFinalizerMockRecorder struct {
	mock *MockRoleBindingFinalizer
}

// NewMockRoleBindingFinalizer creates a new mock instance
func NewMockRoleBindingFinalizer(ctrl *gomock.Controller) *MockRoleBindingFinalizer {
	mock := &MockRoleBindingFinalizer{ctrl: ctrl}
	mock.recorder = &MockRoleBindingFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleBindingFinalizer) EXPECT() *MockRoleBindingFinalizerMockRecorder {
	return m.recorder
}

// ReconcileRoleBinding mocks base method
func (m *MockRoleBindingFinalizer) ReconcileRoleBinding(obj *v1alpha1.RoleBinding) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRoleBinding", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRoleBinding indicates an expected call of ReconcileRoleBinding
func (mr *MockRoleBindingFinalizerMockRecorder) ReconcileRoleBinding(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRoleBinding", reflect.TypeOf((*MockRoleBindingFinalizer)(nil).ReconcileRoleBinding), obj)
}

// RoleBindingFinalizerName mocks base method
func (m *MockRoleBindingFinalizer) RoleBindingFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleBindingFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// RoleBindingFinalizerName indicates an expected call of RoleBindingFinalizerName
func (mr *MockRoleBindingFinalizerMockRecorder) RoleBindingFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleBindingFinalizerName", reflect.TypeOf((*MockRoleBindingFinalizer)(nil).RoleBindingFinalizerName))
}

// FinalizeRoleBinding mocks base method
func (m *MockRoleBindingFinalizer) FinalizeRoleBinding(obj *v1alpha1.RoleBinding) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeRoleBinding", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeRoleBinding indicates an expected call of FinalizeRoleBinding
func (mr *MockRoleBindingFinalizerMockRecorder) FinalizeRoleBinding(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeRoleBinding", reflect.TypeOf((*MockRoleBindingFinalizer)(nil).FinalizeRoleBinding), obj)
}

// MockRoleBindingReconcileLoop is a mock of RoleBindingReconcileLoop interface
type MockRoleBindingReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockRoleBindingReconcileLoopMockRecorder
}

// MockRoleBindingReconcileLoopMockRecorder is the mock recorder for MockRoleBindingReconcileLoop
type MockRoleBindingReconcileLoopMockRecorder struct {
	mock *MockRoleBindingReconcileLoop
}

// NewMockRoleBindingReconcileLoop creates a new mock instance
func NewMockRoleBindingReconcileLoop(ctrl *gomock.Controller) *MockRoleBindingReconcileLoop {
	mock := &MockRoleBindingReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockRoleBindingReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleBindingReconcileLoop) EXPECT() *MockRoleBindingReconcileLoopMockRecorder {
	return m.recorder
}

// RunRoleBindingReconciler mocks base method
func (m *MockRoleBindingReconcileLoop) RunRoleBindingReconciler(ctx context.Context, rec controller.RoleBindingReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunRoleBindingReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunRoleBindingReconciler indicates an expected call of RunRoleBindingReconciler
func (mr *MockRoleBindingReconcileLoopMockRecorder) RunRoleBindingReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunRoleBindingReconciler", reflect.TypeOf((*MockRoleBindingReconcileLoop)(nil).RunRoleBindingReconciler), varargs...)
}
