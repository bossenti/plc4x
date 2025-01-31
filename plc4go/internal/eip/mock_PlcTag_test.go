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

package eip

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	readwritemodel "github.com/apache/plc4x/plc4go/protocols/eip/readwrite/model"

	utils "github.com/apache/plc4x/plc4go/spi/utils"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// MockPlcTag is an autogenerated mock type for the PlcTag type
type MockPlcTag struct {
	mock.Mock
}

type MockPlcTag_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcTag) EXPECT() *MockPlcTag_Expecter {
	return &MockPlcTag_Expecter{mock: &_m.Mock}
}

// GetAddressString provides a mock function with given fields:
func (_m *MockPlcTag) GetAddressString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcTag_GetAddressString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddressString'
type MockPlcTag_GetAddressString_Call struct {
	*mock.Call
}

// GetAddressString is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetAddressString() *MockPlcTag_GetAddressString_Call {
	return &MockPlcTag_GetAddressString_Call{Call: _e.mock.On("GetAddressString")}
}

func (_c *MockPlcTag_GetAddressString_Call) Run(run func()) *MockPlcTag_GetAddressString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetAddressString_Call) Return(_a0 string) *MockPlcTag_GetAddressString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetAddressString_Call) RunAndReturn(run func() string) *MockPlcTag_GetAddressString_Call {
	_c.Call.Return(run)
	return _c
}

// GetArrayInfo provides a mock function with given fields:
func (_m *MockPlcTag) GetArrayInfo() []model.ArrayInfo {
	ret := _m.Called()

	var r0 []model.ArrayInfo
	if rf, ok := ret.Get(0).(func() []model.ArrayInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ArrayInfo)
		}
	}

	return r0
}

// MockPlcTag_GetArrayInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArrayInfo'
type MockPlcTag_GetArrayInfo_Call struct {
	*mock.Call
}

// GetArrayInfo is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetArrayInfo() *MockPlcTag_GetArrayInfo_Call {
	return &MockPlcTag_GetArrayInfo_Call{Call: _e.mock.On("GetArrayInfo")}
}

func (_c *MockPlcTag_GetArrayInfo_Call) Run(run func()) *MockPlcTag_GetArrayInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetArrayInfo_Call) Return(_a0 []model.ArrayInfo) *MockPlcTag_GetArrayInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetArrayInfo_Call) RunAndReturn(run func() []model.ArrayInfo) *MockPlcTag_GetArrayInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetElementNb provides a mock function with given fields:
func (_m *MockPlcTag) GetElementNb() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockPlcTag_GetElementNb_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetElementNb'
type MockPlcTag_GetElementNb_Call struct {
	*mock.Call
}

// GetElementNb is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetElementNb() *MockPlcTag_GetElementNb_Call {
	return &MockPlcTag_GetElementNb_Call{Call: _e.mock.On("GetElementNb")}
}

func (_c *MockPlcTag_GetElementNb_Call) Run(run func()) *MockPlcTag_GetElementNb_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetElementNb_Call) Return(_a0 uint16) *MockPlcTag_GetElementNb_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetElementNb_Call) RunAndReturn(run func() uint16) *MockPlcTag_GetElementNb_Call {
	_c.Call.Return(run)
	return _c
}

// GetTag provides a mock function with given fields:
func (_m *MockPlcTag) GetTag() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcTag_GetTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTag'
type MockPlcTag_GetTag_Call struct {
	*mock.Call
}

// GetTag is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetTag() *MockPlcTag_GetTag_Call {
	return &MockPlcTag_GetTag_Call{Call: _e.mock.On("GetTag")}
}

func (_c *MockPlcTag_GetTag_Call) Run(run func()) *MockPlcTag_GetTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetTag_Call) Return(_a0 string) *MockPlcTag_GetTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetTag_Call) RunAndReturn(run func() string) *MockPlcTag_GetTag_Call {
	_c.Call.Return(run)
	return _c
}

// GetType provides a mock function with given fields:
func (_m *MockPlcTag) GetType() readwritemodel.CIPDataTypeCode {
	ret := _m.Called()

	var r0 readwritemodel.CIPDataTypeCode
	if rf, ok := ret.Get(0).(func() readwritemodel.CIPDataTypeCode); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(readwritemodel.CIPDataTypeCode)
	}

	return r0
}

// MockPlcTag_GetType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetType'
type MockPlcTag_GetType_Call struct {
	*mock.Call
}

// GetType is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetType() *MockPlcTag_GetType_Call {
	return &MockPlcTag_GetType_Call{Call: _e.mock.On("GetType")}
}

func (_c *MockPlcTag_GetType_Call) Run(run func()) *MockPlcTag_GetType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetType_Call) Return(_a0 readwritemodel.CIPDataTypeCode) *MockPlcTag_GetType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetType_Call) RunAndReturn(run func() readwritemodel.CIPDataTypeCode) *MockPlcTag_GetType_Call {
	_c.Call.Return(run)
	return _c
}

// GetValueType provides a mock function with given fields:
func (_m *MockPlcTag) GetValueType() values.PlcValueType {
	ret := _m.Called()

	var r0 values.PlcValueType
	if rf, ok := ret.Get(0).(func() values.PlcValueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(values.PlcValueType)
	}

	return r0
}

// MockPlcTag_GetValueType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValueType'
type MockPlcTag_GetValueType_Call struct {
	*mock.Call
}

// GetValueType is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetValueType() *MockPlcTag_GetValueType_Call {
	return &MockPlcTag_GetValueType_Call{Call: _e.mock.On("GetValueType")}
}

func (_c *MockPlcTag_GetValueType_Call) Run(run func()) *MockPlcTag_GetValueType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetValueType_Call) Return(_a0 values.PlcValueType) *MockPlcTag_GetValueType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetValueType_Call) RunAndReturn(run func() values.PlcValueType) *MockPlcTag_GetValueType_Call {
	_c.Call.Return(run)
	return _c
}

// Serialize provides a mock function with given fields:
func (_m *MockPlcTag) Serialize() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPlcTag_Serialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Serialize'
type MockPlcTag_Serialize_Call struct {
	*mock.Call
}

// Serialize is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) Serialize() *MockPlcTag_Serialize_Call {
	return &MockPlcTag_Serialize_Call{Call: _e.mock.On("Serialize")}
}

func (_c *MockPlcTag_Serialize_Call) Run(run func()) *MockPlcTag_Serialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_Serialize_Call) Return(_a0 []byte, _a1 error) *MockPlcTag_Serialize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPlcTag_Serialize_Call) RunAndReturn(run func() ([]byte, error)) *MockPlcTag_Serialize_Call {
	_c.Call.Return(run)
	return _c
}

// SerializeWithWriteBuffer provides a mock function with given fields: ctx, writeBuffer
func (_m *MockPlcTag) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	ret := _m.Called(ctx, writeBuffer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, utils.WriteBuffer) error); ok {
		r0 = rf(ctx, writeBuffer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcTag_SerializeWithWriteBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SerializeWithWriteBuffer'
type MockPlcTag_SerializeWithWriteBuffer_Call struct {
	*mock.Call
}

// SerializeWithWriteBuffer is a helper method to define mock.On call
//   - ctx context.Context
//   - writeBuffer utils.WriteBuffer
func (_e *MockPlcTag_Expecter) SerializeWithWriteBuffer(ctx interface{}, writeBuffer interface{}) *MockPlcTag_SerializeWithWriteBuffer_Call {
	return &MockPlcTag_SerializeWithWriteBuffer_Call{Call: _e.mock.On("SerializeWithWriteBuffer", ctx, writeBuffer)}
}

func (_c *MockPlcTag_SerializeWithWriteBuffer_Call) Run(run func(ctx context.Context, writeBuffer utils.WriteBuffer)) *MockPlcTag_SerializeWithWriteBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(utils.WriteBuffer))
	})
	return _c
}

func (_c *MockPlcTag_SerializeWithWriteBuffer_Call) Return(_a0 error) *MockPlcTag_SerializeWithWriteBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_SerializeWithWriteBuffer_Call) RunAndReturn(run func(context.Context, utils.WriteBuffer) error) *MockPlcTag_SerializeWithWriteBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcTag) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcTag_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcTag_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) String() *MockPlcTag_String_Call {
	return &MockPlcTag_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcTag_String_Call) Run(run func()) *MockPlcTag_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_String_Call) Return(_a0 string) *MockPlcTag_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_String_Call) RunAndReturn(run func() string) *MockPlcTag_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcTag creates a new instance of MockPlcTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcTag(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcTag {
	mock := &MockPlcTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
