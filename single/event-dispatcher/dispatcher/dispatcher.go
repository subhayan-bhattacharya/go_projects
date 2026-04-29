package dispatcher

import (
	"event-dispatcher/handlers"
	"fmt"
	"sync"
)

type Dispatcher[T any] struct {
	handlers []handlers.Handler[T]
	mu       sync.Mutex
}

func NewDispatcher[T any]() *Dispatcher[T] {
	return &Dispatcher[T]{
		handlers: make([]handlers.Handler[T], 0),
	}
}

func (d *Dispatcher[T]) Subscribe(handler handlers.Handler[T]) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers = append(d.handlers, handler)
	return nil
}

func (d *Dispatcher[T]) Publish(event T) error {
	var wg sync.WaitGroup
	d.mu.Lock()
	copiedHandlers := make([]handlers.Handler[T], len(d.handlers))
	_ = copy(copiedHandlers, d.handlers)
	d.mu.Unlock()
	wg.Add(len(copiedHandlers))
	for _, handler := range copiedHandlers {
		go func() {
			defer wg.Done()
			defer func() {
				if v := recover(); v != nil {
					fmt.Printf("panic was caused by handler function for input %v\n", v)
				}
			}()
			handler(event)
		}()
	}
	wg.Wait()
	return nil
}
