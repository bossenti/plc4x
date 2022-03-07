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
type BACnetApplicationTag struct {
	Header *BACnetTagHeader
	Child  IBACnetApplicationTagChild
}

// The corresponding interface
type IBACnetApplicationTag interface {
	// GetHeader returns Header (property field)
	GetHeader() *BACnetTagHeader
	// GetActualTagNumber returns ActualTagNumber (virtual field)
	GetActualTagNumber() uint8
	// GetActualLength returns ActualLength (virtual field)
	GetActualLength() uint32
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetApplicationTagParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetApplicationTag, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetApplicationTagChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader)
	GetParent() *BACnetApplicationTag

	GetTypeName() string
	IBACnetApplicationTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetApplicationTag) GetHeader() *BACnetTagHeader {
	return m.Header
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////
func (m *BACnetApplicationTag) GetActualTagNumber() uint8 {
	return m.GetHeader().GetActualTagNumber()
}

func (m *BACnetApplicationTag) GetActualLength() uint32 {
	return m.GetHeader().GetActualLength()
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTag factory function for BACnetApplicationTag
func NewBACnetApplicationTag(header *BACnetTagHeader) *BACnetApplicationTag {
	return &BACnetApplicationTag{Header: header}
}

func CastBACnetApplicationTag(structType interface{}) *BACnetApplicationTag {
	if casted, ok := structType.(BACnetApplicationTag); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetApplicationTag); ok {
		return casted
	}
	return nil
}

func (m *BACnetApplicationTag) GetTypeName() string {
	return "BACnetApplicationTag"
}

func (m *BACnetApplicationTag) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetApplicationTag) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetApplicationTag) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetApplicationTag) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetApplicationTagParse(readBuffer utils.ReadBuffer) (*BACnetApplicationTag, error) {
	if pullErr := readBuffer.PullContext("BACnetApplicationTag"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, pullErr
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field")
	}
	header := CastBACnetTagHeader(_header)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, closeErr
	}

	// Validation
	if !(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS))) {
		return nil, utils.ParseAssertError{"should be a application tag"}
	}

	// Virtual field
	_actualTagNumber := header.GetActualTagNumber()
	actualTagNumber := uint8(_actualTagNumber)
	_ = actualTagNumber

	// Virtual field
	_actualLength := header.GetActualLength()
	actualLength := uint32(_actualLength)
	_ = actualLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetApplicationTagChild interface {
		InitializeParent(*BACnetApplicationTag, *BACnetTagHeader)
		GetParent() *BACnetApplicationTag
	}
	var _child BACnetApplicationTagChild
	var typeSwitchError error
	switch {
	case actualTagNumber == 0x0: // BACnetApplicationTagNull
		_child, typeSwitchError = BACnetApplicationTagNullParse(readBuffer)
	case actualTagNumber == 0x1: // BACnetApplicationTagBoolean
		_child, typeSwitchError = BACnetApplicationTagBooleanParse(readBuffer, header)
	case actualTagNumber == 0x2: // BACnetApplicationTagUnsignedInteger
		_child, typeSwitchError = BACnetApplicationTagUnsignedIntegerParse(readBuffer, header)
	case actualTagNumber == 0x3: // BACnetApplicationTagSignedInteger
		_child, typeSwitchError = BACnetApplicationTagSignedIntegerParse(readBuffer, header)
	case actualTagNumber == 0x4: // BACnetApplicationTagReal
		_child, typeSwitchError = BACnetApplicationTagRealParse(readBuffer)
	case actualTagNumber == 0x5: // BACnetApplicationTagDouble
		_child, typeSwitchError = BACnetApplicationTagDoubleParse(readBuffer)
	case actualTagNumber == 0x6: // BACnetApplicationTagOctetString
		_child, typeSwitchError = BACnetApplicationTagOctetStringParse(readBuffer, header)
	case actualTagNumber == 0x7: // BACnetApplicationTagCharacterString
		_child, typeSwitchError = BACnetApplicationTagCharacterStringParse(readBuffer, header)
	case actualTagNumber == 0x8: // BACnetApplicationTagBitString
		_child, typeSwitchError = BACnetApplicationTagBitStringParse(readBuffer, header)
	case actualTagNumber == 0x9: // BACnetApplicationTagEnumerated
		_child, typeSwitchError = BACnetApplicationTagEnumeratedParse(readBuffer, header)
	case actualTagNumber == 0xA: // BACnetApplicationTagDate
		_child, typeSwitchError = BACnetApplicationTagDateParse(readBuffer)
	case actualTagNumber == 0xB: // BACnetApplicationTagTime
		_child, typeSwitchError = BACnetApplicationTagTimeParse(readBuffer)
	case actualTagNumber == 0xC: // BACnetApplicationTagObjectIdentifier
		_child, typeSwitchError = BACnetApplicationTagObjectIdentifierParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetApplicationTag"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), header)
	return _child.GetParent(), nil
}

func (m *BACnetApplicationTag) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetApplicationTag) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetApplicationTag, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetApplicationTag"); pushErr != nil {
		return pushErr
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return pushErr
	}
	_headerErr := m.Header.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return popErr
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}
	// Virtual field
	if _actualTagNumberErr := writeBuffer.WriteVirtual("actualTagNumber", m.GetActualTagNumber()); _actualTagNumberErr != nil {
		return errors.Wrap(_actualTagNumberErr, "Error serializing 'actualTagNumber' field")
	}
	// Virtual field
	if _actualLengthErr := writeBuffer.WriteVirtual("actualLength", m.GetActualLength()); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetApplicationTag"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetApplicationTag) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
