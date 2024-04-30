package command

import (
	"main/net"
	"testing"
)

func Test_Command_ParseMoveEntityCommand(t *testing.T) {
	version := int32(1)
	cmdId := MOVE_COMMAND
	entityId := int32(0)
	deltaX := float64(5.0)
	deltaY := float64(6.0)
	deltaZ := float64(7.0)

	// Set data in network byte order
	var data []byte
	data = append(data, net.Int32ToBytes(int32(entityId))...)
	data = append(data, net.Float64ToBytes(deltaX)...)
	data = append(data, net.Float64ToBytes(deltaY)...)
	data = append(data, net.Float64ToBytes(deltaZ)...)

	nd := net.NetData{Version: byte(version), Id: byte(cmdId), Data: data}
	bytes := nd.Pack()

	mc := MoveCommand{}
	c, err := mc.Parse(bytes)
	if err != nil {
		t.Fatalf("Error parsing move command:\n%s", err.Error())
	}

	mc, ok := c.(MoveCommand)
	if !ok {
		t.Fatalf("Could not cast command to move command")
	}

	if mc.Version != int(version) {
		t.Errorf("Expected version to be %d, got %d", version, mc.Version)
	}
	if mc.EntId != int(entityId) {
		t.Errorf("Expected id to be %d, got %d", entityId, mc.EntId)
	}
	if mc.Delta.X != deltaX {
		t.Errorf("Expected X to be %f, got %f", deltaX, mc.Delta.X)
	}
	if mc.Delta.Y != deltaY {
		t.Errorf("Expected Y to be %f, got %f", deltaY, mc.Delta.Y)
	}
	if mc.Delta.Z != deltaZ {
		t.Errorf("Expected Z to be %f, got %f", deltaZ, mc.Delta.Z)
	}
}
