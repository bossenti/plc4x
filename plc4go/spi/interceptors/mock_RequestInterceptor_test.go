/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.32.0. DO NOT EDIT.

package interceptors

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"
)

// MockRequestInterceptor is an autogenerated mock type for the RequestInterceptor type
type MockRequestInterceptor struct {
	mock.Mock
}

type MockRequestInterceptor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRequestInterceptor) EXPECT() *MockRequestInterceptor_Expecter {
	return &MockRequestInterceptor_Expecter{mock: &_m.Mock}
}

// InterceptReadRequest provides a mock function with given fields: ctx, readRequest
func (_m *MockRequestInterceptor) InterceptReadRequest(ctx context.Context, readRequest model.PlcReadRequest) []model.PlcReadRequest {
	ret := _m.Called(ctx, readRequest)

	var r0 []model.PlcReadRequest
	if rf, ok := ret.Get(0).(func(context.Context, model.PlcReadRequest) []model.PlcReadRequest); ok {
		r0 = rf(ctx, readRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PlcReadRequest)
		}
	}

	return r0
}

// MockRequestInterceptor_InterceptReadRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InterceptReadRequest'
type MockRequestInterceptor_InterceptReadRequest_Call struct {
	*mock.Call
}

// InterceptReadRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - readRequest model.PlcReadRequest
func (_e *MockRequestInterceptor_Expecter) InterceptReadRequest(ctx interface{}, readRequest interface{}) *MockRequestInterceptor_InterceptReadRequest_Call {
	return &MockRequestInterceptor_InterceptReadRequest_Call{Call: _e.mock.On("InterceptReadRequest", ctx, readRequest)}
}

func (_c *MockRequestInterceptor_InterceptReadRequest_Call) Run(run func(ctx context.Context, readRequest model.PlcReadRequest)) *MockRequestInterceptor_InterceptReadRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PlcReadRequest))
	})
	return _c
}

func (_c *MockRequestInterceptor_InterceptReadRequest_Call) Return(_a0 []model.PlcReadRequest) *MockRequestInterceptor_InterceptReadRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestInterceptor_InterceptReadRequest_Call) RunAndReturn(run func(context.Context, model.PlcReadRequest) []model.PlcReadRequest) *MockRequestInterceptor_InterceptReadRequest_Call {
	_c.Call.Return(run)
	return _c
}

// InterceptWriteRequest provides a mock function with given fields: ctx, writeRequest
func (_m *MockRequestInterceptor) InterceptWriteRequest(ctx context.Context, writeRequest model.PlcWriteRequest) []model.PlcWriteRequest {
	ret := _m.Called(ctx, writeRequest)

	var r0 []model.PlcWriteRequest
	if rf, ok := ret.Get(0).(func(context.Context, model.PlcWriteRequest) []model.PlcWriteRequest); ok {
		r0 = rf(ctx, writeRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PlcWriteRequest)
		}
	}

	return r0
}

// MockRequestInterceptor_InterceptWriteRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InterceptWriteRequest'
type MockRequestInterceptor_InterceptWriteRequest_Call struct {
	*mock.Call
}

// InterceptWriteRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - writeRequest model.PlcWriteRequest
func (_e *MockRequestInterceptor_Expecter) InterceptWriteRequest(ctx interface{}, writeRequest interface{}) *MockRequestInterceptor_InterceptWriteRequest_Call {
	return &MockRequestInterceptor_InterceptWriteRequest_Call{Call: _e.mock.On("InterceptWriteRequest", ctx, writeRequest)}
}

func (_c *MockRequestInterceptor_InterceptWriteRequest_Call) Run(run func(ctx context.Context, writeRequest model.PlcWriteRequest)) *MockRequestInterceptor_InterceptWriteRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PlcWriteRequest))
	})
	return _c
}

func (_c *MockRequestInterceptor_InterceptWriteRequest_Call) Return(_a0 []model.PlcWriteRequest) *MockRequestInterceptor_InterceptWriteRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestInterceptor_InterceptWriteRequest_Call) RunAndReturn(run func(context.Context, model.PlcWriteRequest) []model.PlcWriteRequest) *MockRequestInterceptor_InterceptWriteRequest_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessReadResponses provides a mock function with given fields: ctx, readRequest, readResults
func (_m *MockRequestInterceptor) ProcessReadResponses(ctx context.Context, readRequest model.PlcReadRequest, readResults []model.PlcReadRequestResult) model.PlcReadRequestResult {
	ret := _m.Called(ctx, readRequest, readResults)

	var r0 model.PlcReadRequestResult
	if rf, ok := ret.Get(0).(func(context.Context, model.PlcReadRequest, []model.PlcReadRequestResult) model.PlcReadRequestResult); ok {
		r0 = rf(ctx, readRequest, readResults)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcReadRequestResult)
		}
	}

	return r0
}

// MockRequestInterceptor_ProcessReadResponses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessReadResponses'
type MockRequestInterceptor_ProcessReadResponses_Call struct {
	*mock.Call
}

// ProcessReadResponses is a helper method to define mock.On call
//   - ctx context.Context
//   - readRequest model.PlcReadRequest
//   - readResults []model.PlcReadRequestResult
func (_e *MockRequestInterceptor_Expecter) ProcessReadResponses(ctx interface{}, readRequest interface{}, readResults interface{}) *MockRequestInterceptor_ProcessReadResponses_Call {
	return &MockRequestInterceptor_ProcessReadResponses_Call{Call: _e.mock.On("ProcessReadResponses", ctx, readRequest, readResults)}
}

func (_c *MockRequestInterceptor_ProcessReadResponses_Call) Run(run func(ctx context.Context, readRequest model.PlcReadRequest, readResults []model.PlcReadRequestResult)) *MockRequestInterceptor_ProcessReadResponses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PlcReadRequest), args[2].([]model.PlcReadRequestResult))
	})
	return _c
}

func (_c *MockRequestInterceptor_ProcessReadResponses_Call) Return(_a0 model.PlcReadRequestResult) *MockRequestInterceptor_ProcessReadResponses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestInterceptor_ProcessReadResponses_Call) RunAndReturn(run func(context.Context, model.PlcReadRequest, []model.PlcReadRequestResult) model.PlcReadRequestResult) *MockRequestInterceptor_ProcessReadResponses_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessWriteResponses provides a mock function with given fields: ctx, writeRequest, writeResults
func (_m *MockRequestInterceptor) ProcessWriteResponses(ctx context.Context, writeRequest model.PlcWriteRequest, writeResults []model.PlcWriteRequestResult) model.PlcWriteRequestResult {
	ret := _m.Called(ctx, writeRequest, writeResults)

	var r0 model.PlcWriteRequestResult
	if rf, ok := ret.Get(0).(func(context.Context, model.PlcWriteRequest, []model.PlcWriteRequestResult) model.PlcWriteRequestResult); ok {
		r0 = rf(ctx, writeRequest, writeResults)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcWriteRequestResult)
		}
	}

	return r0
}

// MockRequestInterceptor_ProcessWriteResponses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessWriteResponses'
type MockRequestInterceptor_ProcessWriteResponses_Call struct {
	*mock.Call
}

// ProcessWriteResponses is a helper method to define mock.On call
//   - ctx context.Context
//   - writeRequest model.PlcWriteRequest
//   - writeResults []model.PlcWriteRequestResult
func (_e *MockRequestInterceptor_Expecter) ProcessWriteResponses(ctx interface{}, writeRequest interface{}, writeResults interface{}) *MockRequestInterceptor_ProcessWriteResponses_Call {
	return &MockRequestInterceptor_ProcessWriteResponses_Call{Call: _e.mock.On("ProcessWriteResponses", ctx, writeRequest, writeResults)}
}

func (_c *MockRequestInterceptor_ProcessWriteResponses_Call) Run(run func(ctx context.Context, writeRequest model.PlcWriteRequest, writeResults []model.PlcWriteRequestResult)) *MockRequestInterceptor_ProcessWriteResponses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PlcWriteRequest), args[2].([]model.PlcWriteRequestResult))
	})
	return _c
}

func (_c *MockRequestInterceptor_ProcessWriteResponses_Call) Return(_a0 model.PlcWriteRequestResult) *MockRequestInterceptor_ProcessWriteResponses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestInterceptor_ProcessWriteResponses_Call) RunAndReturn(run func(context.Context, model.PlcWriteRequest, []model.PlcWriteRequestResult) model.PlcWriteRequestResult) *MockRequestInterceptor_ProcessWriteResponses_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRequestInterceptor creates a new instance of MockRequestInterceptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRequestInterceptor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRequestInterceptor {
	mock := &MockRequestInterceptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
