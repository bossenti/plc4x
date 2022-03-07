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
type BACnetNotificationParametersChangeOfValueNewValueChangedBits struct {
	*BACnetNotificationParametersChangeOfValueNewValue
	ChangedBits *BACnetContextTagBitString

	// Arguments.
	TagNumber uint8
}

// The corresponding interface
type IBACnetNotificationParametersChangeOfValueNewValueChangedBits interface {
	IBACnetNotificationParametersChangeOfValueNewValue
	// GetChangedBits returns ChangedBits (property field)
	GetChangedBits() *BACnetContextTagBitString
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
///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedBits) InitializeParent(parent *BACnetNotificationParametersChangeOfValueNewValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParametersChangeOfValueNewValue.OpeningTag = openingTag
	m.BACnetNotificationParametersChangeOfValueNewValue.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParametersChangeOfValueNewValue.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetParent() *BACnetNotificationParametersChangeOfValueNewValue {
	return m.BACnetNotificationParametersChangeOfValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetChangedBits() *BACnetContextTagBitString {
	return m.ChangedBits
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfValueNewValueChangedBits factory function for BACnetNotificationParametersChangeOfValueNewValueChangedBits
func NewBACnetNotificationParametersChangeOfValueNewValueChangedBits(changedBits *BACnetContextTagBitString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetNotificationParametersChangeOfValueNewValue {
	child := &BACnetNotificationParametersChangeOfValueNewValueChangedBits{
		ChangedBits: changedBits,
		BACnetNotificationParametersChangeOfValueNewValue: NewBACnetNotificationParametersChangeOfValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	child.Child = child
	return child.BACnetNotificationParametersChangeOfValueNewValue
}

func CastBACnetNotificationParametersChangeOfValueNewValueChangedBits(structType interface{}) *BACnetNotificationParametersChangeOfValueNewValueChangedBits {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfValueNewValueChangedBits); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfValueNewValueChangedBits); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParametersChangeOfValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfValueNewValueChangedBits(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfValueNewValueChangedBits(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfValueNewValueChangedBits"
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (changedBits)
	lengthInBits += m.ChangedBits.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfValueNewValueChangedBitsParse(readBuffer utils.ReadBuffer, tagNumber uint8, peekedTagNumber uint8) (*BACnetNotificationParametersChangeOfValueNewValueChangedBits, error) {
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (changedBits)
	if pullErr := readBuffer.PullContext("changedBits"); pullErr != nil {
		return nil, pullErr
	}
	_changedBits, _changedBitsErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BIT_STRING))
	if _changedBitsErr != nil {
		return nil, errors.Wrap(_changedBitsErr, "Error parsing 'changedBits' field")
	}
	changedBits := CastBACnetContextTagBitString(_changedBits)
	if closeErr := readBuffer.CloseContext("changedBits"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfValueNewValueChangedBits{
		ChangedBits: CastBACnetContextTagBitString(changedBits),
		BACnetNotificationParametersChangeOfValueNewValue: &BACnetNotificationParametersChangeOfValueNewValue{},
	}
	_child.BACnetNotificationParametersChangeOfValueNewValue.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedBits) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); pushErr != nil {
			return pushErr
		}

		// Simple Field (changedBits)
		if pushErr := writeBuffer.PushContext("changedBits"); pushErr != nil {
			return pushErr
		}
		_changedBitsErr := m.ChangedBits.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("changedBits"); popErr != nil {
			return popErr
		}
		if _changedBitsErr != nil {
			return errors.Wrap(_changedBitsErr, "Error serializing 'changedBits' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedBits) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
