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

package bacnetip

import mock "github.com/stretchr/testify/mock"

// mock_IOController is an autogenerated mock type for the _IOController type
type mock_IOController struct {
	mock.Mock
}

type mock_IOController_Expecter struct {
	mock *mock.Mock
}

func (_m *mock_IOController) EXPECT() *mock_IOController_Expecter {
	return &mock_IOController_Expecter{mock: &_m.Mock}
}

// Abort provides a mock function with given fields: err
func (_m *mock_IOController) Abort(err error) error {
	ret := _m.Called(err)

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_IOController_Abort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Abort'
type mock_IOController_Abort_Call struct {
	*mock.Call
}

// Abort is a helper method to define mock.On call
//   - err error
func (_e *mock_IOController_Expecter) Abort(err interface{}) *mock_IOController_Abort_Call {
	return &mock_IOController_Abort_Call{Call: _e.mock.On("Abort", err)}
}

func (_c *mock_IOController_Abort_Call) Run(run func(err error)) *mock_IOController_Abort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *mock_IOController_Abort_Call) Return(_a0 error) *mock_IOController_Abort_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_IOController_Abort_Call) RunAndReturn(run func(error) error) *mock_IOController_Abort_Call {
	_c.Call.Return(run)
	return _c
}

// AbortIO provides a mock function with given fields: iocb, err
func (_m *mock_IOController) AbortIO(iocb _IOCB, err error) error {
	ret := _m.Called(iocb, err)

	var r0 error
	if rf, ok := ret.Get(0).(func(_IOCB, error) error); ok {
		r0 = rf(iocb, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_IOController_AbortIO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AbortIO'
type mock_IOController_AbortIO_Call struct {
	*mock.Call
}

// AbortIO is a helper method to define mock.On call
//   - iocb _IOCB
//   - err error
func (_e *mock_IOController_Expecter) AbortIO(iocb interface{}, err interface{}) *mock_IOController_AbortIO_Call {
	return &mock_IOController_AbortIO_Call{Call: _e.mock.On("AbortIO", iocb, err)}
}

func (_c *mock_IOController_AbortIO_Call) Run(run func(iocb _IOCB, err error)) *mock_IOController_AbortIO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_IOCB), args[1].(error))
	})
	return _c
}

func (_c *mock_IOController_AbortIO_Call) Return(_a0 error) *mock_IOController_AbortIO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_IOController_AbortIO_Call) RunAndReturn(run func(_IOCB, error) error) *mock_IOController_AbortIO_Call {
	_c.Call.Return(run)
	return _c
}

// CompleteIO provides a mock function with given fields: iocb, pdu
func (_m *mock_IOController) CompleteIO(iocb _IOCB, pdu _PDU) error {
	ret := _m.Called(iocb, pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_IOCB, _PDU) error); ok {
		r0 = rf(iocb, pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_IOController_CompleteIO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompleteIO'
type mock_IOController_CompleteIO_Call struct {
	*mock.Call
}

// CompleteIO is a helper method to define mock.On call
//   - iocb _IOCB
//   - pdu _PDU
func (_e *mock_IOController_Expecter) CompleteIO(iocb interface{}, pdu interface{}) *mock_IOController_CompleteIO_Call {
	return &mock_IOController_CompleteIO_Call{Call: _e.mock.On("CompleteIO", iocb, pdu)}
}

func (_c *mock_IOController_CompleteIO_Call) Run(run func(iocb _IOCB, pdu _PDU)) *mock_IOController_CompleteIO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_IOCB), args[1].(_PDU))
	})
	return _c
}

func (_c *mock_IOController_CompleteIO_Call) Return(_a0 error) *mock_IOController_CompleteIO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_IOController_CompleteIO_Call) RunAndReturn(run func(_IOCB, _PDU) error) *mock_IOController_CompleteIO_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessIO provides a mock function with given fields: iocb
func (_m *mock_IOController) ProcessIO(iocb _IOCB) error {
	ret := _m.Called(iocb)

	var r0 error
	if rf, ok := ret.Get(0).(func(_IOCB) error); ok {
		r0 = rf(iocb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_IOController_ProcessIO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessIO'
type mock_IOController_ProcessIO_Call struct {
	*mock.Call
}

// ProcessIO is a helper method to define mock.On call
//   - iocb _IOCB
func (_e *mock_IOController_Expecter) ProcessIO(iocb interface{}) *mock_IOController_ProcessIO_Call {
	return &mock_IOController_ProcessIO_Call{Call: _e.mock.On("ProcessIO", iocb)}
}

func (_c *mock_IOController_ProcessIO_Call) Run(run func(iocb _IOCB)) *mock_IOController_ProcessIO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_IOCB))
	})
	return _c
}

func (_c *mock_IOController_ProcessIO_Call) Return(_a0 error) *mock_IOController_ProcessIO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_IOController_ProcessIO_Call) RunAndReturn(run func(_IOCB) error) *mock_IOController_ProcessIO_Call {
	_c.Call.Return(run)
	return _c
}

// newMock_IOController creates a new instance of mock_IOController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMock_IOController(t interface {
	mock.TestingT
	Cleanup(func())
}) *mock_IOController {
	mock := &mock_IOController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
