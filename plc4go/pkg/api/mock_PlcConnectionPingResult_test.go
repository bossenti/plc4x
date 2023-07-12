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

package plc4go

import mock "github.com/stretchr/testify/mock"

// MockPlcConnectionPingResult is an autogenerated mock type for the PlcConnectionPingResult type
type MockPlcConnectionPingResult struct {
	mock.Mock
}

type MockPlcConnectionPingResult_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcConnectionPingResult) EXPECT() *MockPlcConnectionPingResult_Expecter {
	return &MockPlcConnectionPingResult_Expecter{mock: &_m.Mock}
}

// GetErr provides a mock function with given fields:
func (_m *MockPlcConnectionPingResult) GetErr() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcConnectionPingResult_GetErr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetErr'
type MockPlcConnectionPingResult_GetErr_Call struct {
	*mock.Call
}

// GetErr is a helper method to define mock.On call
func (_e *MockPlcConnectionPingResult_Expecter) GetErr() *MockPlcConnectionPingResult_GetErr_Call {
	return &MockPlcConnectionPingResult_GetErr_Call{Call: _e.mock.On("GetErr")}
}

func (_c *MockPlcConnectionPingResult_GetErr_Call) Run(run func()) *MockPlcConnectionPingResult_GetErr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionPingResult_GetErr_Call) Return(_a0 error) *MockPlcConnectionPingResult_GetErr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionPingResult_GetErr_Call) RunAndReturn(run func() error) *MockPlcConnectionPingResult_GetErr_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcConnectionPingResult creates a new instance of MockPlcConnectionPingResult. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcConnectionPingResult(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcConnectionPingResult {
	mock := &MockPlcConnectionPingResult{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
