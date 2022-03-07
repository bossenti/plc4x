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
type BACnetContextTagObjectIdentifier struct {
	*BACnetContextTag
	Payload *BACnetTagPayloadObjectIdentifier

	// Arguments.
	TagNumberArgument        uint8
	IsNotOpeningOrClosingTag bool
}

// The corresponding interface
type IBACnetContextTagObjectIdentifier interface {
	IBACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() *BACnetTagPayloadObjectIdentifier
	// GetObjectType returns ObjectType (virtual field)
	GetObjectType() BACnetObjectType
	// GetInstanceNumber returns InstanceNumber (virtual field)
	GetInstanceNumber() uint32
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
func (m *BACnetContextTagObjectIdentifier) GetDataType() BACnetDataType {
	return BACnetDataType_BACNET_OBJECT_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetContextTagObjectIdentifier) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

func (m *BACnetContextTagObjectIdentifier) GetParent() *BACnetContextTag {
	return m.BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetContextTagObjectIdentifier) GetPayload() *BACnetTagPayloadObjectIdentifier {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////
func (m *BACnetContextTagObjectIdentifier) GetObjectType() BACnetObjectType {
	return m.GetPayload().GetObjectType()
}

func (m *BACnetContextTagObjectIdentifier) GetInstanceNumber() uint32 {
	return m.GetPayload().GetInstanceNumber()
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagObjectIdentifier factory function for BACnetContextTagObjectIdentifier
func NewBACnetContextTagObjectIdentifier(payload *BACnetTagPayloadObjectIdentifier, header *BACnetTagHeader, tagNumberArgument uint8, isNotOpeningOrClosingTag bool) *BACnetContextTag {
	child := &BACnetContextTagObjectIdentifier{
		Payload:          payload,
		BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagObjectIdentifier(structType interface{}) *BACnetContextTagObjectIdentifier {
	if casted, ok := structType.(BACnetContextTagObjectIdentifier); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetContextTagObjectIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(BACnetContextTag); ok {
		return CastBACnetContextTagObjectIdentifier(casted.Child)
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return CastBACnetContextTagObjectIdentifier(casted.Child)
	}
	return nil
}

func (m *BACnetContextTagObjectIdentifier) GetTypeName() string {
	return "BACnetContextTagObjectIdentifier"
}

func (m *BACnetContextTagObjectIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTagObjectIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTagObjectIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagObjectIdentifierParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, isNotOpeningOrClosingTag bool) (*BACnetContextTagObjectIdentifier, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagObjectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(isNotOpeningOrClosingTag) {
		return nil, utils.ParseAssertError{"length 6 and 7 reserved for opening and closing tag"}
	}

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, pullErr
	}
	_payload, _payloadErr := BACnetTagPayloadObjectIdentifierParse(readBuffer)
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadObjectIdentifier(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_objectType := payload.GetObjectType()
	objectType := BACnetObjectType(_objectType)
	_ = objectType

	// Virtual field
	_instanceNumber := payload.GetInstanceNumber()
	instanceNumber := uint32(_instanceNumber)
	_ = instanceNumber

	if closeErr := readBuffer.CloseContext("BACnetContextTagObjectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagObjectIdentifier{
		Payload:          CastBACnetTagPayloadObjectIdentifier(payload),
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child, nil
}

func (m *BACnetContextTagObjectIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagObjectIdentifier"); pushErr != nil {
			return pushErr
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return pushErr
		}
		_payloadErr := m.Payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return popErr
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _objectTypeErr := writeBuffer.WriteVirtual("objectType", m.GetObjectType()); _objectTypeErr != nil {
			return errors.Wrap(_objectTypeErr, "Error serializing 'objectType' field")
		}
		// Virtual field
		if _instanceNumberErr := writeBuffer.WriteVirtual("instanceNumber", m.GetInstanceNumber()); _instanceNumberErr != nil {
			return errors.Wrap(_instanceNumberErr, "Error serializing 'instanceNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagObjectIdentifier"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagObjectIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
