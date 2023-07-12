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

package _default

import (
	spi "github.com/apache/plc4x/plc4go/spi"
	mock "github.com/stretchr/testify/mock"
)

// MockCustomMessageHandler is an autogenerated mock type for the CustomMessageHandler type
type MockCustomMessageHandler struct {
	mock.Mock
}

type MockCustomMessageHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCustomMessageHandler) EXPECT() *MockCustomMessageHandler_Expecter {
	return &MockCustomMessageHandler_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: codec, message
func (_m *MockCustomMessageHandler) Execute(codec DefaultCodecRequirements, message spi.Message) bool {
	ret := _m.Called(codec, message)

	var r0 bool
	if rf, ok := ret.Get(0).(func(DefaultCodecRequirements, spi.Message) bool); ok {
		r0 = rf(codec, message)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockCustomMessageHandler_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockCustomMessageHandler_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - codec DefaultCodecRequirements
//   - message spi.Message
func (_e *MockCustomMessageHandler_Expecter) Execute(codec interface{}, message interface{}) *MockCustomMessageHandler_Execute_Call {
	return &MockCustomMessageHandler_Execute_Call{Call: _e.mock.On("Execute", codec, message)}
}

func (_c *MockCustomMessageHandler_Execute_Call) Run(run func(codec DefaultCodecRequirements, message spi.Message)) *MockCustomMessageHandler_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(DefaultCodecRequirements), args[1].(spi.Message))
	})
	return _c
}

func (_c *MockCustomMessageHandler_Execute_Call) Return(_a0 bool) *MockCustomMessageHandler_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCustomMessageHandler_Execute_Call) RunAndReturn(run func(DefaultCodecRequirements, spi.Message) bool) *MockCustomMessageHandler_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCustomMessageHandler creates a new instance of MockCustomMessageHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCustomMessageHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCustomMessageHandler {
	mock := &MockCustomMessageHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
