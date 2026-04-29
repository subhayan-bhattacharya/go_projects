package commands

import "context"

type Handler[C Command[R], R any] interface {
	Handle(ctx context.Context, cmd C) (R, error)
}
