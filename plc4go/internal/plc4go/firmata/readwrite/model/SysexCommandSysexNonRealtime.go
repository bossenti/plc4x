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
type SysexCommandSysexNonRealtime struct {
	*SysexCommand
}

// The corresponding interface
type ISysexCommandSysexNonRealtime interface {
	ISysexCommand
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
func (m *SysexCommandSysexNonRealtime) GetCommandType() uint8 {
	return 0x7E
}

func (m *SysexCommandSysexNonRealtime) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *SysexCommandSysexNonRealtime) InitializeParent(parent *SysexCommand) {}

func (m *SysexCommandSysexNonRealtime) GetParent() *SysexCommand {
	return m.SysexCommand
}

// NewSysexCommandSysexNonRealtime factory function for SysexCommandSysexNonRealtime
func NewSysexCommandSysexNonRealtime() *SysexCommand {
	child := &SysexCommandSysexNonRealtime{
		SysexCommand: NewSysexCommand(),
	}
	child.Child = child
	return child.SysexCommand
}

func CastSysexCommandSysexNonRealtime(structType interface{}) *SysexCommandSysexNonRealtime {
	if casted, ok := structType.(SysexCommandSysexNonRealtime); ok {
		return &casted
	}
	if casted, ok := structType.(*SysexCommandSysexNonRealtime); ok {
		return casted
	}
	if casted, ok := structType.(SysexCommand); ok {
		return CastSysexCommandSysexNonRealtime(casted.Child)
	}
	if casted, ok := structType.(*SysexCommand); ok {
		return CastSysexCommandSysexNonRealtime(casted.Child)
	}
	return nil
}

func (m *SysexCommandSysexNonRealtime) GetTypeName() string {
	return "SysexCommandSysexNonRealtime"
}

func (m *SysexCommandSysexNonRealtime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SysexCommandSysexNonRealtime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *SysexCommandSysexNonRealtime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandSysexNonRealtimeParse(readBuffer utils.ReadBuffer, response bool) (*SysexCommandSysexNonRealtime, error) {
	if pullErr := readBuffer.PullContext("SysexCommandSysexNonRealtime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandSysexNonRealtime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandSysexNonRealtime{
		SysexCommand: &SysexCommand{},
	}
	_child.SysexCommand.Child = _child
	return _child, nil
}

func (m *SysexCommandSysexNonRealtime) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSysexNonRealtime"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("SysexCommandSysexNonRealtime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandSysexNonRealtime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
