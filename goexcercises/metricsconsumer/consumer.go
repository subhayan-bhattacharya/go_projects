package metricsconsumer

import (
	"fmt"
	"reflect"
	"time"
)

// NamespaceSnapshot is a single point-in-time reading for one namespace.
// This is YOUR type — you own it.
type NamespaceSnapshot struct {
	Namespace string
	CPUCores  float64 // e.g. 1.75 cores in use
	MemoryMB  float64 // e.g. 512.0 MB in use
	PodCount  int
	Timestamp time.Time
}

func extractStructFields(snapshopt NamespaceSnapshot) []string {
	t := reflect.TypeOf(snapshopt)
	numFields := t.NumField()
	fieldNames := make([]string, numFields)
	for i := range numFields {
		fieldNames[i] = t.Field(i).Name
	}
	return fieldNames
}

type RenderBackend[T any] interface {
	Render(snapshot T) error
}

type TerminalWriterRenderer[T any] struct {
	writer TerminalWriter
}

func NewTerminalWriterRenderer[T any](writer TerminalWriter) TerminalWriterRenderer[T] {
	return TerminalWriterRenderer[T]{
		writer: writer,
	}
}

func (r TerminalWriterRenderer[T]) Render(snapshot T) error {
	nsSnapshot, ok := any(snapshot).(NamespaceSnapshot)
	if ok {
		_ := extractStructFields(nsSnapshot)
	} else {
		return fmt.Errorf("could not send data as data sent is not the correct type.")
	}
	return nil
}
