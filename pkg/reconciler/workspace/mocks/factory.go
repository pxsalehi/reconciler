// Code generated by mockery v2.9.4. DO NOT EDIT.

package mock

import (
	workspace "github.com/kyma-incubator/reconciler/pkg/reconciler/workspace"
	mock "github.com/stretchr/testify/mock"
)

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// Delete provides a mock function with given fields: version
func (_m *Factory) Delete(version string) error {
	ret := _m.Called(version)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: version
func (_m *Factory) Get(version string) (*workspace.Workspace, error) {
	ret := _m.Called(version)

	var r0 *workspace.Workspace
	if rf, ok := ret.Get(0).(func(string) *workspace.Workspace); ok {
		r0 = rf(version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.Workspace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
