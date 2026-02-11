package main

import (
	"errors"
	"reflect"
)

type OrderCreated struct {
	Id       string
	Quantity int
	Product  string
}

type UserCreated struct {
	Id    string
	Name  string
	Email string
}

type Subscriber[T any] interface {
	Handle(T) error
}

type Dispatcher struct {
	Notifiers map[reflect.Type]any
}

func NewDispatcher() *Dispatcher {
	notifiers := make(map[reflect.Type]any)
	return &Dispatcher{
		Notifiers: notifiers,
	}
}

func (d *Dispatcher) Register(eventType any, subscriber any) error {
	t := reflect.TypeOf(eventType)
	sval := reflect.ValueOf(subscriber)
	method := sval.MethodByName("Handle")
	if !method.IsValid() {
		return errors.New("The subscriber needs to have a method called Handle")
	}
	d.Notifiers[t] = subscriber
	return nil
}

func main() {

}
