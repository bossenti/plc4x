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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPriorityArray is the data-structure of this message
type BACnetPriorityArray struct {
	PriorityValue01 *BACnetPriorityValue
	PriorityValue02 *BACnetPriorityValue
	PriorityValue03 *BACnetPriorityValue
	PriorityValue04 *BACnetPriorityValue
	PriorityValue05 *BACnetPriorityValue
	PriorityValue06 *BACnetPriorityValue
	PriorityValue07 *BACnetPriorityValue
	PriorityValue08 *BACnetPriorityValue
	PriorityValue09 *BACnetPriorityValue
	PriorityValue10 *BACnetPriorityValue
	PriorityValue11 *BACnetPriorityValue
	PriorityValue12 *BACnetPriorityValue
	PriorityValue13 *BACnetPriorityValue
	PriorityValue14 *BACnetPriorityValue
	PriorityValue15 *BACnetPriorityValue
	PriorityValue16 *BACnetPriorityValue

	// Arguments.
	ObjectType BACnetObjectType
}

// IBACnetPriorityArray is the corresponding interface of BACnetPriorityArray
type IBACnetPriorityArray interface {
	// GetPriorityValue01 returns PriorityValue01 (property field)
	GetPriorityValue01() *BACnetPriorityValue
	// GetPriorityValue02 returns PriorityValue02 (property field)
	GetPriorityValue02() *BACnetPriorityValue
	// GetPriorityValue03 returns PriorityValue03 (property field)
	GetPriorityValue03() *BACnetPriorityValue
	// GetPriorityValue04 returns PriorityValue04 (property field)
	GetPriorityValue04() *BACnetPriorityValue
	// GetPriorityValue05 returns PriorityValue05 (property field)
	GetPriorityValue05() *BACnetPriorityValue
	// GetPriorityValue06 returns PriorityValue06 (property field)
	GetPriorityValue06() *BACnetPriorityValue
	// GetPriorityValue07 returns PriorityValue07 (property field)
	GetPriorityValue07() *BACnetPriorityValue
	// GetPriorityValue08 returns PriorityValue08 (property field)
	GetPriorityValue08() *BACnetPriorityValue
	// GetPriorityValue09 returns PriorityValue09 (property field)
	GetPriorityValue09() *BACnetPriorityValue
	// GetPriorityValue10 returns PriorityValue10 (property field)
	GetPriorityValue10() *BACnetPriorityValue
	// GetPriorityValue11 returns PriorityValue11 (property field)
	GetPriorityValue11() *BACnetPriorityValue
	// GetPriorityValue12 returns PriorityValue12 (property field)
	GetPriorityValue12() *BACnetPriorityValue
	// GetPriorityValue13 returns PriorityValue13 (property field)
	GetPriorityValue13() *BACnetPriorityValue
	// GetPriorityValue14 returns PriorityValue14 (property field)
	GetPriorityValue14() *BACnetPriorityValue
	// GetPriorityValue15 returns PriorityValue15 (property field)
	GetPriorityValue15() *BACnetPriorityValue
	// GetPriorityValue16 returns PriorityValue16 (property field)
	GetPriorityValue16() *BACnetPriorityValue
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPriorityArray) GetPriorityValue01() *BACnetPriorityValue {
	return m.PriorityValue01
}

func (m *BACnetPriorityArray) GetPriorityValue02() *BACnetPriorityValue {
	return m.PriorityValue02
}

func (m *BACnetPriorityArray) GetPriorityValue03() *BACnetPriorityValue {
	return m.PriorityValue03
}

func (m *BACnetPriorityArray) GetPriorityValue04() *BACnetPriorityValue {
	return m.PriorityValue04
}

func (m *BACnetPriorityArray) GetPriorityValue05() *BACnetPriorityValue {
	return m.PriorityValue05
}

func (m *BACnetPriorityArray) GetPriorityValue06() *BACnetPriorityValue {
	return m.PriorityValue06
}

func (m *BACnetPriorityArray) GetPriorityValue07() *BACnetPriorityValue {
	return m.PriorityValue07
}

func (m *BACnetPriorityArray) GetPriorityValue08() *BACnetPriorityValue {
	return m.PriorityValue08
}

func (m *BACnetPriorityArray) GetPriorityValue09() *BACnetPriorityValue {
	return m.PriorityValue09
}

func (m *BACnetPriorityArray) GetPriorityValue10() *BACnetPriorityValue {
	return m.PriorityValue10
}

func (m *BACnetPriorityArray) GetPriorityValue11() *BACnetPriorityValue {
	return m.PriorityValue11
}

func (m *BACnetPriorityArray) GetPriorityValue12() *BACnetPriorityValue {
	return m.PriorityValue12
}

func (m *BACnetPriorityArray) GetPriorityValue13() *BACnetPriorityValue {
	return m.PriorityValue13
}

func (m *BACnetPriorityArray) GetPriorityValue14() *BACnetPriorityValue {
	return m.PriorityValue14
}

func (m *BACnetPriorityArray) GetPriorityValue15() *BACnetPriorityValue {
	return m.PriorityValue15
}

func (m *BACnetPriorityArray) GetPriorityValue16() *BACnetPriorityValue {
	return m.PriorityValue16
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityArray factory function for BACnetPriorityArray
func NewBACnetPriorityArray(priorityValue01 *BACnetPriorityValue, priorityValue02 *BACnetPriorityValue, priorityValue03 *BACnetPriorityValue, priorityValue04 *BACnetPriorityValue, priorityValue05 *BACnetPriorityValue, priorityValue06 *BACnetPriorityValue, priorityValue07 *BACnetPriorityValue, priorityValue08 *BACnetPriorityValue, priorityValue09 *BACnetPriorityValue, priorityValue10 *BACnetPriorityValue, priorityValue11 *BACnetPriorityValue, priorityValue12 *BACnetPriorityValue, priorityValue13 *BACnetPriorityValue, priorityValue14 *BACnetPriorityValue, priorityValue15 *BACnetPriorityValue, priorityValue16 *BACnetPriorityValue, objectType BACnetObjectType) *BACnetPriorityArray {
	return &BACnetPriorityArray{PriorityValue01: priorityValue01, PriorityValue02: priorityValue02, PriorityValue03: priorityValue03, PriorityValue04: priorityValue04, PriorityValue05: priorityValue05, PriorityValue06: priorityValue06, PriorityValue07: priorityValue07, PriorityValue08: priorityValue08, PriorityValue09: priorityValue09, PriorityValue10: priorityValue10, PriorityValue11: priorityValue11, PriorityValue12: priorityValue12, PriorityValue13: priorityValue13, PriorityValue14: priorityValue14, PriorityValue15: priorityValue15, PriorityValue16: priorityValue16, ObjectType: objectType}
}

func CastBACnetPriorityArray(structType interface{}) *BACnetPriorityArray {
	if casted, ok := structType.(BACnetPriorityArray); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPriorityArray); ok {
		return casted
	}
	return nil
}

func (m *BACnetPriorityArray) GetTypeName() string {
	return "BACnetPriorityArray"
}

func (m *BACnetPriorityArray) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPriorityArray) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (priorityValue01)
	lengthInBits += m.PriorityValue01.GetLengthInBits()

	// Simple field (priorityValue02)
	lengthInBits += m.PriorityValue02.GetLengthInBits()

	// Simple field (priorityValue03)
	lengthInBits += m.PriorityValue03.GetLengthInBits()

	// Simple field (priorityValue04)
	lengthInBits += m.PriorityValue04.GetLengthInBits()

	// Simple field (priorityValue05)
	lengthInBits += m.PriorityValue05.GetLengthInBits()

	// Simple field (priorityValue06)
	lengthInBits += m.PriorityValue06.GetLengthInBits()

	// Simple field (priorityValue07)
	lengthInBits += m.PriorityValue07.GetLengthInBits()

	// Simple field (priorityValue08)
	lengthInBits += m.PriorityValue08.GetLengthInBits()

	// Simple field (priorityValue09)
	lengthInBits += m.PriorityValue09.GetLengthInBits()

	// Simple field (priorityValue10)
	lengthInBits += m.PriorityValue10.GetLengthInBits()

	// Simple field (priorityValue11)
	lengthInBits += m.PriorityValue11.GetLengthInBits()

	// Simple field (priorityValue12)
	lengthInBits += m.PriorityValue12.GetLengthInBits()

	// Simple field (priorityValue13)
	lengthInBits += m.PriorityValue13.GetLengthInBits()

	// Simple field (priorityValue14)
	lengthInBits += m.PriorityValue14.GetLengthInBits()

	// Simple field (priorityValue15)
	lengthInBits += m.PriorityValue15.GetLengthInBits()

	// Simple field (priorityValue16)
	lengthInBits += m.PriorityValue16.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPriorityArray) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityArrayParse(readBuffer utils.ReadBuffer, objectType BACnetObjectType) (*BACnetPriorityArray, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityArray"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (priorityValue01)
	if pullErr := readBuffer.PullContext("priorityValue01"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue01, _priorityValue01Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue01Err != nil {
		return nil, errors.Wrap(_priorityValue01Err, "Error parsing 'priorityValue01' field")
	}
	priorityValue01 := CastBACnetPriorityValue(_priorityValue01)
	if closeErr := readBuffer.CloseContext("priorityValue01"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue02)
	if pullErr := readBuffer.PullContext("priorityValue02"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue02, _priorityValue02Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue02Err != nil {
		return nil, errors.Wrap(_priorityValue02Err, "Error parsing 'priorityValue02' field")
	}
	priorityValue02 := CastBACnetPriorityValue(_priorityValue02)
	if closeErr := readBuffer.CloseContext("priorityValue02"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue03)
	if pullErr := readBuffer.PullContext("priorityValue03"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue03, _priorityValue03Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue03Err != nil {
		return nil, errors.Wrap(_priorityValue03Err, "Error parsing 'priorityValue03' field")
	}
	priorityValue03 := CastBACnetPriorityValue(_priorityValue03)
	if closeErr := readBuffer.CloseContext("priorityValue03"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue04)
	if pullErr := readBuffer.PullContext("priorityValue04"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue04, _priorityValue04Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue04Err != nil {
		return nil, errors.Wrap(_priorityValue04Err, "Error parsing 'priorityValue04' field")
	}
	priorityValue04 := CastBACnetPriorityValue(_priorityValue04)
	if closeErr := readBuffer.CloseContext("priorityValue04"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue05)
	if pullErr := readBuffer.PullContext("priorityValue05"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue05, _priorityValue05Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue05Err != nil {
		return nil, errors.Wrap(_priorityValue05Err, "Error parsing 'priorityValue05' field")
	}
	priorityValue05 := CastBACnetPriorityValue(_priorityValue05)
	if closeErr := readBuffer.CloseContext("priorityValue05"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue06)
	if pullErr := readBuffer.PullContext("priorityValue06"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue06, _priorityValue06Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue06Err != nil {
		return nil, errors.Wrap(_priorityValue06Err, "Error parsing 'priorityValue06' field")
	}
	priorityValue06 := CastBACnetPriorityValue(_priorityValue06)
	if closeErr := readBuffer.CloseContext("priorityValue06"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue07)
	if pullErr := readBuffer.PullContext("priorityValue07"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue07, _priorityValue07Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue07Err != nil {
		return nil, errors.Wrap(_priorityValue07Err, "Error parsing 'priorityValue07' field")
	}
	priorityValue07 := CastBACnetPriorityValue(_priorityValue07)
	if closeErr := readBuffer.CloseContext("priorityValue07"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue08)
	if pullErr := readBuffer.PullContext("priorityValue08"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue08, _priorityValue08Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue08Err != nil {
		return nil, errors.Wrap(_priorityValue08Err, "Error parsing 'priorityValue08' field")
	}
	priorityValue08 := CastBACnetPriorityValue(_priorityValue08)
	if closeErr := readBuffer.CloseContext("priorityValue08"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue09)
	if pullErr := readBuffer.PullContext("priorityValue09"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue09, _priorityValue09Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue09Err != nil {
		return nil, errors.Wrap(_priorityValue09Err, "Error parsing 'priorityValue09' field")
	}
	priorityValue09 := CastBACnetPriorityValue(_priorityValue09)
	if closeErr := readBuffer.CloseContext("priorityValue09"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue10)
	if pullErr := readBuffer.PullContext("priorityValue10"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue10, _priorityValue10Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue10Err != nil {
		return nil, errors.Wrap(_priorityValue10Err, "Error parsing 'priorityValue10' field")
	}
	priorityValue10 := CastBACnetPriorityValue(_priorityValue10)
	if closeErr := readBuffer.CloseContext("priorityValue10"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue11)
	if pullErr := readBuffer.PullContext("priorityValue11"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue11, _priorityValue11Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue11Err != nil {
		return nil, errors.Wrap(_priorityValue11Err, "Error parsing 'priorityValue11' field")
	}
	priorityValue11 := CastBACnetPriorityValue(_priorityValue11)
	if closeErr := readBuffer.CloseContext("priorityValue11"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue12)
	if pullErr := readBuffer.PullContext("priorityValue12"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue12, _priorityValue12Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue12Err != nil {
		return nil, errors.Wrap(_priorityValue12Err, "Error parsing 'priorityValue12' field")
	}
	priorityValue12 := CastBACnetPriorityValue(_priorityValue12)
	if closeErr := readBuffer.CloseContext("priorityValue12"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue13)
	if pullErr := readBuffer.PullContext("priorityValue13"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue13, _priorityValue13Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue13Err != nil {
		return nil, errors.Wrap(_priorityValue13Err, "Error parsing 'priorityValue13' field")
	}
	priorityValue13 := CastBACnetPriorityValue(_priorityValue13)
	if closeErr := readBuffer.CloseContext("priorityValue13"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue14)
	if pullErr := readBuffer.PullContext("priorityValue14"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue14, _priorityValue14Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue14Err != nil {
		return nil, errors.Wrap(_priorityValue14Err, "Error parsing 'priorityValue14' field")
	}
	priorityValue14 := CastBACnetPriorityValue(_priorityValue14)
	if closeErr := readBuffer.CloseContext("priorityValue14"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue15)
	if pullErr := readBuffer.PullContext("priorityValue15"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue15, _priorityValue15Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue15Err != nil {
		return nil, errors.Wrap(_priorityValue15Err, "Error parsing 'priorityValue15' field")
	}
	priorityValue15 := CastBACnetPriorityValue(_priorityValue15)
	if closeErr := readBuffer.CloseContext("priorityValue15"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priorityValue16)
	if pullErr := readBuffer.PullContext("priorityValue16"); pullErr != nil {
		return nil, pullErr
	}
	_priorityValue16, _priorityValue16Err := BACnetPriorityValueParse(readBuffer, BACnetObjectType(objectType))
	if _priorityValue16Err != nil {
		return nil, errors.Wrap(_priorityValue16Err, "Error parsing 'priorityValue16' field")
	}
	priorityValue16 := CastBACnetPriorityValue(_priorityValue16)
	if closeErr := readBuffer.CloseContext("priorityValue16"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityArray"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetPriorityArray(priorityValue01, priorityValue02, priorityValue03, priorityValue04, priorityValue05, priorityValue06, priorityValue07, priorityValue08, priorityValue09, priorityValue10, priorityValue11, priorityValue12, priorityValue13, priorityValue14, priorityValue15, priorityValue16, objectType), nil
}

func (m *BACnetPriorityArray) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPriorityArray"); pushErr != nil {
		return pushErr
	}

	// Simple Field (priorityValue01)
	if pushErr := writeBuffer.PushContext("priorityValue01"); pushErr != nil {
		return pushErr
	}
	_priorityValue01Err := m.PriorityValue01.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue01"); popErr != nil {
		return popErr
	}
	if _priorityValue01Err != nil {
		return errors.Wrap(_priorityValue01Err, "Error serializing 'priorityValue01' field")
	}

	// Simple Field (priorityValue02)
	if pushErr := writeBuffer.PushContext("priorityValue02"); pushErr != nil {
		return pushErr
	}
	_priorityValue02Err := m.PriorityValue02.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue02"); popErr != nil {
		return popErr
	}
	if _priorityValue02Err != nil {
		return errors.Wrap(_priorityValue02Err, "Error serializing 'priorityValue02' field")
	}

	// Simple Field (priorityValue03)
	if pushErr := writeBuffer.PushContext("priorityValue03"); pushErr != nil {
		return pushErr
	}
	_priorityValue03Err := m.PriorityValue03.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue03"); popErr != nil {
		return popErr
	}
	if _priorityValue03Err != nil {
		return errors.Wrap(_priorityValue03Err, "Error serializing 'priorityValue03' field")
	}

	// Simple Field (priorityValue04)
	if pushErr := writeBuffer.PushContext("priorityValue04"); pushErr != nil {
		return pushErr
	}
	_priorityValue04Err := m.PriorityValue04.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue04"); popErr != nil {
		return popErr
	}
	if _priorityValue04Err != nil {
		return errors.Wrap(_priorityValue04Err, "Error serializing 'priorityValue04' field")
	}

	// Simple Field (priorityValue05)
	if pushErr := writeBuffer.PushContext("priorityValue05"); pushErr != nil {
		return pushErr
	}
	_priorityValue05Err := m.PriorityValue05.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue05"); popErr != nil {
		return popErr
	}
	if _priorityValue05Err != nil {
		return errors.Wrap(_priorityValue05Err, "Error serializing 'priorityValue05' field")
	}

	// Simple Field (priorityValue06)
	if pushErr := writeBuffer.PushContext("priorityValue06"); pushErr != nil {
		return pushErr
	}
	_priorityValue06Err := m.PriorityValue06.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue06"); popErr != nil {
		return popErr
	}
	if _priorityValue06Err != nil {
		return errors.Wrap(_priorityValue06Err, "Error serializing 'priorityValue06' field")
	}

	// Simple Field (priorityValue07)
	if pushErr := writeBuffer.PushContext("priorityValue07"); pushErr != nil {
		return pushErr
	}
	_priorityValue07Err := m.PriorityValue07.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue07"); popErr != nil {
		return popErr
	}
	if _priorityValue07Err != nil {
		return errors.Wrap(_priorityValue07Err, "Error serializing 'priorityValue07' field")
	}

	// Simple Field (priorityValue08)
	if pushErr := writeBuffer.PushContext("priorityValue08"); pushErr != nil {
		return pushErr
	}
	_priorityValue08Err := m.PriorityValue08.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue08"); popErr != nil {
		return popErr
	}
	if _priorityValue08Err != nil {
		return errors.Wrap(_priorityValue08Err, "Error serializing 'priorityValue08' field")
	}

	// Simple Field (priorityValue09)
	if pushErr := writeBuffer.PushContext("priorityValue09"); pushErr != nil {
		return pushErr
	}
	_priorityValue09Err := m.PriorityValue09.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue09"); popErr != nil {
		return popErr
	}
	if _priorityValue09Err != nil {
		return errors.Wrap(_priorityValue09Err, "Error serializing 'priorityValue09' field")
	}

	// Simple Field (priorityValue10)
	if pushErr := writeBuffer.PushContext("priorityValue10"); pushErr != nil {
		return pushErr
	}
	_priorityValue10Err := m.PriorityValue10.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue10"); popErr != nil {
		return popErr
	}
	if _priorityValue10Err != nil {
		return errors.Wrap(_priorityValue10Err, "Error serializing 'priorityValue10' field")
	}

	// Simple Field (priorityValue11)
	if pushErr := writeBuffer.PushContext("priorityValue11"); pushErr != nil {
		return pushErr
	}
	_priorityValue11Err := m.PriorityValue11.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue11"); popErr != nil {
		return popErr
	}
	if _priorityValue11Err != nil {
		return errors.Wrap(_priorityValue11Err, "Error serializing 'priorityValue11' field")
	}

	// Simple Field (priorityValue12)
	if pushErr := writeBuffer.PushContext("priorityValue12"); pushErr != nil {
		return pushErr
	}
	_priorityValue12Err := m.PriorityValue12.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue12"); popErr != nil {
		return popErr
	}
	if _priorityValue12Err != nil {
		return errors.Wrap(_priorityValue12Err, "Error serializing 'priorityValue12' field")
	}

	// Simple Field (priorityValue13)
	if pushErr := writeBuffer.PushContext("priorityValue13"); pushErr != nil {
		return pushErr
	}
	_priorityValue13Err := m.PriorityValue13.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue13"); popErr != nil {
		return popErr
	}
	if _priorityValue13Err != nil {
		return errors.Wrap(_priorityValue13Err, "Error serializing 'priorityValue13' field")
	}

	// Simple Field (priorityValue14)
	if pushErr := writeBuffer.PushContext("priorityValue14"); pushErr != nil {
		return pushErr
	}
	_priorityValue14Err := m.PriorityValue14.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue14"); popErr != nil {
		return popErr
	}
	if _priorityValue14Err != nil {
		return errors.Wrap(_priorityValue14Err, "Error serializing 'priorityValue14' field")
	}

	// Simple Field (priorityValue15)
	if pushErr := writeBuffer.PushContext("priorityValue15"); pushErr != nil {
		return pushErr
	}
	_priorityValue15Err := m.PriorityValue15.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue15"); popErr != nil {
		return popErr
	}
	if _priorityValue15Err != nil {
		return errors.Wrap(_priorityValue15Err, "Error serializing 'priorityValue15' field")
	}

	// Simple Field (priorityValue16)
	if pushErr := writeBuffer.PushContext("priorityValue16"); pushErr != nil {
		return pushErr
	}
	_priorityValue16Err := m.PriorityValue16.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityValue16"); popErr != nil {
		return popErr
	}
	if _priorityValue16Err != nil {
		return errors.Wrap(_priorityValue16Err, "Error serializing 'priorityValue16' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPriorityArray"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetPriorityArray) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
