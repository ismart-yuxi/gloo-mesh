// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo-mesh/pkg/openmeshctl/runtime (interfaces: Context)

// Package mock_runtime is a generated GoMock package.
package mock_runtime

import (
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	registry "github.com/solo-io/gloo-mesh/pkg/openmeshctl/resource/registry"
	meta "k8s.io/apimachinery/pkg/api/meta"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// Deadline mocks base method.
func (m *MockContext) Deadline() (time.Time, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deadline")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Deadline indicates an expected call of Deadline.
func (mr *MockContextMockRecorder) Deadline() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deadline", reflect.TypeOf((*MockContext)(nil).Deadline))
}

// Done mocks base method.
func (m *MockContext) Done() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockContextMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockContext)(nil).Done))
}

// Err mocks base method.
func (m *MockContext) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockContextMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockContext)(nil).Err))
}

// ErrOut mocks base method.
func (m *MockContext) ErrOut() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrOut")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// ErrOut indicates an expected call of ErrOut.
func (mr *MockContextMockRecorder) ErrOut() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrOut", reflect.TypeOf((*MockContext)(nil).ErrOut))
}

// KubeClient mocks base method.
func (m *MockContext) KubeClient() (client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeClient")
	ret0, _ := ret[0].(client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KubeClient indicates an expected call of KubeClient.
func (mr *MockContextMockRecorder) KubeClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeClient", reflect.TypeOf((*MockContext)(nil).KubeClient))
}

// KubeConfig mocks base method.
func (m *MockContext) KubeConfig() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeConfig")
	ret0, _ := ret[0].(string)
	return ret0
}

// KubeConfig indicates an expected call of KubeConfig.
func (mr *MockContextMockRecorder) KubeConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeConfig", reflect.TypeOf((*MockContext)(nil).KubeConfig))
}

// KubeContext mocks base method.
func (m *MockContext) KubeContext() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeContext")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KubeContext indicates an expected call of KubeContext.
func (mr *MockContextMockRecorder) KubeContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeContext", reflect.TypeOf((*MockContext)(nil).KubeContext))
}

// Namespace mocks base method.
func (m *MockContext) Namespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// Namespace indicates an expected call of Namespace.
func (mr *MockContextMockRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockContext)(nil).Namespace))
}

// Out mocks base method.
func (m *MockContext) Out() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Out")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// Out indicates an expected call of Out.
func (mr *MockContextMockRecorder) Out() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Out", reflect.TypeOf((*MockContext)(nil).Out))
}

// Registry mocks base method.
func (m *MockContext) Registry() registry.TypeRegistry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Registry")
	ret0, _ := ret[0].(registry.TypeRegistry)
	return ret0
}

// Registry indicates an expected call of Registry.
func (mr *MockContextMockRecorder) Registry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Registry", reflect.TypeOf((*MockContext)(nil).Registry))
}

// ToDiscoveryClient mocks base method.
func (m *MockContext) ToDiscoveryClient() (discovery.CachedDiscoveryInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToDiscoveryClient")
	ret0, _ := ret[0].(discovery.CachedDiscoveryInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToDiscoveryClient indicates an expected call of ToDiscoveryClient.
func (mr *MockContextMockRecorder) ToDiscoveryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToDiscoveryClient", reflect.TypeOf((*MockContext)(nil).ToDiscoveryClient))
}

// ToRESTConfig mocks base method.
func (m *MockContext) ToRESTConfig() (*rest.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToRESTConfig")
	ret0, _ := ret[0].(*rest.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToRESTConfig indicates an expected call of ToRESTConfig.
func (mr *MockContextMockRecorder) ToRESTConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToRESTConfig", reflect.TypeOf((*MockContext)(nil).ToRESTConfig))
}

// ToRESTMapper mocks base method.
func (m *MockContext) ToRESTMapper() (meta.RESTMapper, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToRESTMapper")
	ret0, _ := ret[0].(meta.RESTMapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToRESTMapper indicates an expected call of ToRESTMapper.
func (mr *MockContextMockRecorder) ToRESTMapper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToRESTMapper", reflect.TypeOf((*MockContext)(nil).ToRESTMapper))
}

// ToRawKubeConfigLoader mocks base method.
func (m *MockContext) ToRawKubeConfigLoader() clientcmd.ClientConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToRawKubeConfigLoader")
	ret0, _ := ret[0].(clientcmd.ClientConfig)
	return ret0
}

// ToRawKubeConfigLoader indicates an expected call of ToRawKubeConfigLoader.
func (mr *MockContextMockRecorder) ToRawKubeConfigLoader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToRawKubeConfigLoader", reflect.TypeOf((*MockContext)(nil).ToRawKubeConfigLoader))
}

// Value mocks base method.
func (m *MockContext) Value(arg0 interface{}) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockContextMockRecorder) Value(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockContext)(nil).Value), arg0)
}

// Verbose mocks base method.
func (m *MockContext) Verbose() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verbose")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Verbose indicates an expected call of Verbose.
func (mr *MockContextMockRecorder) Verbose() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verbose", reflect.TypeOf((*MockContext)(nil).Verbose))
}