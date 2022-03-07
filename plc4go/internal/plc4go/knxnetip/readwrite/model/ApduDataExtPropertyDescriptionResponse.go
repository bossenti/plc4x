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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ApduDataExtPropertyDescriptionResponse struct {
	*ApduDataExt
	ObjectIndex      uint8
	PropertyId       uint8
	Index            uint8
	WriteEnabled     bool
	PropertyDataType KnxPropertyDataType
	MaxNrOfElements  uint16
	ReadLevel        AccessLevel
	WriteLevel       AccessLevel

	// Arguments.
	Length uint8
}

// The corresponding interface
type IApduDataExtPropertyDescriptionResponse interface {
	IApduDataExt
	// GetObjectIndex returns ObjectIndex (property field)
	GetObjectIndex() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetIndex returns Index (property field)
	GetIndex() uint8
	// GetWriteEnabled returns WriteEnabled (property field)
	GetWriteEnabled() bool
	// GetPropertyDataType returns PropertyDataType (property field)
	GetPropertyDataType() KnxPropertyDataType
	// GetMaxNrOfElements returns MaxNrOfElements (property field)
	GetMaxNrOfElements() uint16
	// GetReadLevel returns ReadLevel (property field)
	GetReadLevel() AccessLevel
	// GetWriteLevel returns WriteLevel (property field)
	GetWriteLevel() AccessLevel
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
func (m *ApduDataExtPropertyDescriptionResponse) GetExtApciType() uint8 {
	return 0x19
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtPropertyDescriptionResponse) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtPropertyDescriptionResponse) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *ApduDataExtPropertyDescriptionResponse) GetObjectIndex() uint8 {
	return m.ObjectIndex
}

func (m *ApduDataExtPropertyDescriptionResponse) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *ApduDataExtPropertyDescriptionResponse) GetIndex() uint8 {
	return m.Index
}

func (m *ApduDataExtPropertyDescriptionResponse) GetWriteEnabled() bool {
	return m.WriteEnabled
}

func (m *ApduDataExtPropertyDescriptionResponse) GetPropertyDataType() KnxPropertyDataType {
	return m.PropertyDataType
}

func (m *ApduDataExtPropertyDescriptionResponse) GetMaxNrOfElements() uint16 {
	return m.MaxNrOfElements
}

func (m *ApduDataExtPropertyDescriptionResponse) GetReadLevel() AccessLevel {
	return m.ReadLevel
}

func (m *ApduDataExtPropertyDescriptionResponse) GetWriteLevel() AccessLevel {
	return m.WriteLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataExtPropertyDescriptionResponse factory function for ApduDataExtPropertyDescriptionResponse
func NewApduDataExtPropertyDescriptionResponse(objectIndex uint8, propertyId uint8, index uint8, writeEnabled bool, propertyDataType KnxPropertyDataType, maxNrOfElements uint16, readLevel AccessLevel, writeLevel AccessLevel, length uint8) *ApduDataExt {
	child := &ApduDataExtPropertyDescriptionResponse{
		ObjectIndex:      objectIndex,
		PropertyId:       propertyId,
		Index:            index,
		WriteEnabled:     writeEnabled,
		PropertyDataType: propertyDataType,
		MaxNrOfElements:  maxNrOfElements,
		ReadLevel:        readLevel,
		WriteLevel:       writeLevel,
		ApduDataExt:      NewApduDataExt(length),
	}
	child.Child = child
	return child.ApduDataExt
}

func CastApduDataExtPropertyDescriptionResponse(structType interface{}) *ApduDataExtPropertyDescriptionResponse {
	if casted, ok := structType.(ApduDataExtPropertyDescriptionResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtPropertyDescriptionResponse); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtPropertyDescriptionResponse(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtPropertyDescriptionResponse(casted.Child)
	}
	return nil
}

func (m *ApduDataExtPropertyDescriptionResponse) GetTypeName() string {
	return "ApduDataExtPropertyDescriptionResponse"
}

func (m *ApduDataExtPropertyDescriptionResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtPropertyDescriptionResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (index)
	lengthInBits += 8

	// Simple field (writeEnabled)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (propertyDataType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (maxNrOfElements)
	lengthInBits += 12

	// Simple field (readLevel)
	lengthInBits += 4

	// Simple field (writeLevel)
	lengthInBits += 4

	return lengthInBits
}

func (m *ApduDataExtPropertyDescriptionResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtPropertyDescriptionResponseParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtPropertyDescriptionResponse, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtPropertyDescriptionResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (objectIndex)
	_objectIndex, _objectIndexErr := readBuffer.ReadUint8("objectIndex", 8)
	if _objectIndexErr != nil {
		return nil, errors.Wrap(_objectIndexErr, "Error parsing 'objectIndex' field")
	}
	objectIndex := _objectIndex

	// Simple Field (propertyId)
	_propertyId, _propertyIdErr := readBuffer.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field")
	}
	propertyId := _propertyId

	// Simple Field (index)
	_index, _indexErr := readBuffer.ReadUint8("index", 8)
	if _indexErr != nil {
		return nil, errors.Wrap(_indexErr, "Error parsing 'index' field")
	}
	index := _index

	// Simple Field (writeEnabled)
	_writeEnabled, _writeEnabledErr := readBuffer.ReadBit("writeEnabled")
	if _writeEnabledErr != nil {
		return nil, errors.Wrap(_writeEnabledErr, "Error parsing 'writeEnabled' field")
	}
	writeEnabled := _writeEnabled

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (propertyDataType)
	if pullErr := readBuffer.PullContext("propertyDataType"); pullErr != nil {
		return nil, pullErr
	}
	_propertyDataType, _propertyDataTypeErr := KnxPropertyDataTypeParse(readBuffer)
	if _propertyDataTypeErr != nil {
		return nil, errors.Wrap(_propertyDataTypeErr, "Error parsing 'propertyDataType' field")
	}
	propertyDataType := _propertyDataType
	if closeErr := readBuffer.CloseContext("propertyDataType"); closeErr != nil {
		return nil, closeErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (maxNrOfElements)
	_maxNrOfElements, _maxNrOfElementsErr := readBuffer.ReadUint16("maxNrOfElements", 12)
	if _maxNrOfElementsErr != nil {
		return nil, errors.Wrap(_maxNrOfElementsErr, "Error parsing 'maxNrOfElements' field")
	}
	maxNrOfElements := _maxNrOfElements

	// Simple Field (readLevel)
	if pullErr := readBuffer.PullContext("readLevel"); pullErr != nil {
		return nil, pullErr
	}
	_readLevel, _readLevelErr := AccessLevelParse(readBuffer)
	if _readLevelErr != nil {
		return nil, errors.Wrap(_readLevelErr, "Error parsing 'readLevel' field")
	}
	readLevel := _readLevel
	if closeErr := readBuffer.CloseContext("readLevel"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (writeLevel)
	if pullErr := readBuffer.PullContext("writeLevel"); pullErr != nil {
		return nil, pullErr
	}
	_writeLevel, _writeLevelErr := AccessLevelParse(readBuffer)
	if _writeLevelErr != nil {
		return nil, errors.Wrap(_writeLevelErr, "Error parsing 'writeLevel' field")
	}
	writeLevel := _writeLevel
	if closeErr := readBuffer.CloseContext("writeLevel"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtPropertyDescriptionResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtPropertyDescriptionResponse{
		ObjectIndex:      objectIndex,
		PropertyId:       propertyId,
		Index:            index,
		WriteEnabled:     writeEnabled,
		PropertyDataType: propertyDataType,
		MaxNrOfElements:  maxNrOfElements,
		ReadLevel:        readLevel,
		WriteLevel:       writeLevel,
		ApduDataExt:      &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtPropertyDescriptionResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtPropertyDescriptionResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (objectIndex)
		objectIndex := uint8(m.ObjectIndex)
		_objectIndexErr := writeBuffer.WriteUint8("objectIndex", 8, (objectIndex))
		if _objectIndexErr != nil {
			return errors.Wrap(_objectIndexErr, "Error serializing 'objectIndex' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.PropertyId)
		_propertyIdErr := writeBuffer.WriteUint8("propertyId", 8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (index)
		index := uint8(m.Index)
		_indexErr := writeBuffer.WriteUint8("index", 8, (index))
		if _indexErr != nil {
			return errors.Wrap(_indexErr, "Error serializing 'index' field")
		}

		// Simple Field (writeEnabled)
		writeEnabled := bool(m.WriteEnabled)
		_writeEnabledErr := writeBuffer.WriteBit("writeEnabled", (writeEnabled))
		if _writeEnabledErr != nil {
			return errors.Wrap(_writeEnabledErr, "Error serializing 'writeEnabled' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 1, uint8(0x0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (propertyDataType)
		if pushErr := writeBuffer.PushContext("propertyDataType"); pushErr != nil {
			return pushErr
		}
		_propertyDataTypeErr := m.PropertyDataType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("propertyDataType"); popErr != nil {
			return popErr
		}
		if _propertyDataTypeErr != nil {
			return errors.Wrap(_propertyDataTypeErr, "Error serializing 'propertyDataType' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 4, uint8(0x0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (maxNrOfElements)
		maxNrOfElements := uint16(m.MaxNrOfElements)
		_maxNrOfElementsErr := writeBuffer.WriteUint16("maxNrOfElements", 12, (maxNrOfElements))
		if _maxNrOfElementsErr != nil {
			return errors.Wrap(_maxNrOfElementsErr, "Error serializing 'maxNrOfElements' field")
		}

		// Simple Field (readLevel)
		if pushErr := writeBuffer.PushContext("readLevel"); pushErr != nil {
			return pushErr
		}
		_readLevelErr := m.ReadLevel.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("readLevel"); popErr != nil {
			return popErr
		}
		if _readLevelErr != nil {
			return errors.Wrap(_readLevelErr, "Error serializing 'readLevel' field")
		}

		// Simple Field (writeLevel)
		if pushErr := writeBuffer.PushContext("writeLevel"); pushErr != nil {
			return pushErr
		}
		_writeLevelErr := m.WriteLevel.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("writeLevel"); popErr != nil {
			return popErr
		}
		if _writeLevelErr != nil {
			return errors.Wrap(_writeLevelErr, "Error serializing 'writeLevel' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtPropertyDescriptionResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtPropertyDescriptionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
