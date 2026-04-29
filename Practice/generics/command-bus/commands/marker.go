package commands

type Command[R any] interface {
	isCommand()
}
type Marker[R any] struct{}

func (m Marker[R]) isCommand() {}
