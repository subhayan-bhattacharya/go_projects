package handlers

import (
	"event-handler/resources"
	"fmt"
	"reflect"
)

type LoggingHandler[T resources.LoggableAndValidatable] struct {
}

func (l *LoggingHandler[T]) OnAdd(resource T) bool {
	fmt.Printf("ADD: %s\n", resource.Log())
	return true
}

func (l *LoggingHandler[T]) OnDelete(resource T) bool {
	fmt.Printf("DELETE: %s\n", resource.Log())
	return true
}

func (l *LoggingHandler[T]) OnUpdate(old, new T) bool {
	fmt.Printf("UPDATE from : %s to new : %s\n", old.Log(), new.Log())
	return true
}

func (l *LoggingHandler[T]) Name() string {
	var t T
	typeOfT := reflect.TypeOf(t)
	if typeOfT.Kind() == reflect.Ptr {
		return fmt.Sprintf("LoggingHandler<%s>", typeOfT.Elem().Name())
	}
	return fmt.Sprintf("LoggingHandler<%s>", typeOfT.Name())
}
