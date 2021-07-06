// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"
	databases "klever/grpc/databases"

	mock "github.com/stretchr/testify/mock"
)

// CollectionHelper is an autogenerated mock type for the CollectionHelper type
type CollectionHelper struct {
	mock.Mock
}

// DeleteOne provides a mock function with given fields: ctx, filter
func (_m *CollectionHelper) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	ret := _m.Called(ctx, filter)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) int64); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: _a0, _a1
func (_m *CollectionHelper) Find(_a0 context.Context, _a1 interface{}) (databases.CursorResultHelper, error) {
	ret := _m.Called(_a0, _a1)

	var r0 databases.CursorResultHelper
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) databases.CursorResultHelper); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(databases.CursorResultHelper)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: _a0, _a1
func (_m *CollectionHelper) FindOne(_a0 context.Context, _a1 interface{}) databases.SingleResultHelper {
	ret := _m.Called(_a0, _a1)

	var r0 databases.SingleResultHelper
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) databases.SingleResultHelper); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(databases.SingleResultHelper)
		}
	}

	return r0
}

// FindOneAndUpdate provides a mock function with given fields: _a0, _a1, _a2
func (_m *CollectionHelper) FindOneAndUpdate(_a0 context.Context, _a1 interface{}, _a2 interface{}) databases.SingleResultHelper {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 databases.SingleResultHelper
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}) databases.SingleResultHelper); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(databases.SingleResultHelper)
		}
	}

	return r0
}

// InsertOne provides a mock function with given fields: _a0, _a1
func (_m *CollectionHelper) InsertOne(_a0 context.Context, _a1 interface{}) (interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
