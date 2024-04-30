package command

import (
	"main/net"
	"testing"
)

func Test_Commander_ExecuteCommand(t *testing.T) {
	commander := Commander{}

	t.Run("TestParseCreateEntityCommand", func(t *testing.T) {
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

		// This has to execute to test the move command
		err = commander.ExecuteCommand(c)
		if err != nil {
			t.Fatalf("Error executing create entity command:\n%s", err.Error())
		}
	})

	t.Run("TestParseMoveCommand", func(t *testing.T) {
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

		err = commander.ExecuteCommand(c)
		if err != nil {
			t.Fatalf("Error executing move command:\n%s", err.Error())
		}
	})
}
