package command

type Command interface {
	Parse(data []byte) (Command, error)
}
