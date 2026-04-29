package plugin

import (
	"context"
	"fmt"
	"sync"
)
import datamodel "plugin-registry/datamodel"
import errors "plugin-registry/errors"

type PipelinePlugin[T any] interface {
	Name() string
	Execute(ctx context.Context, data T) error
}

type Sanitizer[T datamodel.Sanitizable] struct{}

func (s Sanitizer[T]) Name() string {
	return "DataSanitizer"
}

func (s Sanitizer[T]) Execute(ctx context.Context, data T) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("the context was cancelled before the sanitizer could be run %w", ctx.Err())
	default:
	}
	value := data.GetValue()
	tags := data.GetTags()
	if value < -273.15 {
		message := fmt.Sprintf("the value %f is less than -273", value)
		return errors.CriticalError{Message: message}
	}
	if tags == nil {
		tags := make(map[string]string)
		tags["Sanitized"] = "true"
		data.SetTags(tags)
	}
	return nil
}

type MemoryStore[T datamodel.Sanitizable] struct {
	Db map[string]T
	mu sync.RWMutex
}

func NewMemoryStore[T datamodel.Sanitizable]() *MemoryStore[T] {
	db := make(map[string]T)
	return &MemoryStore[T]{
		Db: db,
	}
}

func (m *MemoryStore[T]) Name() string {
	return "MemoryStore"
}

func (m *MemoryStore[T]) Get(signaId string) T {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.Db[signaId]
}

func (m *MemoryStore[T]) Execute(ctx context.Context, data T) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("the context was cancelled before the data could be persisted %w", ctx.Err())
	default:
	}
	sensorId := data.GetId()
	m.mu.Lock()
	m.Db[sensorId] = data
	defer m.mu.Unlock()
	return nil
}
