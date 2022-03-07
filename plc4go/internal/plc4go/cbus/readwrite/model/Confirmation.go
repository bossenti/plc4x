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
type Confirmation struct {
	Alpha *Alpha
	Child IConfirmationChild
}

// The corresponding interface
type IConfirmation interface {
	// GetConfirmationType returns ConfirmationType (discriminator field)
	GetConfirmationType() byte
	// GetAlpha returns Alpha (property field)
	GetAlpha() *Alpha
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IConfirmationParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IConfirmation, serializeChildFunction func() error) error
	GetTypeName() string
}

type IConfirmationChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *Confirmation, alpha *Alpha)
	GetParent() *Confirmation

	GetTypeName() string
	IConfirmation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *Confirmation) GetAlpha() *Alpha {
	return m.Alpha
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConfirmation factory function for Confirmation
func NewConfirmation(alpha *Alpha) *Confirmation {
	return &Confirmation{Alpha: alpha}
}

func CastConfirmation(structType interface{}) *Confirmation {
	if casted, ok := structType.(Confirmation); ok {
		return &casted
	}
	if casted, ok := structType.(*Confirmation); ok {
		return casted
	}
	return nil
}

func (m *Confirmation) GetTypeName() string {
	return "Confirmation"
}

func (m *Confirmation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *Confirmation) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *Confirmation) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (alpha)
	lengthInBits += m.Alpha.GetLengthInBits()
	// Discriminator Field (confirmationType)
	lengthInBits += 8

	return lengthInBits
}

func (m *Confirmation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ConfirmationParse(readBuffer utils.ReadBuffer) (*Confirmation, error) {
	if pullErr := readBuffer.PullContext("Confirmation"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (alpha)
	if pullErr := readBuffer.PullContext("alpha"); pullErr != nil {
		return nil, pullErr
	}
	_alpha, _alphaErr := AlphaParse(readBuffer)
	if _alphaErr != nil {
		return nil, errors.Wrap(_alphaErr, "Error parsing 'alpha' field")
	}
	alpha := CastAlpha(_alpha)
	if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
		return nil, closeErr
	}

	// Discriminator Field (confirmationType) (Used as input to a switch field)
	confirmationType, _confirmationTypeErr := readBuffer.ReadByte("confirmationType")
	if _confirmationTypeErr != nil {
		return nil, errors.Wrap(_confirmationTypeErr, "Error parsing 'confirmationType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ConfirmationChild interface {
		InitializeParent(*Confirmation, *Alpha)
		GetParent() *Confirmation
	}
	var _child ConfirmationChild
	var typeSwitchError error
	switch {
	case confirmationType == 0x2E: // ConfirmationSuccessful
		_child, typeSwitchError = ConfirmationSuccessfulParse(readBuffer)
	case confirmationType == 0x23: // NotTransmittedToManyReTransmissions
		_child, typeSwitchError = NotTransmittedToManyReTransmissionsParse(readBuffer)
	case confirmationType == 0x24: // NotTransmittedCorruption
		_child, typeSwitchError = NotTransmittedCorruptionParse(readBuffer)
	case confirmationType == 0x25: // NotTransmittedSyncLoss
		_child, typeSwitchError = NotTransmittedSyncLossParse(readBuffer)
	case confirmationType == 0x27: // NotTransmittedTooLong
		_child, typeSwitchError = NotTransmittedTooLongParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("Confirmation"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), alpha)
	return _child.GetParent(), nil
}

func (m *Confirmation) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *Confirmation) SerializeParent(writeBuffer utils.WriteBuffer, child IConfirmation, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("Confirmation"); pushErr != nil {
		return pushErr
	}

	// Simple Field (alpha)
	if pushErr := writeBuffer.PushContext("alpha"); pushErr != nil {
		return pushErr
	}
	_alphaErr := m.Alpha.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("alpha"); popErr != nil {
		return popErr
	}
	if _alphaErr != nil {
		return errors.Wrap(_alphaErr, "Error serializing 'alpha' field")
	}

	// Discriminator Field (confirmationType) (Used as input to a switch field)
	confirmationType := byte(child.GetConfirmationType())
	_confirmationTypeErr := writeBuffer.WriteByte("confirmationType", (confirmationType))

	if _confirmationTypeErr != nil {
		return errors.Wrap(_confirmationTypeErr, "Error serializing 'confirmationType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("Confirmation"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *Confirmation) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
