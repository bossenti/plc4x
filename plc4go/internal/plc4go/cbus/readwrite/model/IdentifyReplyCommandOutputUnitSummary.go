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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type IdentifyReplyCommandOutputUnitSummary struct {
	*IdentifyReplyCommand
}

// The corresponding interface
type IIdentifyReplyCommandOutputUnitSummary interface {
	IIdentifyReplyCommand
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
func (m *IdentifyReplyCommandOutputUnitSummary) GetAttribute() Attribute {
	return Attribute_OutputUnitSummary
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *IdentifyReplyCommandOutputUnitSummary) InitializeParent(parent *IdentifyReplyCommand) {}

func (m *IdentifyReplyCommandOutputUnitSummary) GetParent() *IdentifyReplyCommand {
	return m.IdentifyReplyCommand
}

// NewIdentifyReplyCommandOutputUnitSummary factory function for IdentifyReplyCommandOutputUnitSummary
func NewIdentifyReplyCommandOutputUnitSummary() *IdentifyReplyCommand {
	child := &IdentifyReplyCommandOutputUnitSummary{
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	child.Child = child
	return child.IdentifyReplyCommand
}

func CastIdentifyReplyCommandOutputUnitSummary(structType interface{}) *IdentifyReplyCommandOutputUnitSummary {
	if casted, ok := structType.(IdentifyReplyCommandOutputUnitSummary); ok {
		return &casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandOutputUnitSummary); ok {
		return casted
	}
	if casted, ok := structType.(IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandOutputUnitSummary(casted.Child)
	}
	if casted, ok := structType.(*IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandOutputUnitSummary(casted.Child)
	}
	return nil
}

func (m *IdentifyReplyCommandOutputUnitSummary) GetTypeName() string {
	return "IdentifyReplyCommandOutputUnitSummary"
}

func (m *IdentifyReplyCommandOutputUnitSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandOutputUnitSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *IdentifyReplyCommandOutputUnitSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandOutputUnitSummaryParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommandOutputUnitSummary, error) {
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandOutputUnitSummary"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandOutputUnitSummary"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandOutputUnitSummary{
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child, nil
}

func (m *IdentifyReplyCommandOutputUnitSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandOutputUnitSummary"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandOutputUnitSummary"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandOutputUnitSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
