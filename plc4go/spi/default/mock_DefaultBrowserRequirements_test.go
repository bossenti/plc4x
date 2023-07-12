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
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"
)

// MockDefaultBrowserRequirements is an autogenerated mock type for the DefaultBrowserRequirements type
type MockDefaultBrowserRequirements struct {
	mock.Mock
}

type MockDefaultBrowserRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDefaultBrowserRequirements) EXPECT() *MockDefaultBrowserRequirements_Expecter {
	return &MockDefaultBrowserRequirements_Expecter{mock: &_m.Mock}
}

// BrowseQuery provides a mock function with given fields: ctx, interceptor, queryName, query
func (_m *MockDefaultBrowserRequirements) BrowseQuery(ctx context.Context, interceptor func(model.PlcBrowseItem) bool, queryName string, query model.PlcQuery) (model.PlcResponseCode, []model.PlcBrowseItem) {
	ret := _m.Called(ctx, interceptor, queryName, query)

	var r0 model.PlcResponseCode
	var r1 []model.PlcBrowseItem
	if rf, ok := ret.Get(0).(func(context.Context, func(model.PlcBrowseItem) bool, string, model.PlcQuery) (model.PlcResponseCode, []model.PlcBrowseItem)); ok {
		return rf(ctx, interceptor, queryName, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, func(model.PlcBrowseItem) bool, string, model.PlcQuery) model.PlcResponseCode); ok {
		r0 = rf(ctx, interceptor, queryName, query)
	} else {
		r0 = ret.Get(0).(model.PlcResponseCode)
	}

	if rf, ok := ret.Get(1).(func(context.Context, func(model.PlcBrowseItem) bool, string, model.PlcQuery) []model.PlcBrowseItem); ok {
		r1 = rf(ctx, interceptor, queryName, query)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]model.PlcBrowseItem)
		}
	}

	return r0, r1
}

// MockDefaultBrowserRequirements_BrowseQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BrowseQuery'
type MockDefaultBrowserRequirements_BrowseQuery_Call struct {
	*mock.Call
}

// BrowseQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - interceptor func(model.PlcBrowseItem) bool
//   - queryName string
//   - query model.PlcQuery
func (_e *MockDefaultBrowserRequirements_Expecter) BrowseQuery(ctx interface{}, interceptor interface{}, queryName interface{}, query interface{}) *MockDefaultBrowserRequirements_BrowseQuery_Call {
	return &MockDefaultBrowserRequirements_BrowseQuery_Call{Call: _e.mock.On("BrowseQuery", ctx, interceptor, queryName, query)}
}

func (_c *MockDefaultBrowserRequirements_BrowseQuery_Call) Run(run func(ctx context.Context, interceptor func(model.PlcBrowseItem) bool, queryName string, query model.PlcQuery)) *MockDefaultBrowserRequirements_BrowseQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(model.PlcBrowseItem) bool), args[2].(string), args[3].(model.PlcQuery))
	})
	return _c
}

func (_c *MockDefaultBrowserRequirements_BrowseQuery_Call) Return(_a0 model.PlcResponseCode, _a1 []model.PlcBrowseItem) *MockDefaultBrowserRequirements_BrowseQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDefaultBrowserRequirements_BrowseQuery_Call) RunAndReturn(run func(context.Context, func(model.PlcBrowseItem) bool, string, model.PlcQuery) (model.PlcResponseCode, []model.PlcBrowseItem)) *MockDefaultBrowserRequirements_BrowseQuery_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDefaultBrowserRequirements creates a new instance of MockDefaultBrowserRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDefaultBrowserRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDefaultBrowserRequirements {
	mock := &MockDefaultBrowserRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
