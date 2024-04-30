package model

import (
	"testing"
)

func TestNewEntity(t *testing.T) {
	e := Entity{}
	pos := e.Position
	expectedPos := Vector{X: 0, Y: 0, Z: 0}
	if pos != expectedPos {
		t.Fail()
	}
}

func TestNewEntityWithPosition(t *testing.T) {
	e := Entity{Position: Vector{X: 1, Y: 1, Z: 1}}
	pos := e.Position
	expectedPos := Vector{X: 1, Y: 1, Z: 1}
	if pos != expectedPos {
		t.Fail()
	}
}

func TestEntityMove(t *testing.T) {
	e := Entity{}
	e.Move(Vector{1, 1, 1})
	pos := e.Position
	expectedPos := Vector{X: 1, Y: 1, Z: 1}
	if pos != expectedPos {
		t.Fail()
	}
}

func TestEntitySetPosition(t *testing.T) {
	e := Entity{}
	e.SetPosition(Vector{1, 1, 1})
	pos := e.GetPosition()
	expectedPos := Vector{1, 1, 1}
	if pos != expectedPos {
		t.Fail()
	}
}
