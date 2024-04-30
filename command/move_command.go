package command

import (
	"errors"
	"fmt"
	"main/model"
	"main/net"
)

const (
	MOVE_COMMAND_BYTE_SIZE = 28
)

type MoveCommand struct {
	Version int
	EntId   int
	Delta   model.Vector
}

func (mc MoveCommand) Parse(data []byte) (Command, error) {
	nd, err := net.TryUnpackData(data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unpacking data: %s", err.Error()))
	}

	if len(nd.Data) != MOVE_COMMAND_BYTE_SIZE {
		return nil, errors.New(fmt.Sprintf("Data was not the correct length: Expected: %d, Actual: %d", MOVE_COMMAND_BYTE_SIZE, len(nd.Data)))
	}

	// Data: idxyz
	entId := net.Int32FromBytes(nd.Data[:4])
	x := net.Float64FromBytes(nd.Data[4:12])
	y := net.Float64FromBytes(nd.Data[12:20])
	z := net.Float64FromBytes(nd.Data[20:28])

	return MoveCommand{Version: int(nd.Version), EntId: int(entId), Delta: model.Vector{X: x, Y: y, Z: z}}, nil
}
