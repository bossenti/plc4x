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

// ApplicationType is an enum
type ApplicationType uint32

type IApplicationType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ApplicationType_applicationTypeServer          ApplicationType = 0
	ApplicationType_applicationTypeClient          ApplicationType = 1
	ApplicationType_applicationTypeClientAndServer ApplicationType = 2
	ApplicationType_applicationTypeDiscoveryServer ApplicationType = 3
)

var ApplicationTypeValues []ApplicationType

func init() {
	_ = errors.New
	ApplicationTypeValues = []ApplicationType{
		ApplicationType_applicationTypeServer,
		ApplicationType_applicationTypeClient,
		ApplicationType_applicationTypeClientAndServer,
		ApplicationType_applicationTypeDiscoveryServer,
	}
}

func ApplicationTypeByValue(value uint32) (enum ApplicationType, ok bool) {
	switch value {
	case 0:
		return ApplicationType_applicationTypeServer, true
	case 1:
		return ApplicationType_applicationTypeClient, true
	case 2:
		return ApplicationType_applicationTypeClientAndServer, true
	case 3:
		return ApplicationType_applicationTypeDiscoveryServer, true
	}
	return 0, false
}

func ApplicationTypeByName(value string) (enum ApplicationType, ok bool) {
	switch value {
	case "applicationTypeServer":
		return ApplicationType_applicationTypeServer, true
	case "applicationTypeClient":
		return ApplicationType_applicationTypeClient, true
	case "applicationTypeClientAndServer":
		return ApplicationType_applicationTypeClientAndServer, true
	case "applicationTypeDiscoveryServer":
		return ApplicationType_applicationTypeDiscoveryServer, true
	}
	return 0, false
}

func ApplicationTypeKnows(value uint32) bool {
	for _, typeValue := range ApplicationTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastApplicationType(structType any) ApplicationType {
	castFunc := func(typ any) ApplicationType {
		if sApplicationType, ok := typ.(ApplicationType); ok {
			return sApplicationType
		}
		return 0
	}
	return castFunc(structType)
}

func (m ApplicationType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m ApplicationType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApplicationTypeParse(ctx context.Context, theBytes []byte) (ApplicationType, error) {
	return ApplicationTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ApplicationTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ApplicationType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("ApplicationType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ApplicationType")
	}
	if enum, ok := ApplicationTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ApplicationType")
		return ApplicationType(val), nil
	} else {
		return enum, nil
	}
}

func (e ApplicationType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ApplicationType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("ApplicationType", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ApplicationType) PLC4XEnumName() string {
	switch e {
	case ApplicationType_applicationTypeServer:
		return "applicationTypeServer"
	case ApplicationType_applicationTypeClient:
		return "applicationTypeClient"
	case ApplicationType_applicationTypeClientAndServer:
		return "applicationTypeClientAndServer"
	case ApplicationType_applicationTypeDiscoveryServer:
		return "applicationTypeDiscoveryServer"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e ApplicationType) String() string {
	return e.PLC4XEnumName()
}
