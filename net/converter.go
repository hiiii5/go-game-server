package net

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"slices"
)

const (
	// 2 bytes for the version and id, and 512 bytes for the data
	MAX_NET_DATA_LENGTH = 514
)

type NetData struct {
	Version byte
	Id      byte
	Data    []byte
}

func TryUnpackData(data []byte) (*NetData, error) {
	// We recieve in big endian
	slices.Reverse(data)

	if len(data) > MAX_NET_DATA_LENGTH {
		return nil, errors.New(fmt.Sprintf("Data was longer than the max length:\nExpected: %d, Actual: %d", MAX_NET_DATA_LENGTH, len(data)))
	}

	version := data[0]
	id := data[1]
	data = data[2:]

	return &NetData{Version: version, Id: id, Data: data}, nil
}

func (nd NetData) Pack() []byte {
	data := make([]byte, 0)
	data = append(data, nd.Version, nd.Id)
	data = append(data, nd.Data...)
	// We send in big endian
	slices.Reverse(data)
	return data
}

func (nd NetData) String() string {
	return fmt.Sprintf("Version: %d, Id: %d, Data: %s", nd.Version, nd.Id, string(nd.Data))
}

func Int32FromBytes(bytes []byte) int32 {
	return int32(binary.LittleEndian.Uint32(bytes))
}

func Int32ToBytes(i int32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(i))
	return bytes
}

func Float32FromBytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func Float32ToBytes(f float32) []byte {
	bits := math.Float32bits(f)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func Float64FromBytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

func Float64ToBytes(f float64) []byte {
	bits := math.Float64bits(f)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}
