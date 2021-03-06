// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../formula/math.go

package mock_formula

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Math interface
type MockMath struct {
	ctrl     *gomock.Controller
	recorder *_MockMathRecorder
}

// Recorder for MockMath (not exported)
type _MockMathRecorder struct {
	mock *MockMath
}

func NewMockMath(ctrl *gomock.Controller) *MockMath {
	mock := &MockMath{ctrl: ctrl}
	mock.recorder = &_MockMathRecorder{mock}
	return mock
}

func (_m *MockMath) EXPECT() *_MockMathRecorder {
	return _m.recorder
}

func (_m *MockMath) DoSum() int {
	ret := _m.ctrl.Call(_m, "DoSum")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockMathRecorder) DoSum() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DoSum")
}

func (_m *MockMath) DoArea() int {
	ret := _m.ctrl.Call(_m, "DoArea")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockMathRecorder) DoArea() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DoArea")
}
