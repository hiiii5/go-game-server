package command

import (
	"errors"
	"fmt"
	"log"
	"main/model"
	"main/net"
	//"strconv"
	//"strings"
)

const (
	CREATE_ENTITY_COMMAND = iota
	MOVE_COMMAND
)

type Commander struct {
	entities []model.Entity
}

func TryParse(data []byte) (Command, error) {
	if len(data) < 1 {
		return nil, errors.New("Data is empty")
	}

	nd, err := net.TryUnpackData(data)
	if err != nil {
		return nil, errors.New("Error unpacking data: " + err.Error())
	}

	if nd.Id != CREATE_ENTITY_COMMAND && nd.Id != MOVE_COMMAND {
		return nil, errors.New(fmt.Sprintf("Command not recognized: %d", nd.Id))
	}

	var c Command
	switch nd.Id {
	case CREATE_ENTITY_COMMAND:
		cec := CreateEntityCommand{}
		c, err = cec.Parse(data)
	case MOVE_COMMAND:
		mc := MoveCommand{}
		c, err = mc.Parse(data)
	default:
		return nil, errors.New(fmt.Sprintf("Command not recognized: %d", nd.Id))
	}

	return c, err
}

func (c *Commander) AddEntity(e model.Entity) {
	id := len(c.entities)
	e.Id = id
	c.entities = append(c.entities, e)
}

func (c *Commander) GetEntity(id int) model.Entity {
	for _, e := range c.entities {
		if e.Id == id {
			return e
		}
	}

	return model.Entity{}
}

func (c *Commander) ExecuteCommand(cmd Command) error {
	switch cmd.(type) {
	case CreateEntityCommand:
		log.Printf("Executing create entity command with data: %v", cmd)
		cec := cmd.(CreateEntityCommand)
		c.AddEntity(cec.Entity)
		return nil
	case MoveCommand:
		log.Printf("Executing move command with data: %v", cmd)
		mc := cmd.(MoveCommand)
		e := c.GetEntity(mc.EntId)
		e.Move(mc.Delta)
		c.entities[mc.EntId] = e
		return nil
	default:
		return errors.New(fmt.Sprintf("Command not recognized: type: %T, value: %v", cmd, cmd))
	}
}
