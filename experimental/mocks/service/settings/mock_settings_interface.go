// Code generated by mockery v2.43.0. DO NOT EDIT.

package settings

import (
	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockSettingsInterface is an autogenerated mock type for the SettingsInterface type
type MockSettingsInterface struct {
	mock.Mock
}

type MockSettingsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSettingsInterface) EXPECT() *MockSettingsInterface_Expecter {
	return &MockSettingsInterface_Expecter{mock: &_m.Mock}
}

// AutomaticClusterUpdate provides a mock function with given fields:
func (_m *MockSettingsInterface) AutomaticClusterUpdate() settings.AutomaticClusterUpdateInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AutomaticClusterUpdate")
	}

	var r0 settings.AutomaticClusterUpdateInterface
	if rf, ok := ret.Get(0).(func() settings.AutomaticClusterUpdateInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.AutomaticClusterUpdateInterface)
		}
	}

	return r0
}

// MockSettingsInterface_AutomaticClusterUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AutomaticClusterUpdate'
type MockSettingsInterface_AutomaticClusterUpdate_Call struct {
	*mock.Call
}

// AutomaticClusterUpdate is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) AutomaticClusterUpdate() *MockSettingsInterface_AutomaticClusterUpdate_Call {
	return &MockSettingsInterface_AutomaticClusterUpdate_Call{Call: _e.mock.On("AutomaticClusterUpdate")}
}

func (_c *MockSettingsInterface_AutomaticClusterUpdate_Call) Run(run func()) *MockSettingsInterface_AutomaticClusterUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_AutomaticClusterUpdate_Call) Return(_a0 settings.AutomaticClusterUpdateInterface) *MockSettingsInterface_AutomaticClusterUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_AutomaticClusterUpdate_Call) RunAndReturn(run func() settings.AutomaticClusterUpdateInterface) *MockSettingsInterface_AutomaticClusterUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// CspEnablement provides a mock function with given fields:
func (_m *MockSettingsInterface) CspEnablement() settings.CspEnablementInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CspEnablement")
	}

	var r0 settings.CspEnablementInterface
	if rf, ok := ret.Get(0).(func() settings.CspEnablementInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.CspEnablementInterface)
		}
	}

	return r0
}

// MockSettingsInterface_CspEnablement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CspEnablement'
type MockSettingsInterface_CspEnablement_Call struct {
	*mock.Call
}

// CspEnablement is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) CspEnablement() *MockSettingsInterface_CspEnablement_Call {
	return &MockSettingsInterface_CspEnablement_Call{Call: _e.mock.On("CspEnablement")}
}

func (_c *MockSettingsInterface_CspEnablement_Call) Run(run func()) *MockSettingsInterface_CspEnablement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_CspEnablement_Call) Return(_a0 settings.CspEnablementInterface) *MockSettingsInterface_CspEnablement_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_CspEnablement_Call) RunAndReturn(run func() settings.CspEnablementInterface) *MockSettingsInterface_CspEnablement_Call {
	_c.Call.Return(run)
	return _c
}

// DefaultNamespace provides a mock function with given fields:
func (_m *MockSettingsInterface) DefaultNamespace() settings.DefaultNamespaceInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DefaultNamespace")
	}

	var r0 settings.DefaultNamespaceInterface
	if rf, ok := ret.Get(0).(func() settings.DefaultNamespaceInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.DefaultNamespaceInterface)
		}
	}

	return r0
}

// MockSettingsInterface_DefaultNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DefaultNamespace'
type MockSettingsInterface_DefaultNamespace_Call struct {
	*mock.Call
}

// DefaultNamespace is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) DefaultNamespace() *MockSettingsInterface_DefaultNamespace_Call {
	return &MockSettingsInterface_DefaultNamespace_Call{Call: _e.mock.On("DefaultNamespace")}
}

func (_c *MockSettingsInterface_DefaultNamespace_Call) Run(run func()) *MockSettingsInterface_DefaultNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_DefaultNamespace_Call) Return(_a0 settings.DefaultNamespaceInterface) *MockSettingsInterface_DefaultNamespace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_DefaultNamespace_Call) RunAndReturn(run func() settings.DefaultNamespaceInterface) *MockSettingsInterface_DefaultNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// EsmEnablement provides a mock function with given fields:
func (_m *MockSettingsInterface) EsmEnablement() settings.EsmEnablementInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EsmEnablement")
	}

	var r0 settings.EsmEnablementInterface
	if rf, ok := ret.Get(0).(func() settings.EsmEnablementInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.EsmEnablementInterface)
		}
	}

	return r0
}

// MockSettingsInterface_EsmEnablement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EsmEnablement'
type MockSettingsInterface_EsmEnablement_Call struct {
	*mock.Call
}

// EsmEnablement is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) EsmEnablement() *MockSettingsInterface_EsmEnablement_Call {
	return &MockSettingsInterface_EsmEnablement_Call{Call: _e.mock.On("EsmEnablement")}
}

func (_c *MockSettingsInterface_EsmEnablement_Call) Run(run func()) *MockSettingsInterface_EsmEnablement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_EsmEnablement_Call) Return(_a0 settings.EsmEnablementInterface) *MockSettingsInterface_EsmEnablement_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_EsmEnablement_Call) RunAndReturn(run func() settings.EsmEnablementInterface) *MockSettingsInterface_EsmEnablement_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockSettingsInterface) Impl() settings.SettingsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 settings.SettingsService
	if rf, ok := ret.Get(0).(func() settings.SettingsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.SettingsService)
		}
	}

	return r0
}

// MockSettingsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockSettingsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) Impl() *MockSettingsInterface_Impl_Call {
	return &MockSettingsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockSettingsInterface_Impl_Call) Run(run func()) *MockSettingsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_Impl_Call) Return(_a0 settings.SettingsService) *MockSettingsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_Impl_Call) RunAndReturn(run func() settings.SettingsService) *MockSettingsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// RestrictWorkspaceAdmins provides a mock function with given fields:
func (_m *MockSettingsInterface) RestrictWorkspaceAdmins() settings.RestrictWorkspaceAdminsInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RestrictWorkspaceAdmins")
	}

	var r0 settings.RestrictWorkspaceAdminsInterface
	if rf, ok := ret.Get(0).(func() settings.RestrictWorkspaceAdminsInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.RestrictWorkspaceAdminsInterface)
		}
	}

	return r0
}

// MockSettingsInterface_RestrictWorkspaceAdmins_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RestrictWorkspaceAdmins'
type MockSettingsInterface_RestrictWorkspaceAdmins_Call struct {
	*mock.Call
}

// RestrictWorkspaceAdmins is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) RestrictWorkspaceAdmins() *MockSettingsInterface_RestrictWorkspaceAdmins_Call {
	return &MockSettingsInterface_RestrictWorkspaceAdmins_Call{Call: _e.mock.On("RestrictWorkspaceAdmins")}
}

func (_c *MockSettingsInterface_RestrictWorkspaceAdmins_Call) Run(run func()) *MockSettingsInterface_RestrictWorkspaceAdmins_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_RestrictWorkspaceAdmins_Call) Return(_a0 settings.RestrictWorkspaceAdminsInterface) *MockSettingsInterface_RestrictWorkspaceAdmins_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_RestrictWorkspaceAdmins_Call) RunAndReturn(run func() settings.RestrictWorkspaceAdminsInterface) *MockSettingsInterface_RestrictWorkspaceAdmins_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockSettingsInterface) WithImpl(impl settings.SettingsService) settings.SettingsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 settings.SettingsInterface
	if rf, ok := ret.Get(0).(func(settings.SettingsService) settings.SettingsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.SettingsInterface)
		}
	}

	return r0
}

// MockSettingsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockSettingsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl settings.SettingsService
func (_e *MockSettingsInterface_Expecter) WithImpl(impl interface{}) *MockSettingsInterface_WithImpl_Call {
	return &MockSettingsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockSettingsInterface_WithImpl_Call) Run(run func(impl settings.SettingsService)) *MockSettingsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(settings.SettingsService))
	})
	return _c
}

func (_c *MockSettingsInterface_WithImpl_Call) Return(_a0 settings.SettingsInterface) *MockSettingsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_WithImpl_Call) RunAndReturn(run func(settings.SettingsService) settings.SettingsInterface) *MockSettingsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSettingsInterface creates a new instance of MockSettingsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSettingsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSettingsInterface {
	mock := &MockSettingsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
