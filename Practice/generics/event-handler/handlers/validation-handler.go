package handlers

import (
	. "event-handler/resources"
	"fmt"
	"reflect"
)

type ValidationHandler[T LoggableAndValidatable] struct {
}

func (l *ValidationHandler[T]) OnAdd(resource T) bool {
	valid, err := resource.IsValid()
	if !valid || err != nil {
		if err != nil {
			fmt.Printf("Validation failed for resource %v: %v\n", resource.Log(), err)
		} else {
			fmt.Printf("The resource %v is not a valid resource (no specific error provided).\n", resource.Log())
		}
		return false
	}
	return true
}

func (l *ValidationHandler[T]) OnDelete(resource T) bool {
	return true
}

func (l *ValidationHandler[T]) OnUpdate(old, new T) bool {
	valid, err := new.IsValid() // Validate the new resource
	if !valid || err != nil {
		if err != nil {
			fmt.Printf("Validation failed for updated resource %v: %v\n", new.Log(), err)
		} else {
			fmt.Printf("The updated resource %v is not a valid resource (no specific error provided).\n", new.Log())
		}
		return false
	}
	return true
}

func (l *ValidationHandler[T]) Name() string {
	var t T
	typeOfT := reflect.TypeOf(t)
	if typeOfT.Kind() == reflect.Ptr {
		return fmt.Sprintf("ValidationHandler<%s>", typeOfT.Elem().Name())
	}
	return fmt.Sprintf("ValidationHandler<%s>", typeOfT.Name())
}
