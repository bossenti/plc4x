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
type AdsMultiRequestItemRead struct {
	*AdsMultiRequestItem
	ItemIndexGroup  uint32
	ItemIndexOffset uint32
	ItemReadLength  uint32
}

// The corresponding interface
type IAdsMultiRequestItemRead interface {
	IAdsMultiRequestItem
	// GetItemIndexGroup returns ItemIndexGroup (property field)
	GetItemIndexGroup() uint32
	// GetItemIndexOffset returns ItemIndexOffset (property field)
	GetItemIndexOffset() uint32
	// GetItemReadLength returns ItemReadLength (property field)
	GetItemReadLength() uint32
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
func (m *AdsMultiRequestItemRead) GetIndexGroup() uint32 {
	return uint32(61568)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *AdsMultiRequestItemRead) InitializeParent(parent *AdsMultiRequestItem) {}

func (m *AdsMultiRequestItemRead) GetParent() *AdsMultiRequestItem {
	return m.AdsMultiRequestItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *AdsMultiRequestItemRead) GetItemIndexGroup() uint32 {
	return m.ItemIndexGroup
}

func (m *AdsMultiRequestItemRead) GetItemIndexOffset() uint32 {
	return m.ItemIndexOffset
}

func (m *AdsMultiRequestItemRead) GetItemReadLength() uint32 {
	return m.ItemReadLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsMultiRequestItemRead factory function for AdsMultiRequestItemRead
func NewAdsMultiRequestItemRead(itemIndexGroup uint32, itemIndexOffset uint32, itemReadLength uint32) *AdsMultiRequestItem {
	child := &AdsMultiRequestItemRead{
		ItemIndexGroup:      itemIndexGroup,
		ItemIndexOffset:     itemIndexOffset,
		ItemReadLength:      itemReadLength,
		AdsMultiRequestItem: NewAdsMultiRequestItem(),
	}
	child.Child = child
	return child.AdsMultiRequestItem
}

func CastAdsMultiRequestItemRead(structType interface{}) *AdsMultiRequestItemRead {
	if casted, ok := structType.(AdsMultiRequestItemRead); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsMultiRequestItemRead); ok {
		return casted
	}
	if casted, ok := structType.(AdsMultiRequestItem); ok {
		return CastAdsMultiRequestItemRead(casted.Child)
	}
	if casted, ok := structType.(*AdsMultiRequestItem); ok {
		return CastAdsMultiRequestItemRead(casted.Child)
	}
	return nil
}

func (m *AdsMultiRequestItemRead) GetTypeName() string {
	return "AdsMultiRequestItemRead"
}

func (m *AdsMultiRequestItemRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsMultiRequestItemRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (itemIndexGroup)
	lengthInBits += 32

	// Simple field (itemIndexOffset)
	lengthInBits += 32

	// Simple field (itemReadLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsMultiRequestItemRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsMultiRequestItemReadParse(readBuffer utils.ReadBuffer, indexGroup uint32) (*AdsMultiRequestItemRead, error) {
	if pullErr := readBuffer.PullContext("AdsMultiRequestItemRead"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (itemIndexGroup)
	_itemIndexGroup, _itemIndexGroupErr := readBuffer.ReadUint32("itemIndexGroup", 32)
	if _itemIndexGroupErr != nil {
		return nil, errors.Wrap(_itemIndexGroupErr, "Error parsing 'itemIndexGroup' field")
	}
	itemIndexGroup := _itemIndexGroup

	// Simple Field (itemIndexOffset)
	_itemIndexOffset, _itemIndexOffsetErr := readBuffer.ReadUint32("itemIndexOffset", 32)
	if _itemIndexOffsetErr != nil {
		return nil, errors.Wrap(_itemIndexOffsetErr, "Error parsing 'itemIndexOffset' field")
	}
	itemIndexOffset := _itemIndexOffset

	// Simple Field (itemReadLength)
	_itemReadLength, _itemReadLengthErr := readBuffer.ReadUint32("itemReadLength", 32)
	if _itemReadLengthErr != nil {
		return nil, errors.Wrap(_itemReadLengthErr, "Error parsing 'itemReadLength' field")
	}
	itemReadLength := _itemReadLength

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItemRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsMultiRequestItemRead{
		ItemIndexGroup:      itemIndexGroup,
		ItemIndexOffset:     itemIndexOffset,
		ItemReadLength:      itemReadLength,
		AdsMultiRequestItem: &AdsMultiRequestItem{},
	}
	_child.AdsMultiRequestItem.Child = _child
	return _child, nil
}

func (m *AdsMultiRequestItemRead) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsMultiRequestItemRead"); pushErr != nil {
			return pushErr
		}

		// Simple Field (itemIndexGroup)
		itemIndexGroup := uint32(m.ItemIndexGroup)
		_itemIndexGroupErr := writeBuffer.WriteUint32("itemIndexGroup", 32, (itemIndexGroup))
		if _itemIndexGroupErr != nil {
			return errors.Wrap(_itemIndexGroupErr, "Error serializing 'itemIndexGroup' field")
		}

		// Simple Field (itemIndexOffset)
		itemIndexOffset := uint32(m.ItemIndexOffset)
		_itemIndexOffsetErr := writeBuffer.WriteUint32("itemIndexOffset", 32, (itemIndexOffset))
		if _itemIndexOffsetErr != nil {
			return errors.Wrap(_itemIndexOffsetErr, "Error serializing 'itemIndexOffset' field")
		}

		// Simple Field (itemReadLength)
		itemReadLength := uint32(m.ItemReadLength)
		_itemReadLengthErr := writeBuffer.WriteUint32("itemReadLength", 32, (itemReadLength))
		if _itemReadLengthErr != nil {
			return errors.Wrap(_itemReadLengthErr, "Error serializing 'itemReadLength' field")
		}

		if popErr := writeBuffer.PopContext("AdsMultiRequestItemRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsMultiRequestItemRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
