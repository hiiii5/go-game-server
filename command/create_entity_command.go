package command

import (
	"errors"
	"fmt"
	"main/model"
	"main/net"
)

const (
	CREATE_ENTITY_BYTE_SIZE = 32
)

type CreateEntityCommand struct {
	Version int
	Entity  model.Entity
}

func (cec CreateEntityCommand) Parse(data []byte) (Command, error) {
	nd, err := net.TryUnpackData(data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unpacking data: %s", err.Error()))
	}

	if len(nd.Data) != CREATE_ENTITY_BYTE_SIZE {
		return nil, errors.New(fmt.Sprintf("Data was not the correct length:\nExpected: %d, Actual: %d\ndata: %v", CREATE_ENTITY_BYTE_SIZE, len(nd.Data), nd.Data))
	}

	// Data: xyzspeed
	x := net.Float64FromBytes(nd.Data[:8])
	y := net.Float64FromBytes(nd.Data[8:16])
	z := net.Float64FromBytes(nd.Data[16:24])
	speed := net.Float64FromBytes(nd.Data[24:32])

	return CreateEntityCommand{Version: int(nd.Version), Entity: model.Entity{Position: model.Vector{X: x, Y: y, Z: z}, MoveSpeed: speed}}, nil
}
