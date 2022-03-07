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
type S7PayloadAlarmSQ struct {
	*S7PayloadUserDataItem
	AlarmMessage *AlarmMessagePushType
}

// The corresponding interface
type IS7PayloadAlarmSQ interface {
	IS7PayloadUserDataItem
	// GetAlarmMessage returns AlarmMessage (property field)
	GetAlarmMessage() *AlarmMessagePushType
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
func (m *S7PayloadAlarmSQ) GetCpuFunctionType() uint8 {
	return 0x00
}

func (m *S7PayloadAlarmSQ) GetCpuSubfunction() uint8 {
	return 0x11
}

func (m *S7PayloadAlarmSQ) GetDataLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7PayloadAlarmSQ) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.S7PayloadUserDataItem.ReturnCode = returnCode
	m.S7PayloadUserDataItem.TransportSize = transportSize
}

func (m *S7PayloadAlarmSQ) GetParent() *S7PayloadUserDataItem {
	return m.S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *S7PayloadAlarmSQ) GetAlarmMessage() *AlarmMessagePushType {
	return m.AlarmMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadAlarmSQ factory function for S7PayloadAlarmSQ
func NewS7PayloadAlarmSQ(alarmMessage *AlarmMessagePushType, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItem {
	child := &S7PayloadAlarmSQ{
		AlarmMessage:          alarmMessage,
		S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	child.Child = child
	return child.S7PayloadUserDataItem
}

func CastS7PayloadAlarmSQ(structType interface{}) *S7PayloadAlarmSQ {
	if casted, ok := structType.(S7PayloadAlarmSQ); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadAlarmSQ); ok {
		return casted
	}
	if casted, ok := structType.(S7PayloadUserDataItem); ok {
		return CastS7PayloadAlarmSQ(casted.Child)
	}
	if casted, ok := structType.(*S7PayloadUserDataItem); ok {
		return CastS7PayloadAlarmSQ(casted.Child)
	}
	return nil
}

func (m *S7PayloadAlarmSQ) GetTypeName() string {
	return "S7PayloadAlarmSQ"
}

func (m *S7PayloadAlarmSQ) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadAlarmSQ) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (alarmMessage)
	lengthInBits += m.AlarmMessage.GetLengthInBits()

	return lengthInBits
}

func (m *S7PayloadAlarmSQ) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadAlarmSQParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadAlarmSQ, error) {
	if pullErr := readBuffer.PullContext("S7PayloadAlarmSQ"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (alarmMessage)
	if pullErr := readBuffer.PullContext("alarmMessage"); pullErr != nil {
		return nil, pullErr
	}
	_alarmMessage, _alarmMessageErr := AlarmMessagePushTypeParse(readBuffer)
	if _alarmMessageErr != nil {
		return nil, errors.Wrap(_alarmMessageErr, "Error parsing 'alarmMessage' field")
	}
	alarmMessage := CastAlarmMessagePushType(_alarmMessage)
	if closeErr := readBuffer.CloseContext("alarmMessage"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadAlarmSQ"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadAlarmSQ{
		AlarmMessage:          CastAlarmMessagePushType(alarmMessage),
		S7PayloadUserDataItem: &S7PayloadUserDataItem{},
	}
	_child.S7PayloadUserDataItem.Child = _child
	return _child, nil
}

func (m *S7PayloadAlarmSQ) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadAlarmSQ"); pushErr != nil {
			return pushErr
		}

		// Simple Field (alarmMessage)
		if pushErr := writeBuffer.PushContext("alarmMessage"); pushErr != nil {
			return pushErr
		}
		_alarmMessageErr := m.AlarmMessage.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("alarmMessage"); popErr != nil {
			return popErr
		}
		if _alarmMessageErr != nil {
			return errors.Wrap(_alarmMessageErr, "Error serializing 'alarmMessage' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadAlarmSQ"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadAlarmSQ) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
