package handlers

import (
	"event-handler/customerrors"
	"event-handler/resources"
)

type Handler[T resources.LoggableAndValidatable] interface {
	OnAdd(resource T) bool
	OnDelete(resource T) bool
	OnUpdate(old, new T) bool
	Name() string
}

type EventHandlerChain[T resources.LoggableAndValidatable] struct {
	Handlers []Handler[T]
}

func NewEventHandlerChain[T resources.LoggableAndValidatable]() *EventHandlerChain[T] {
	handlers := make([]Handler[T], 0)
	return &EventHandlerChain[T]{
		Handlers: handlers,
	}
}
func (e *EventHandlerChain[T]) Add(handler Handler[T]) *EventHandlerChain[T] {
	e.Handlers = append(e.Handlers, handler)
	return e
}

func (e *EventHandlerChain[T]) OnDelete(resource T) error {
	for _, handler := range e.Handlers {
		if !handler.OnDelete(resource) {
			return customerrors.ResourceDeleteError{
				BaseError: customerrors.BaseError{
					Message: "Could not delete resource",
					Handler: handler.Name(),
				},
			}
		}
	}
	return nil
}

func (e *EventHandlerChain[T]) OnAdd(resource T) error {
	for _, handler := range e.Handlers {
		if !handler.OnAdd(resource) {
			return customerrors.ResourceAddError{
				BaseError: customerrors.BaseError{
					Message: "Could not add resource",
					Handler: handler.Name(),
				},
			}
		}
	}
	return nil
}

func (e *EventHandlerChain[T]) OnUpdate(old, new T) error {
	for _, handler := range e.Handlers {
		if !handler.OnUpdate(old, new) {
			return customerrors.ResourceUpdateError{
				BaseError: customerrors.BaseError{
					Message: "Could not update resource",
					Handler: handler.Name(),
				},
			}
		}
	}
	return nil
}
