/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"fmt"

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// MonitoringMode is an enum
type MonitoringMode uint32

type IMonitoringMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	MonitoringMode_monitoringModeDisabled  MonitoringMode = 0
	MonitoringMode_monitoringModeSampling  MonitoringMode = 1
	MonitoringMode_monitoringModeReporting MonitoringMode = 2
)

var MonitoringModeValues []MonitoringMode

func init() {
	_ = errors.New
	MonitoringModeValues = []MonitoringMode{
		MonitoringMode_monitoringModeDisabled,
		MonitoringMode_monitoringModeSampling,
		MonitoringMode_monitoringModeReporting,
	}
}

func MonitoringModeByValue(value uint32) (enum MonitoringMode, ok bool) {
	switch value {
	case 0:
		return MonitoringMode_monitoringModeDisabled, true
	case 1:
		return MonitoringMode_monitoringModeSampling, true
	case 2:
		return MonitoringMode_monitoringModeReporting, true
	}
	return 0, false
}

func MonitoringModeByName(value string) (enum MonitoringMode, ok bool) {
	switch value {
	case "monitoringModeDisabled":
		return MonitoringMode_monitoringModeDisabled, true
	case "monitoringModeSampling":
		return MonitoringMode_monitoringModeSampling, true
	case "monitoringModeReporting":
		return MonitoringMode_monitoringModeReporting, true
	}
	return 0, false
}

func MonitoringModeKnows(value uint32) bool {
	for _, typeValue := range MonitoringModeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastMonitoringMode(structType any) MonitoringMode {
	castFunc := func(typ any) MonitoringMode {
		if sMonitoringMode, ok := typ.(MonitoringMode); ok {
			return sMonitoringMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m MonitoringMode) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m MonitoringMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoringModeParse(ctx context.Context, theBytes []byte) (MonitoringMode, error) {
	return MonitoringModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MonitoringModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MonitoringMode, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("MonitoringMode", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading MonitoringMode")
	}
	if enum, ok := MonitoringModeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for MonitoringMode")
		return MonitoringMode(val), nil
	} else {
		return enum, nil
	}
}

func (e MonitoringMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e MonitoringMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("MonitoringMode", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e MonitoringMode) PLC4XEnumName() string {
	switch e {
	case MonitoringMode_monitoringModeDisabled:
		return "monitoringModeDisabled"
	case MonitoringMode_monitoringModeSampling:
		return "monitoringModeSampling"
	case MonitoringMode_monitoringModeReporting:
		return "monitoringModeReporting"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e MonitoringMode) String() string {
	return e.PLC4XEnumName()
}
