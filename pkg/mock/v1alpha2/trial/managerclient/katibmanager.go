// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeflow/katib/pkg/controller/v1alpha2/trial/managerclient (interfaces: ManagerClient)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/kubeflow/katib/pkg/api/operators/apis/trial/v1alpha2"
	v1alpha20 "github.com/kubeflow/katib/pkg/api/v1alpha2"
	reflect "reflect"
)

// MockManagerClient is a mock of ManagerClient interface
type MockManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagerClientMockRecorder
}

// MockManagerClientMockRecorder is the mock recorder for MockManagerClient
type MockManagerClientMockRecorder struct {
	mock *MockManagerClient
}

// NewMockManagerClient creates a new mock instance
func NewMockManagerClient(ctrl *gomock.Controller) *MockManagerClient {
	mock := &MockManagerClient{ctrl: ctrl}
	mock.recorder = &MockManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManagerClient) EXPECT() *MockManagerClientMockRecorder {
	return m.recorder
}

// CreateTrialInDB mocks base method
func (m *MockManagerClient) CreateTrialInDB(arg0 *v1alpha2.Trial) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrialInDB", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTrialInDB indicates an expected call of CreateTrialInDB
func (mr *MockManagerClientMockRecorder) CreateTrialInDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrialInDB", reflect.TypeOf((*MockManagerClient)(nil).CreateTrialInDB), arg0)
}

// GetTrialConf mocks base method
func (m *MockManagerClient) GetTrialConf(arg0 *v1alpha2.Trial) *v1alpha20.Trial {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrialConf", arg0)
	ret0, _ := ret[0].(*v1alpha20.Trial)
	return ret0
}

// GetTrialConf indicates an expected call of GetTrialConf
func (mr *MockManagerClientMockRecorder) GetTrialConf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrialConf", reflect.TypeOf((*MockManagerClient)(nil).GetTrialConf), arg0)
}

// GetTrialObservation mocks base method
func (m *MockManagerClient) GetTrialObservation(arg0 *v1alpha2.Trial) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrialObservation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTrialObservation indicates an expected call of GetTrialObservation
func (mr *MockManagerClientMockRecorder) GetTrialObservation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrialObservation", reflect.TypeOf((*MockManagerClient)(nil).GetTrialObservation), arg0)
}

// UpdateTrialStatusInDB mocks base method
func (m *MockManagerClient) UpdateTrialStatusInDB(arg0 *v1alpha2.Trial) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrialStatusInDB", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTrialStatusInDB indicates an expected call of UpdateTrialStatusInDB
func (mr *MockManagerClientMockRecorder) UpdateTrialStatusInDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrialStatusInDB", reflect.TypeOf((*MockManagerClient)(nil).UpdateTrialStatusInDB), arg0)
}