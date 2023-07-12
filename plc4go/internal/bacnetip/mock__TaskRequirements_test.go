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

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// mock_TaskRequirements is an autogenerated mock type for the _TaskRequirements type
type mock_TaskRequirements struct {
	mock.Mock
}

type mock_TaskRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *mock_TaskRequirements) EXPECT() *mock_TaskRequirements_Expecter {
	return &mock_TaskRequirements_Expecter{mock: &_m.Mock}
}

// InstallTask provides a mock function with given fields: when, delta
func (_m *mock_TaskRequirements) InstallTask(when *time.Time, delta *time.Duration) {
	_m.Called(when, delta)
}

// mock_TaskRequirements_InstallTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InstallTask'
type mock_TaskRequirements_InstallTask_Call struct {
	*mock.Call
}

// InstallTask is a helper method to define mock.On call
//   - when *time.Time
//   - delta *time.Duration
func (_e *mock_TaskRequirements_Expecter) InstallTask(when interface{}, delta interface{}) *mock_TaskRequirements_InstallTask_Call {
	return &mock_TaskRequirements_InstallTask_Call{Call: _e.mock.On("InstallTask", when, delta)}
}

func (_c *mock_TaskRequirements_InstallTask_Call) Run(run func(when *time.Time, delta *time.Duration)) *mock_TaskRequirements_InstallTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*time.Time), args[1].(*time.Duration))
	})
	return _c
}

func (_c *mock_TaskRequirements_InstallTask_Call) Return() *mock_TaskRequirements_InstallTask_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_TaskRequirements_InstallTask_Call) RunAndReturn(run func(*time.Time, *time.Duration)) *mock_TaskRequirements_InstallTask_Call {
	_c.Call.Return(run)
	return _c
}

// getIsScheduled provides a mock function with given fields:
func (_m *mock_TaskRequirements) getIsScheduled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// mock_TaskRequirements_getIsScheduled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getIsScheduled'
type mock_TaskRequirements_getIsScheduled_Call struct {
	*mock.Call
}

// getIsScheduled is a helper method to define mock.On call
func (_e *mock_TaskRequirements_Expecter) getIsScheduled() *mock_TaskRequirements_getIsScheduled_Call {
	return &mock_TaskRequirements_getIsScheduled_Call{Call: _e.mock.On("getIsScheduled")}
}

func (_c *mock_TaskRequirements_getIsScheduled_Call) Run(run func()) *mock_TaskRequirements_getIsScheduled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_TaskRequirements_getIsScheduled_Call) Return(_a0 bool) *mock_TaskRequirements_getIsScheduled_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_TaskRequirements_getIsScheduled_Call) RunAndReturn(run func() bool) *mock_TaskRequirements_getIsScheduled_Call {
	_c.Call.Return(run)
	return _c
}

// getTaskTime provides a mock function with given fields:
func (_m *mock_TaskRequirements) getTaskTime() *time.Time {
	ret := _m.Called()

	var r0 *time.Time
	if rf, ok := ret.Get(0).(func() *time.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Time)
		}
	}

	return r0
}

// mock_TaskRequirements_getTaskTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getTaskTime'
type mock_TaskRequirements_getTaskTime_Call struct {
	*mock.Call
}

// getTaskTime is a helper method to define mock.On call
func (_e *mock_TaskRequirements_Expecter) getTaskTime() *mock_TaskRequirements_getTaskTime_Call {
	return &mock_TaskRequirements_getTaskTime_Call{Call: _e.mock.On("getTaskTime")}
}

func (_c *mock_TaskRequirements_getTaskTime_Call) Run(run func()) *mock_TaskRequirements_getTaskTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_TaskRequirements_getTaskTime_Call) Return(_a0 *time.Time) *mock_TaskRequirements_getTaskTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_TaskRequirements_getTaskTime_Call) RunAndReturn(run func() *time.Time) *mock_TaskRequirements_getTaskTime_Call {
	_c.Call.Return(run)
	return _c
}

// processTask provides a mock function with given fields:
func (_m *mock_TaskRequirements) processTask() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_TaskRequirements_processTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'processTask'
type mock_TaskRequirements_processTask_Call struct {
	*mock.Call
}

// processTask is a helper method to define mock.On call
func (_e *mock_TaskRequirements_Expecter) processTask() *mock_TaskRequirements_processTask_Call {
	return &mock_TaskRequirements_processTask_Call{Call: _e.mock.On("processTask")}
}

func (_c *mock_TaskRequirements_processTask_Call) Run(run func()) *mock_TaskRequirements_processTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_TaskRequirements_processTask_Call) Return(_a0 error) *mock_TaskRequirements_processTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_TaskRequirements_processTask_Call) RunAndReturn(run func() error) *mock_TaskRequirements_processTask_Call {
	_c.Call.Return(run)
	return _c
}

// setIsScheduled provides a mock function with given fields: isScheduled
func (_m *mock_TaskRequirements) setIsScheduled(isScheduled bool) {
	_m.Called(isScheduled)
}

// mock_TaskRequirements_setIsScheduled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setIsScheduled'
type mock_TaskRequirements_setIsScheduled_Call struct {
	*mock.Call
}

// setIsScheduled is a helper method to define mock.On call
//   - isScheduled bool
func (_e *mock_TaskRequirements_Expecter) setIsScheduled(isScheduled interface{}) *mock_TaskRequirements_setIsScheduled_Call {
	return &mock_TaskRequirements_setIsScheduled_Call{Call: _e.mock.On("setIsScheduled", isScheduled)}
}

func (_c *mock_TaskRequirements_setIsScheduled_Call) Run(run func(isScheduled bool)) *mock_TaskRequirements_setIsScheduled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *mock_TaskRequirements_setIsScheduled_Call) Return() *mock_TaskRequirements_setIsScheduled_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_TaskRequirements_setIsScheduled_Call) RunAndReturn(run func(bool)) *mock_TaskRequirements_setIsScheduled_Call {
	_c.Call.Return(run)
	return _c
}

// newMock_TaskRequirements creates a new instance of mock_TaskRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMock_TaskRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *mock_TaskRequirements {
	mock := &mock_TaskRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
