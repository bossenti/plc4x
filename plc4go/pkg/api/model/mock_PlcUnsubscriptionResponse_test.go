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

// Code generated by mockery v2.32.2. DO NOT EDIT.

package model

import mock "github.com/stretchr/testify/mock"

// MockPlcUnsubscriptionResponse is an autogenerated mock type for the PlcUnsubscriptionResponse type
type MockPlcUnsubscriptionResponse struct {
	mock.Mock
}

type MockPlcUnsubscriptionResponse_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcUnsubscriptionResponse) EXPECT() *MockPlcUnsubscriptionResponse_Expecter {
	return &MockPlcUnsubscriptionResponse_Expecter{mock: &_m.Mock}
}

// GetRequest provides a mock function with given fields:
func (_m *MockPlcUnsubscriptionResponse) GetRequest() PlcUnsubscriptionRequest {
	ret := _m.Called()

	var r0 PlcUnsubscriptionRequest
	if rf, ok := ret.Get(0).(func() PlcUnsubscriptionRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcUnsubscriptionRequest)
		}
	}

	return r0
}

// MockPlcUnsubscriptionResponse_GetRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRequest'
type MockPlcUnsubscriptionResponse_GetRequest_Call struct {
	*mock.Call
}

// GetRequest is a helper method to define mock.On call
func (_e *MockPlcUnsubscriptionResponse_Expecter) GetRequest() *MockPlcUnsubscriptionResponse_GetRequest_Call {
	return &MockPlcUnsubscriptionResponse_GetRequest_Call{Call: _e.mock.On("GetRequest")}
}

func (_c *MockPlcUnsubscriptionResponse_GetRequest_Call) Run(run func()) *MockPlcUnsubscriptionResponse_GetRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcUnsubscriptionResponse_GetRequest_Call) Return(_a0 PlcUnsubscriptionRequest) *MockPlcUnsubscriptionResponse_GetRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcUnsubscriptionResponse_GetRequest_Call) RunAndReturn(run func() PlcUnsubscriptionRequest) *MockPlcUnsubscriptionResponse_GetRequest_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcUnsubscriptionResponse) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcUnsubscriptionResponse_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcUnsubscriptionResponse_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcUnsubscriptionResponse_Expecter) String() *MockPlcUnsubscriptionResponse_String_Call {
	return &MockPlcUnsubscriptionResponse_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcUnsubscriptionResponse_String_Call) Run(run func()) *MockPlcUnsubscriptionResponse_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcUnsubscriptionResponse_String_Call) Return(_a0 string) *MockPlcUnsubscriptionResponse_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcUnsubscriptionResponse_String_Call) RunAndReturn(run func() string) *MockPlcUnsubscriptionResponse_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcUnsubscriptionResponse creates a new instance of MockPlcUnsubscriptionResponse. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcUnsubscriptionResponse(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcUnsubscriptionResponse {
	mock := &MockPlcUnsubscriptionResponse{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
