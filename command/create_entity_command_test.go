package command

import (
	"main/net"
	"testing"
)

func Test_Command_ParseCreateEntity(t *testing.T) {
	version := int32(1)
	id := CREATE_ENTITY_COMMAND
	posX := float64(5.0)
	posY := float64(6.0)
	posZ := float64(7.0)
	moveSpeed := float64(4.5)

	// Set data in network byte order
	var data []byte
	data = append(data, net.Float64ToBytes(posX)...)
	data = append(data, net.Float64ToBytes(posY)...)
	data = append(data, net.Float64ToBytes(posZ)...)
	data = append(data, net.Float64ToBytes(moveSpeed)...)

	nd := net.NetData{Version: byte(version), Id: byte(id), Data: data}
	bytes := nd.Pack()
	t.Logf("\nData: %v, length: %d", bytes, len(bytes))

	cec := CreateEntityCommand{}
	c, err := cec.Parse(bytes)
	if err != nil {
		t.Fatalf("Error parsing create entity command:\n%s", err.Error())
	}

	cec, ok := c.(CreateEntityCommand)
	if !ok {
		t.Fatalf("Could not cast command to create entity command")
	}

	if cec.Version != int(version) {
		t.Errorf("Expected version to be %d, got %d", version, cec.Version)
	}
	if cec.Entity.Position.X != posX {
		t.Errorf("Expected X to be %f, got %f", posX, cec.Entity.Position.X)
	}
	if cec.Entity.Position.Y != posY {
		t.Errorf("Expected Y to be %f, got %f", posY, cec.Entity.Position.Y)
	}
	if cec.Entity.Position.Z != posZ {
		t.Errorf("Expected Z to be %f, got %f", posZ, cec.Entity.Position.Z)
	}
	if cec.Entity.MoveSpeed != moveSpeed {
		t.Errorf("Expected move speed to be %f, got %f", moveSpeed, cec.Entity.MoveSpeed)
	}
}
