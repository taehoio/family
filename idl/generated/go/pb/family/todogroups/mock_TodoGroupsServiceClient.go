// Code generated by mockery v1.0.0. DO NOT EDIT.

package todogroups

import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"

// MockTodoGroupsServiceClient is an autogenerated mock type for the TodoGroupsServiceClient type
type MockTodoGroupsServiceClient struct {
	mock.Mock
}

// CreateTodoGroup provides a mock function with given fields: ctx, in, opts
func (_m *MockTodoGroupsServiceClient) CreateTodoGroup(ctx context.Context, in *CreateTodoGroupRequest, opts ...grpc.CallOption) (*CreateTodoGroupResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *CreateTodoGroupResponse
	if rf, ok := ret.Get(0).(func(context.Context, *CreateTodoGroupRequest, ...grpc.CallOption) *CreateTodoGroupResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CreateTodoGroupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *CreateTodoGroupRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTodoGroup provides a mock function with given fields: ctx, in, opts
func (_m *MockTodoGroupsServiceClient) DeleteTodoGroup(ctx context.Context, in *DeleteTodoGroupRequest, opts ...grpc.CallOption) (*DeleteTodoGroupResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *DeleteTodoGroupResponse
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteTodoGroupRequest, ...grpc.CallOption) *DeleteTodoGroupResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeleteTodoGroupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DeleteTodoGroupRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoGroup provides a mock function with given fields: ctx, in, opts
func (_m *MockTodoGroupsServiceClient) GetTodoGroup(ctx context.Context, in *GetTodoGroupRequest, opts ...grpc.CallOption) (*GetTodoGroupResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetTodoGroupResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTodoGroupRequest, ...grpc.CallOption) *GetTodoGroupResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTodoGroupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTodoGroupRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTodoGroups provides a mock function with given fields: ctx, in, opts
func (_m *MockTodoGroupsServiceClient) ListTodoGroups(ctx context.Context, in *ListTodoGroupsRequest, opts ...grpc.CallOption) (*ListTodoGroupsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ListTodoGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ListTodoGroupsRequest, ...grpc.CallOption) *ListTodoGroupsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ListTodoGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ListTodoGroupsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTodoGroup provides a mock function with given fields: ctx, in, opts
func (_m *MockTodoGroupsServiceClient) UpdateTodoGroup(ctx context.Context, in *UpdateTodoGroupRequest, opts ...grpc.CallOption) (*UpdateTodoGroupResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *UpdateTodoGroupResponse
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateTodoGroupRequest, ...grpc.CallOption) *UpdateTodoGroupResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateTodoGroupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateTodoGroupRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}