/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type S7ParameterUserDataItemCPUFunctions struct {
	*S7ParameterUserDataItem
	Method                  uint8
	CpuFunctionType         uint8
	CpuFunctionGroup        uint8
	CpuSubfunction          uint8
	SequenceNumber          uint8
	DataUnitReferenceNumber *uint8
	LastDataUnit            *uint8
	ErrorCode               *uint16
}

// The corresponding interface
type IS7ParameterUserDataItemCPUFunctions interface {
	IS7ParameterUserDataItem
	// GetMethod returns Method (property field)
	GetMethod() uint8
	// GetCpuFunctionType returns CpuFunctionType (property field)
	GetCpuFunctionType() uint8
	// GetCpuFunctionGroup returns CpuFunctionGroup (property field)
	GetCpuFunctionGroup() uint8
	// GetCpuSubfunction returns CpuSubfunction (property field)
	GetCpuSubfunction() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
	// GetDataUnitReferenceNumber returns DataUnitReferenceNumber (property field)
	GetDataUnitReferenceNumber() *uint8
	// GetLastDataUnit returns LastDataUnit (property field)
	GetLastDataUnit() *uint8
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() *uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////
func (m *S7ParameterUserDataItemCPUFunctions) GetItemType() uint8 {
	return 0x12
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7ParameterUserDataItemCPUFunctions) InitializeParent(parent *S7ParameterUserDataItem) {}

func (m *S7ParameterUserDataItemCPUFunctions) GetParent() *S7ParameterUserDataItem {
	return m.S7ParameterUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *S7ParameterUserDataItemCPUFunctions) GetMethod() uint8 {
	return m.Method
}

func (m *S7ParameterUserDataItemCPUFunctions) GetCpuFunctionType() uint8 {
	return m.CpuFunctionType
}

func (m *S7ParameterUserDataItemCPUFunctions) GetCpuFunctionGroup() uint8 {
	return m.CpuFunctionGroup
}

func (m *S7ParameterUserDataItemCPUFunctions) GetCpuSubfunction() uint8 {
	return m.CpuSubfunction
}

func (m *S7ParameterUserDataItemCPUFunctions) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

func (m *S7ParameterUserDataItemCPUFunctions) GetDataUnitReferenceNumber() *uint8 {
	return m.DataUnitReferenceNumber
}

func (m *S7ParameterUserDataItemCPUFunctions) GetLastDataUnit() *uint8 {
	return m.LastDataUnit
}

func (m *S7ParameterUserDataItemCPUFunctions) GetErrorCode() *uint16 {
	return m.ErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterUserDataItemCPUFunctions factory function for S7ParameterUserDataItemCPUFunctions
func NewS7ParameterUserDataItemCPUFunctions(method uint8, cpuFunctionType uint8, cpuFunctionGroup uint8, cpuSubfunction uint8, sequenceNumber uint8, dataUnitReferenceNumber *uint8, lastDataUnit *uint8, errorCode *uint16) *S7ParameterUserDataItem {
	child := &S7ParameterUserDataItemCPUFunctions{
		Method:                  method,
		CpuFunctionType:         cpuFunctionType,
		CpuFunctionGroup:        cpuFunctionGroup,
		CpuSubfunction:          cpuSubfunction,
		SequenceNumber:          sequenceNumber,
		DataUnitReferenceNumber: dataUnitReferenceNumber,
		LastDataUnit:            lastDataUnit,
		ErrorCode:               errorCode,
		S7ParameterUserDataItem: NewS7ParameterUserDataItem(),
	}
	child.Child = child
	return child.S7ParameterUserDataItem
}

func CastS7ParameterUserDataItemCPUFunctions(structType interface{}) *S7ParameterUserDataItemCPUFunctions {
	if casted, ok := structType.(S7ParameterUserDataItemCPUFunctions); ok {
		return &casted
	}
	if casted, ok := structType.(*S7ParameterUserDataItemCPUFunctions); ok {
		return casted
	}
	if casted, ok := structType.(S7ParameterUserDataItem); ok {
		return CastS7ParameterUserDataItemCPUFunctions(casted.Child)
	}
	if casted, ok := structType.(*S7ParameterUserDataItem); ok {
		return CastS7ParameterUserDataItemCPUFunctions(casted.Child)
	}
	return nil
}

func (m *S7ParameterUserDataItemCPUFunctions) GetTypeName() string {
	return "S7ParameterUserDataItemCPUFunctions"
}

func (m *S7ParameterUserDataItemCPUFunctions) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7ParameterUserDataItemCPUFunctions) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (itemLength)
	lengthInBits += 8

	// Simple field (method)
	lengthInBits += 8

	// Simple field (cpuFunctionType)
	lengthInBits += 4

	// Simple field (cpuFunctionGroup)
	lengthInBits += 4

	// Simple field (cpuSubfunction)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	// Optional Field (dataUnitReferenceNumber)
	if m.DataUnitReferenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (lastDataUnit)
	if m.LastDataUnit != nil {
		lengthInBits += 8
	}

	// Optional Field (errorCode)
	if m.ErrorCode != nil {
		lengthInBits += 16
	}

	return lengthInBits
}

func (m *S7ParameterUserDataItemCPUFunctions) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7ParameterUserDataItemCPUFunctionsParse(readBuffer utils.ReadBuffer) (*S7ParameterUserDataItemCPUFunctions, error) {
	if pullErr := readBuffer.PullContext("S7ParameterUserDataItemCPUFunctions"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Implicit Field (itemLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	itemLength, _itemLengthErr := readBuffer.ReadUint8("itemLength", 8)
	_ = itemLength
	if _itemLengthErr != nil {
		return nil, errors.Wrap(_itemLengthErr, "Error parsing 'itemLength' field")
	}

	// Simple Field (method)
	_method, _methodErr := readBuffer.ReadUint8("method", 8)
	if _methodErr != nil {
		return nil, errors.Wrap(_methodErr, "Error parsing 'method' field")
	}
	method := _method

	// Simple Field (cpuFunctionType)
	_cpuFunctionType, _cpuFunctionTypeErr := readBuffer.ReadUint8("cpuFunctionType", 4)
	if _cpuFunctionTypeErr != nil {
		return nil, errors.Wrap(_cpuFunctionTypeErr, "Error parsing 'cpuFunctionType' field")
	}
	cpuFunctionType := _cpuFunctionType

	// Simple Field (cpuFunctionGroup)
	_cpuFunctionGroup, _cpuFunctionGroupErr := readBuffer.ReadUint8("cpuFunctionGroup", 4)
	if _cpuFunctionGroupErr != nil {
		return nil, errors.Wrap(_cpuFunctionGroupErr, "Error parsing 'cpuFunctionGroup' field")
	}
	cpuFunctionGroup := _cpuFunctionGroup

	// Simple Field (cpuSubfunction)
	_cpuSubfunction, _cpuSubfunctionErr := readBuffer.ReadUint8("cpuSubfunction", 8)
	if _cpuSubfunctionErr != nil {
		return nil, errors.Wrap(_cpuSubfunctionErr, "Error parsing 'cpuSubfunction' field")
	}
	cpuSubfunction := _cpuSubfunction

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := readBuffer.ReadUint8("sequenceNumber", 8)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field")
	}
	sequenceNumber := _sequenceNumber

	// Optional Field (dataUnitReferenceNumber) (Can be skipped, if a given expression evaluates to false)
	var dataUnitReferenceNumber *uint8 = nil
	if bool((cpuFunctionType) == (8)) {
		_val, _err := readBuffer.ReadUint8("dataUnitReferenceNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'dataUnitReferenceNumber' field")
		}
		dataUnitReferenceNumber = &_val
	}

	// Optional Field (lastDataUnit) (Can be skipped, if a given expression evaluates to false)
	var lastDataUnit *uint8 = nil
	if bool((cpuFunctionType) == (8)) {
		_val, _err := readBuffer.ReadUint8("lastDataUnit", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'lastDataUnit' field")
		}
		lastDataUnit = &_val
	}

	// Optional Field (errorCode) (Can be skipped, if a given expression evaluates to false)
	var errorCode *uint16 = nil
	if bool((cpuFunctionType) == (8)) {
		_val, _err := readBuffer.ReadUint16("errorCode", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'errorCode' field")
		}
		errorCode = &_val
	}

	if closeErr := readBuffer.CloseContext("S7ParameterUserDataItemCPUFunctions"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7ParameterUserDataItemCPUFunctions{
		Method:                  method,
		CpuFunctionType:         cpuFunctionType,
		CpuFunctionGroup:        cpuFunctionGroup,
		CpuSubfunction:          cpuSubfunction,
		SequenceNumber:          sequenceNumber,
		DataUnitReferenceNumber: dataUnitReferenceNumber,
		LastDataUnit:            lastDataUnit,
		ErrorCode:               errorCode,
		S7ParameterUserDataItem: &S7ParameterUserDataItem{},
	}
	_child.S7ParameterUserDataItem.Child = _child
	return _child, nil
}

func (m *S7ParameterUserDataItemCPUFunctions) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterUserDataItemCPUFunctions"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		itemLength := uint8(uint8(uint8(m.GetLengthInBytes())) - uint8(uint8(2)))
		_itemLengthErr := writeBuffer.WriteUint8("itemLength", 8, (itemLength))
		if _itemLengthErr != nil {
			return errors.Wrap(_itemLengthErr, "Error serializing 'itemLength' field")
		}

		// Simple Field (method)
		method := uint8(m.Method)
		_methodErr := writeBuffer.WriteUint8("method", 8, (method))
		if _methodErr != nil {
			return errors.Wrap(_methodErr, "Error serializing 'method' field")
		}

		// Simple Field (cpuFunctionType)
		cpuFunctionType := uint8(m.CpuFunctionType)
		_cpuFunctionTypeErr := writeBuffer.WriteUint8("cpuFunctionType", 4, (cpuFunctionType))
		if _cpuFunctionTypeErr != nil {
			return errors.Wrap(_cpuFunctionTypeErr, "Error serializing 'cpuFunctionType' field")
		}

		// Simple Field (cpuFunctionGroup)
		cpuFunctionGroup := uint8(m.CpuFunctionGroup)
		_cpuFunctionGroupErr := writeBuffer.WriteUint8("cpuFunctionGroup", 4, (cpuFunctionGroup))
		if _cpuFunctionGroupErr != nil {
			return errors.Wrap(_cpuFunctionGroupErr, "Error serializing 'cpuFunctionGroup' field")
		}

		// Simple Field (cpuSubfunction)
		cpuSubfunction := uint8(m.CpuSubfunction)
		_cpuSubfunctionErr := writeBuffer.WriteUint8("cpuSubfunction", 8, (cpuSubfunction))
		if _cpuSubfunctionErr != nil {
			return errors.Wrap(_cpuSubfunctionErr, "Error serializing 'cpuSubfunction' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := uint8(m.SequenceNumber)
		_sequenceNumberErr := writeBuffer.WriteUint8("sequenceNumber", 8, (sequenceNumber))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		// Optional Field (dataUnitReferenceNumber) (Can be skipped, if the value is null)
		var dataUnitReferenceNumber *uint8 = nil
		if m.DataUnitReferenceNumber != nil {
			dataUnitReferenceNumber = m.DataUnitReferenceNumber
			_dataUnitReferenceNumberErr := writeBuffer.WriteUint8("dataUnitReferenceNumber", 8, *(dataUnitReferenceNumber))
			if _dataUnitReferenceNumberErr != nil {
				return errors.Wrap(_dataUnitReferenceNumberErr, "Error serializing 'dataUnitReferenceNumber' field")
			}
		}

		// Optional Field (lastDataUnit) (Can be skipped, if the value is null)
		var lastDataUnit *uint8 = nil
		if m.LastDataUnit != nil {
			lastDataUnit = m.LastDataUnit
			_lastDataUnitErr := writeBuffer.WriteUint8("lastDataUnit", 8, *(lastDataUnit))
			if _lastDataUnitErr != nil {
				return errors.Wrap(_lastDataUnitErr, "Error serializing 'lastDataUnit' field")
			}
		}

		// Optional Field (errorCode) (Can be skipped, if the value is null)
		var errorCode *uint16 = nil
		if m.ErrorCode != nil {
			errorCode = m.ErrorCode
			_errorCodeErr := writeBuffer.WriteUint16("errorCode", 16, *(errorCode))
			if _errorCodeErr != nil {
				return errors.Wrap(_errorCodeErr, "Error serializing 'errorCode' field")
			}
		}

		if popErr := writeBuffer.PopContext("S7ParameterUserDataItemCPUFunctions"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7ParameterUserDataItemCPUFunctions) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
