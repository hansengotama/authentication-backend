// Code generated by mockery v2.38.0. DO NOT EDIT.

package validateotpauthservice

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockValidateOTPAuthServiceInterface is an autogenerated mock type for the ValidateOTPAuthServiceInterface type
type MockValidateOTPAuthServiceInterface struct {
	mock.Mock
}

// ValidateOTPAuth provides a mock function with given fields: _a0, _a1
func (_m *MockValidateOTPAuthServiceInterface) ValidateOTPAuth(_a0 context.Context, _a1 ValidateOTPAuthParam) (ValidateOTPAuthRes, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ValidateOTPAuth")
	}

	var r0 ValidateOTPAuthRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ValidateOTPAuthParam) (ValidateOTPAuthRes, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ValidateOTPAuthParam) ValidateOTPAuthRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(ValidateOTPAuthRes)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ValidateOTPAuthParam) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockValidateOTPAuthServiceInterface creates a new instance of MockValidateOTPAuthServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockValidateOTPAuthServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockValidateOTPAuthServiceInterface {
	mock := &MockValidateOTPAuthServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
