// Code generated by mockery v1.0.0. DO NOT EDIT.

package auth

import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"

// MockAuthServiceClient is an autogenerated mock type for the AuthServiceClient type
type MockAuthServiceClient struct {
	mock.Mock
}

// Auth provides a mock function with given fields: ctx, in, opts
func (_m *MockAuthServiceClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *AuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *AuthRequest, ...grpc.CallOption) *AuthResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *AuthRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Parse provides a mock function with given fields: ctx, in, opts
func (_m *MockAuthServiceClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ParseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ParseRequest, ...grpc.CallOption) *ParseResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ParseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ParseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Refresh provides a mock function with given fields: ctx, in, opts
func (_m *MockAuthServiceClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *RefreshResponse
	if rf, ok := ret.Get(0).(func(context.Context, *RefreshRequest, ...grpc.CallOption) *RefreshResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RefreshResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *RefreshRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: ctx, in, opts
func (_m *MockAuthServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ValidateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ValidateRequest, ...grpc.CallOption) *ValidateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ValidateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ValidateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
