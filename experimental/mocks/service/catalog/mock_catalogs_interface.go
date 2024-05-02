// Code generated by mockery v2.43.0. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockCatalogsInterface is an autogenerated mock type for the CatalogsInterface type
type MockCatalogsInterface struct {
	mock.Mock
}

type MockCatalogsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCatalogsInterface) EXPECT() *MockCatalogsInterface_Expecter {
	return &MockCatalogsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockCatalogsInterface) Create(ctx context.Context, request catalog.CreateCatalog) (*catalog.CatalogInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.CatalogInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateCatalog) (*catalog.CatalogInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateCatalog) *catalog.CatalogInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.CatalogInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.CreateCatalog) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCatalogsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockCatalogsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.CreateCatalog
func (_e *MockCatalogsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockCatalogsInterface_Create_Call {
	return &MockCatalogsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockCatalogsInterface_Create_Call) Run(run func(ctx context.Context, request catalog.CreateCatalog)) *MockCatalogsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.CreateCatalog))
	})
	return _c
}

func (_c *MockCatalogsInterface_Create_Call) Return(_a0 *catalog.CatalogInfo, _a1 error) *MockCatalogsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCatalogsInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.CreateCatalog) (*catalog.CatalogInfo, error)) *MockCatalogsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockCatalogsInterface) Delete(ctx context.Context, request catalog.DeleteCatalogRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteCatalogRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCatalogsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockCatalogsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteCatalogRequest
func (_e *MockCatalogsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockCatalogsInterface_Delete_Call {
	return &MockCatalogsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockCatalogsInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteCatalogRequest)) *MockCatalogsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteCatalogRequest))
	})
	return _c
}

func (_c *MockCatalogsInterface_Delete_Call) Return(_a0 error) *MockCatalogsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCatalogsInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteCatalogRequest) error) *MockCatalogsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByName provides a mock function with given fields: ctx, name
func (_m *MockCatalogsInterface) DeleteByName(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCatalogsInterface_DeleteByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByName'
type MockCatalogsInterface_DeleteByName_Call struct {
	*mock.Call
}

// DeleteByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockCatalogsInterface_Expecter) DeleteByName(ctx interface{}, name interface{}) *MockCatalogsInterface_DeleteByName_Call {
	return &MockCatalogsInterface_DeleteByName_Call{Call: _e.mock.On("DeleteByName", ctx, name)}
}

func (_c *MockCatalogsInterface_DeleteByName_Call) Run(run func(ctx context.Context, name string)) *MockCatalogsInterface_DeleteByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCatalogsInterface_DeleteByName_Call) Return(_a0 error) *MockCatalogsInterface_DeleteByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCatalogsInterface_DeleteByName_Call) RunAndReturn(run func(context.Context, string) error) *MockCatalogsInterface_DeleteByName_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockCatalogsInterface) Get(ctx context.Context, request catalog.GetCatalogRequest) (*catalog.CatalogInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.CatalogInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetCatalogRequest) (*catalog.CatalogInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetCatalogRequest) *catalog.CatalogInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.CatalogInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetCatalogRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCatalogsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCatalogsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetCatalogRequest
func (_e *MockCatalogsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockCatalogsInterface_Get_Call {
	return &MockCatalogsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockCatalogsInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetCatalogRequest)) *MockCatalogsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetCatalogRequest))
	})
	return _c
}

func (_c *MockCatalogsInterface_Get_Call) Return(_a0 *catalog.CatalogInfo, _a1 error) *MockCatalogsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCatalogsInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetCatalogRequest) (*catalog.CatalogInfo, error)) *MockCatalogsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockCatalogsInterface) GetByName(ctx context.Context, name string) (*catalog.CatalogInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *catalog.CatalogInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.CatalogInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.CatalogInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.CatalogInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCatalogsInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockCatalogsInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockCatalogsInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockCatalogsInterface_GetByName_Call {
	return &MockCatalogsInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockCatalogsInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockCatalogsInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCatalogsInterface_GetByName_Call) Return(_a0 *catalog.CatalogInfo, _a1 error) *MockCatalogsInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCatalogsInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.CatalogInfo, error)) *MockCatalogsInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockCatalogsInterface) Impl() catalog.CatalogsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 catalog.CatalogsService
	if rf, ok := ret.Get(0).(func() catalog.CatalogsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.CatalogsService)
		}
	}

	return r0
}

// MockCatalogsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockCatalogsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockCatalogsInterface_Expecter) Impl() *MockCatalogsInterface_Impl_Call {
	return &MockCatalogsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockCatalogsInterface_Impl_Call) Run(run func()) *MockCatalogsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCatalogsInterface_Impl_Call) Return(_a0 catalog.CatalogsService) *MockCatalogsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCatalogsInterface_Impl_Call) RunAndReturn(run func() catalog.CatalogsService) *MockCatalogsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockCatalogsInterface) List(ctx context.Context, request catalog.ListCatalogsRequest) listing.Iterator[catalog.CatalogInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[catalog.CatalogInfo]
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListCatalogsRequest) listing.Iterator[catalog.CatalogInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[catalog.CatalogInfo])
		}
	}

	return r0
}

// MockCatalogsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockCatalogsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListCatalogsRequest
func (_e *MockCatalogsInterface_Expecter) List(ctx interface{}, request interface{}) *MockCatalogsInterface_List_Call {
	return &MockCatalogsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockCatalogsInterface_List_Call) Run(run func(ctx context.Context, request catalog.ListCatalogsRequest)) *MockCatalogsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListCatalogsRequest))
	})
	return _c
}

func (_c *MockCatalogsInterface_List_Call) Return(_a0 listing.Iterator[catalog.CatalogInfo]) *MockCatalogsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCatalogsInterface_List_Call) RunAndReturn(run func(context.Context, catalog.ListCatalogsRequest) listing.Iterator[catalog.CatalogInfo]) *MockCatalogsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockCatalogsInterface) ListAll(ctx context.Context, request catalog.ListCatalogsRequest) ([]catalog.CatalogInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []catalog.CatalogInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListCatalogsRequest) ([]catalog.CatalogInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListCatalogsRequest) []catalog.CatalogInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.CatalogInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListCatalogsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCatalogsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockCatalogsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListCatalogsRequest
func (_e *MockCatalogsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockCatalogsInterface_ListAll_Call {
	return &MockCatalogsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockCatalogsInterface_ListAll_Call) Run(run func(ctx context.Context, request catalog.ListCatalogsRequest)) *MockCatalogsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListCatalogsRequest))
	})
	return _c
}

func (_c *MockCatalogsInterface_ListAll_Call) Return(_a0 []catalog.CatalogInfo, _a1 error) *MockCatalogsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCatalogsInterface_ListAll_Call) RunAndReturn(run func(context.Context, catalog.ListCatalogsRequest) ([]catalog.CatalogInfo, error)) *MockCatalogsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockCatalogsInterface) Update(ctx context.Context, request catalog.UpdateCatalog) (*catalog.CatalogInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *catalog.CatalogInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateCatalog) (*catalog.CatalogInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateCatalog) *catalog.CatalogInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.CatalogInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.UpdateCatalog) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCatalogsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockCatalogsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdateCatalog
func (_e *MockCatalogsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockCatalogsInterface_Update_Call {
	return &MockCatalogsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockCatalogsInterface_Update_Call) Run(run func(ctx context.Context, request catalog.UpdateCatalog)) *MockCatalogsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdateCatalog))
	})
	return _c
}

func (_c *MockCatalogsInterface_Update_Call) Return(_a0 *catalog.CatalogInfo, _a1 error) *MockCatalogsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCatalogsInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.UpdateCatalog) (*catalog.CatalogInfo, error)) *MockCatalogsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockCatalogsInterface) WithImpl(impl catalog.CatalogsService) catalog.CatalogsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 catalog.CatalogsInterface
	if rf, ok := ret.Get(0).(func(catalog.CatalogsService) catalog.CatalogsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.CatalogsInterface)
		}
	}

	return r0
}

// MockCatalogsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockCatalogsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl catalog.CatalogsService
func (_e *MockCatalogsInterface_Expecter) WithImpl(impl interface{}) *MockCatalogsInterface_WithImpl_Call {
	return &MockCatalogsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockCatalogsInterface_WithImpl_Call) Run(run func(impl catalog.CatalogsService)) *MockCatalogsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(catalog.CatalogsService))
	})
	return _c
}

func (_c *MockCatalogsInterface_WithImpl_Call) Return(_a0 catalog.CatalogsInterface) *MockCatalogsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCatalogsInterface_WithImpl_Call) RunAndReturn(run func(catalog.CatalogsService) catalog.CatalogsInterface) *MockCatalogsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCatalogsInterface creates a new instance of MockCatalogsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCatalogsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCatalogsInterface {
	mock := &MockCatalogsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
