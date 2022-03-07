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
type COTPPacketDisconnectResponse struct {
	*COTPPacket
	DestinationReference uint16
	SourceReference      uint16

	// Arguments.
	CotpLen uint16
}

// The corresponding interface
type ICOTPPacketDisconnectResponse interface {
	ICOTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetSourceReference returns SourceReference (property field)
	GetSourceReference() uint16
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
func (m *COTPPacketDisconnectResponse) GetTpduCode() uint8 {
	return 0xC0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *COTPPacketDisconnectResponse) InitializeParent(parent *COTPPacket, parameters []*COTPParameter, payload *S7Message) {
	m.COTPPacket.Parameters = parameters
	m.COTPPacket.Payload = payload
}

func (m *COTPPacketDisconnectResponse) GetParent() *COTPPacket {
	return m.COTPPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *COTPPacketDisconnectResponse) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *COTPPacketDisconnectResponse) GetSourceReference() uint16 {
	return m.SourceReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPPacketDisconnectResponse factory function for COTPPacketDisconnectResponse
func NewCOTPPacketDisconnectResponse(destinationReference uint16, sourceReference uint16, parameters []*COTPParameter, payload *S7Message, cotpLen uint16) *COTPPacket {
	child := &COTPPacketDisconnectResponse{
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		COTPPacket:           NewCOTPPacket(parameters, payload, cotpLen),
	}
	child.Child = child
	return child.COTPPacket
}

func CastCOTPPacketDisconnectResponse(structType interface{}) *COTPPacketDisconnectResponse {
	if casted, ok := structType.(COTPPacketDisconnectResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*COTPPacketDisconnectResponse); ok {
		return casted
	}
	if casted, ok := structType.(COTPPacket); ok {
		return CastCOTPPacketDisconnectResponse(casted.Child)
	}
	if casted, ok := structType.(*COTPPacket); ok {
		return CastCOTPPacketDisconnectResponse(casted.Child)
	}
	return nil
}

func (m *COTPPacketDisconnectResponse) GetTypeName() string {
	return "COTPPacketDisconnectResponse"
}

func (m *COTPPacketDisconnectResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *COTPPacketDisconnectResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (sourceReference)
	lengthInBits += 16

	return lengthInBits
}

func (m *COTPPacketDisconnectResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPPacketDisconnectResponseParse(readBuffer utils.ReadBuffer, cotpLen uint16) (*COTPPacketDisconnectResponse, error) {
	if pullErr := readBuffer.PullContext("COTPPacketDisconnectResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (destinationReference)
	_destinationReference, _destinationReferenceErr := readBuffer.ReadUint16("destinationReference", 16)
	if _destinationReferenceErr != nil {
		return nil, errors.Wrap(_destinationReferenceErr, "Error parsing 'destinationReference' field")
	}
	destinationReference := _destinationReference

	// Simple Field (sourceReference)
	_sourceReference, _sourceReferenceErr := readBuffer.ReadUint16("sourceReference", 16)
	if _sourceReferenceErr != nil {
		return nil, errors.Wrap(_sourceReferenceErr, "Error parsing 'sourceReference' field")
	}
	sourceReference := _sourceReference

	if closeErr := readBuffer.CloseContext("COTPPacketDisconnectResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &COTPPacketDisconnectResponse{
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		COTPPacket:           &COTPPacket{},
	}
	_child.COTPPacket.Child = _child
	return _child, nil
}

func (m *COTPPacketDisconnectResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketDisconnectResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (destinationReference)
		destinationReference := uint16(m.DestinationReference)
		_destinationReferenceErr := writeBuffer.WriteUint16("destinationReference", 16, (destinationReference))
		if _destinationReferenceErr != nil {
			return errors.Wrap(_destinationReferenceErr, "Error serializing 'destinationReference' field")
		}

		// Simple Field (sourceReference)
		sourceReference := uint16(m.SourceReference)
		_sourceReferenceErr := writeBuffer.WriteUint16("sourceReference", 16, (sourceReference))
		if _sourceReferenceErr != nil {
			return errors.Wrap(_sourceReferenceErr, "Error serializing 'sourceReference' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketDisconnectResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *COTPPacketDisconnectResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
