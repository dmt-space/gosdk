// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package zcnmocks

import mock "github.com/stretchr/testify/mock"

// GetBalanceCallback is an autogenerated mock type for the GetBalanceCallback type
type GetBalanceCallback struct {
	mock.Mock
}

// OnBalanceAvailable provides a mock function with given fields: status, value, info
func (_m *GetBalanceCallback) OnBalanceAvailable(status int, value int64, info string) {
	_m.Called(status, value, info)
}