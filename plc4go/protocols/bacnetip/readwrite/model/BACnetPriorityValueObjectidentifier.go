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

// BACnetPriorityValueObjectidentifier is the data-structure of this message
type BACnetPriorityValueObjectidentifier struct {
	*BACnetPriorityValue
	ObjectidentifierValue *BACnetApplicationTagObjectIdentifier

	// Arguments.
	ObjectType BACnetObjectType
}

// IBACnetPriorityValueObjectidentifier is the corresponding interface of BACnetPriorityValueObjectidentifier
type IBACnetPriorityValueObjectidentifier interface {
	IBACnetPriorityValue
	// GetObjectidentifierValue returns ObjectidentifierValue (property field)
	GetObjectidentifierValue() *BACnetApplicationTagObjectIdentifier
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

func (m *BACnetPriorityValueObjectidentifier) InitializeParent(parent *BACnetPriorityValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPriorityValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPriorityValueObjectidentifier) GetParent() *BACnetPriorityValue {
	return m.BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPriorityValueObjectidentifier) GetObjectidentifierValue() *BACnetApplicationTagObjectIdentifier {
	return m.ObjectidentifierValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueObjectidentifier factory function for BACnetPriorityValueObjectidentifier
func NewBACnetPriorityValueObjectidentifier(objectidentifierValue *BACnetApplicationTagObjectIdentifier, peekedTagHeader *BACnetTagHeader, objectType BACnetObjectType) *BACnetPriorityValueObjectidentifier {
	_result := &BACnetPriorityValueObjectidentifier{
		ObjectidentifierValue: objectidentifierValue,
		BACnetPriorityValue:   NewBACnetPriorityValue(peekedTagHeader, objectType),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPriorityValueObjectidentifier(structType interface{}) *BACnetPriorityValueObjectidentifier {
	if casted, ok := structType.(BACnetPriorityValueObjectidentifier); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPriorityValueObjectidentifier); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPriorityValue); ok {
		return CastBACnetPriorityValueObjectidentifier(casted.Child)
	}
	if casted, ok := structType.(*BACnetPriorityValue); ok {
		return CastBACnetPriorityValueObjectidentifier(casted.Child)
	}
	return nil
}

func (m *BACnetPriorityValueObjectidentifier) GetTypeName() string {
	return "BACnetPriorityValueObjectidentifier"
}

func (m *BACnetPriorityValueObjectidentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPriorityValueObjectidentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectidentifierValue)
	lengthInBits += m.ObjectidentifierValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPriorityValueObjectidentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueObjectidentifierParse(readBuffer utils.ReadBuffer, objectType BACnetObjectType) (*BACnetPriorityValueObjectidentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueObjectidentifier"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectidentifierValue)
	if pullErr := readBuffer.PullContext("objectidentifierValue"); pullErr != nil {
		return nil, pullErr
	}
	_objectidentifierValue, _objectidentifierValueErr := BACnetApplicationTagParse(readBuffer)
	if _objectidentifierValueErr != nil {
		return nil, errors.Wrap(_objectidentifierValueErr, "Error parsing 'objectidentifierValue' field")
	}
	objectidentifierValue := CastBACnetApplicationTagObjectIdentifier(_objectidentifierValue)
	if closeErr := readBuffer.CloseContext("objectidentifierValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueObjectidentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPriorityValueObjectidentifier{
		ObjectidentifierValue: CastBACnetApplicationTagObjectIdentifier(objectidentifierValue),
		BACnetPriorityValue:   &BACnetPriorityValue{},
	}
	_child.BACnetPriorityValue.Child = _child
	return _child, nil
}

func (m *BACnetPriorityValueObjectidentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueObjectidentifier"); pushErr != nil {
			return pushErr
		}

		// Simple Field (objectidentifierValue)
		if pushErr := writeBuffer.PushContext("objectidentifierValue"); pushErr != nil {
			return pushErr
		}
		_objectidentifierValueErr := m.ObjectidentifierValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("objectidentifierValue"); popErr != nil {
			return popErr
		}
		if _objectidentifierValueErr != nil {
			return errors.Wrap(_objectidentifierValueErr, "Error serializing 'objectidentifierValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueObjectidentifier"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPriorityValueObjectidentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
