package metricsconsumer

import (
	"strconv"
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

func NamespaceSnapshotConverter(snapshot NamespaceSnapshot) ([]string, []string) {
	headers := []string{"Namespace", "CPUCores", "MemoryMB", "PodCount", "Timestamp"}
	values := []string{
		snapshot.Namespace,
		strconv.FormatFloat(snapshot.CPUCores, 'f', -1, 64),
		strconv.FormatFloat(snapshot.MemoryMB, 'f', -1, 64),
		strconv.Itoa(snapshot.PodCount),
		snapshot.Timestamp.Format(time.RFC3339),
	}
	return headers, values
}

type RenderBackend[T any] interface {
	RenderAll(snapshpts []T) error
}

type converterFunc[T any] func(T) (headers []string, values []string)

type TerminalWriterRenderer[T any] struct {
	writer    TerminalWriter
	converter converterFunc[T]
}

func NewTerminalWriterRenderer[T any](writer TerminalWriter, converter converterFunc[T]) TerminalWriterRenderer[T] {
	return TerminalWriterRenderer[T]{
		writer:    writer,
		converter: converter,
	}
}

func (r TerminalWriterRenderer[T]) RenderAll(snapshots []T) error {
	var headers []string
	var data [][]string
	for _, snapshot := range snapshots {
		h, v := r.converter(snapshot)
		if len(headers) == 0 {
			headers = h
		}
		data = append(data, v)
	}
	r.writer.WriteTable(headers, data)
	return nil
}
