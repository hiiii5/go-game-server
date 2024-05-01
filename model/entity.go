package model

const (
	MAX_MOVE_SPEED = 10.0
)

type Vector struct {
	X, Y, Z float64
}

type Entity struct {
	Position  Vector
	MoveSpeed float64
	Health    float32
	// Do not modify directly
	Id int
}

func (e *Entity) Move(v Vector) {
	e.Position.X += v.X
	e.Position.Y += v.Y
	e.Position.Z += v.Z
}

func (e *Entity) SetPosition(v Vector) {
	e.Position = v
}

func (e *Entity) GetPosition() Vector {
	return e.Position
}

func (e *Entity) SetMoveSpeed(speed float64) {
	e.MoveSpeed = min(speed, MAX_MOVE_SPEED) // clamp
}

func (e *Entity) GetMoveSpeed() float64 {
	return e.MoveSpeed
}
