// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package zcnmocks

import mock "github.com/stretchr/testify/mock"

// AuthCallback is an autogenerated mock type for the AuthCallback type
type AuthCallback struct {
	mock.Mock
}

// OnSetupComplete provides a mock function with given fields: status, err
func (_m *AuthCallback) OnSetupComplete(status int, err string) {
	_m.Called(status, err)
}