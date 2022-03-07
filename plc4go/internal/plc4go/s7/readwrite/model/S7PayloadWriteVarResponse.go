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
type S7PayloadWriteVarResponse struct {
	*S7Payload
	Items []*S7VarPayloadStatusItem

	// Arguments.
	Parameter S7Parameter
}

// The corresponding interface
type IS7PayloadWriteVarResponse interface {
	IS7Payload
	// GetItems returns Items (property field)
	GetItems() []*S7VarPayloadStatusItem
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
func (m *S7PayloadWriteVarResponse) GetParameterParameterType() uint8 {
	return 0x05
}

func (m *S7PayloadWriteVarResponse) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7PayloadWriteVarResponse) InitializeParent(parent *S7Payload) {}

func (m *S7PayloadWriteVarResponse) GetParent() *S7Payload {
	return m.S7Payload
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *S7PayloadWriteVarResponse) GetItems() []*S7VarPayloadStatusItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadWriteVarResponse factory function for S7PayloadWriteVarResponse
func NewS7PayloadWriteVarResponse(items []*S7VarPayloadStatusItem, parameter S7Parameter) *S7Payload {
	child := &S7PayloadWriteVarResponse{
		Items:     items,
		S7Payload: NewS7Payload(parameter),
	}
	child.Child = child
	return child.S7Payload
}

func CastS7PayloadWriteVarResponse(structType interface{}) *S7PayloadWriteVarResponse {
	if casted, ok := structType.(S7PayloadWriteVarResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadWriteVarResponse); ok {
		return casted
	}
	if casted, ok := structType.(S7Payload); ok {
		return CastS7PayloadWriteVarResponse(casted.Child)
	}
	if casted, ok := structType.(*S7Payload); ok {
		return CastS7PayloadWriteVarResponse(casted.Child)
	}
	return nil
}

func (m *S7PayloadWriteVarResponse) GetTypeName() string {
	return "S7PayloadWriteVarResponse"
}

func (m *S7PayloadWriteVarResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadWriteVarResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7PayloadWriteVarResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadWriteVarResponseParse(readBuffer utils.ReadBuffer, messageType uint8, parameter *S7Parameter) (*S7PayloadWriteVarResponse, error) {
	if pullErr := readBuffer.PullContext("S7PayloadWriteVarResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	items := make([]*S7VarPayloadStatusItem, CastS7ParameterWriteVarResponse(parameter).GetNumItems())
	{
		for curItem := uint16(0); curItem < uint16(CastS7ParameterWriteVarResponse(parameter).GetNumItems()); curItem++ {
			_item, _err := S7VarPayloadStatusItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items[curItem] = CastS7VarPayloadStatusItem(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadWriteVarResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadWriteVarResponse{
		Items:     items,
		S7Payload: &S7Payload{},
	}
	_child.S7Payload.Child = _child
	return _child, nil
}

func (m *S7PayloadWriteVarResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadWriteVarResponse"); pushErr != nil {
			return pushErr
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadWriteVarResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadWriteVarResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
