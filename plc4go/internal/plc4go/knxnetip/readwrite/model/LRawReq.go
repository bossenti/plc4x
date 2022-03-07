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
type LRawReq struct {
	*CEMI

	// Arguments.
	Size uint16
}

// The corresponding interface
type ILRawReq interface {
	ICEMI
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
func (m *LRawReq) GetMessageCode() uint8 {
	return 0x10
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *LRawReq) InitializeParent(parent *CEMI) {}

func (m *LRawReq) GetParent() *CEMI {
	return m.CEMI
}

// NewLRawReq factory function for LRawReq
func NewLRawReq(size uint16) *CEMI {
	child := &LRawReq{
		CEMI: NewCEMI(size),
	}
	child.Child = child
	return child.CEMI
}

func CastLRawReq(structType interface{}) *LRawReq {
	if casted, ok := structType.(LRawReq); ok {
		return &casted
	}
	if casted, ok := structType.(*LRawReq); ok {
		return casted
	}
	if casted, ok := structType.(CEMI); ok {
		return CastLRawReq(casted.Child)
	}
	if casted, ok := structType.(*CEMI); ok {
		return CastLRawReq(casted.Child)
	}
	return nil
}

func (m *LRawReq) GetTypeName() string {
	return "LRawReq"
}

func (m *LRawReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *LRawReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *LRawReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LRawReqParse(readBuffer utils.ReadBuffer, size uint16) (*LRawReq, error) {
	if pullErr := readBuffer.PullContext("LRawReq"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LRawReq"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &LRawReq{
		CEMI: &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child, nil
}

func (m *LRawReq) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LRawReq"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("LRawReq"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *LRawReq) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
