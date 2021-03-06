// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iost-official/Go-IOS-Protocol/protocol (interfaces: View)

// Package protocol_mock is a generated GoMock package.
package protocol_mock

import (
	gomock "github.com/golang/mock/gomock"
	iosbase "github.com/iost-official/Go-IOS-Protocol/iosbase"
	reflect "reflect"
)

// MockView is a mock of View interface
type MockView struct {
	ctrl     *gomock.Controller
	recorder *MockViewMockRecorder
}

// MockViewMockRecorder is the mock recorder for MockView
type MockViewMockRecorder struct {
	mock *MockView
}

// NewMockView creates a new mock instance
func NewMockView(ctrl *gomock.Controller) *MockView {
	mock := &MockView{ctrl: ctrl}
	mock.recorder = &MockViewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockView) EXPECT() *MockViewMockRecorder {
	return m.recorder
}

// ByzantineTolerance mocks base method
func (m *MockView) ByzantineTolerance() int {
	ret := m.ctrl.Call(m, "ByzantineTolerance")
	ret0, _ := ret[0].(int)
	return ret0
}

// ByzantineTolerance indicates an expected call of ByzantineTolerance
func (mr *MockViewMockRecorder) ByzantineTolerance() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByzantineTolerance", reflect.TypeOf((*MockView)(nil).ByzantineTolerance))
}

// CommitteeSize mocks base method
func (m *MockView) CommitteeSize() int {
	ret := m.ctrl.Call(m, "CommitteeSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// CommitteeSize indicates an expected call of CommitteeSize
func (mr *MockViewMockRecorder) CommitteeSize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitteeSize", reflect.TypeOf((*MockView)(nil).CommitteeSize))
}

// GetBackup mocks base method
func (m *MockView) GetBackup() []iosbase.Member {
	ret := m.ctrl.Call(m, "GetBackup")
	ret0, _ := ret[0].([]iosbase.Member)
	return ret0
}

// GetBackup indicates an expected call of GetBackup
func (mr *MockViewMockRecorder) GetBackup() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackup", reflect.TypeOf((*MockView)(nil).GetBackup))
}

// GetPrimary mocks base method
func (m *MockView) GetPrimary() iosbase.Member {
	ret := m.ctrl.Call(m, "GetPrimary")
	ret0, _ := ret[0].(iosbase.Member)
	return ret0
}

// GetPrimary indicates an expected call of GetPrimary
func (mr *MockViewMockRecorder) GetPrimary() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimary", reflect.TypeOf((*MockView)(nil).GetPrimary))
}

// Init mocks base method
func (m *MockView) Init(arg0 iosbase.BlockChain) {
	m.ctrl.Call(m, "Init", arg0)
}

// Init indicates an expected call of Init
func (mr *MockViewMockRecorder) Init(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockView)(nil).Init), arg0)
}

// IsBackup mocks base method
func (m *MockView) IsBackup(arg0 string) bool {
	ret := m.ctrl.Call(m, "IsBackup", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBackup indicates an expected call of IsBackup
func (mr *MockViewMockRecorder) IsBackup(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBackup", reflect.TypeOf((*MockView)(nil).IsBackup), arg0)
}

// IsPrimary mocks base method
func (m *MockView) IsPrimary(arg0 string) bool {
	ret := m.ctrl.Call(m, "IsPrimary", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPrimary indicates an expected call of IsPrimary
func (mr *MockViewMockRecorder) IsPrimary(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPrimary", reflect.TypeOf((*MockView)(nil).IsPrimary), arg0)
}
