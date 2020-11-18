// Code generated by MockGen. DO NOT EDIT.
// Source: clientset.go

// Package mock_versioned is a generated GoMock package.
package mock_versioned

import (
	gomock "github.com/golang/mock/gomock"
	discovery "k8s.io/client-go/discovery"
	v1alpha1 "k8s.restdev.com/operators/pkg/client/clientset/versioned/typed/scaling/v1alpha1"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Discovery mocks base method
func (m *MockInterface) Discovery() discovery.DiscoveryInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discovery")
	ret0, _ := ret[0].(discovery.DiscoveryInterface)
	return ret0
}

// Discovery indicates an expected call of Discovery
func (mr *MockInterfaceMockRecorder) Discovery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discovery", reflect.TypeOf((*MockInterface)(nil).Discovery))
}

// ScalingV1alpha1 mocks base method
func (m *MockInterface) ScalingV1alpha1() v1alpha1.ScalingV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScalingV1alpha1")
	ret0, _ := ret[0].(v1alpha1.ScalingV1alpha1Interface)
	return ret0
}

// ScalingV1alpha1 indicates an expected call of ScalingV1alpha1
func (mr *MockInterfaceMockRecorder) ScalingV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScalingV1alpha1", reflect.TypeOf((*MockInterface)(nil).ScalingV1alpha1))
}

// Scaling mocks base method
func (m *MockInterface) Scaling() v1alpha1.ScalingV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scaling")
	ret0, _ := ret[0].(v1alpha1.ScalingV1alpha1Interface)
	return ret0
}

// Scaling indicates an expected call of Scaling
func (mr *MockInterfaceMockRecorder) Scaling() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scaling", reflect.TypeOf((*MockInterface)(nil).Scaling))
}