package plugin

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

import internalErrors "plugin-registry/errors"

type PipelineOption[T any] func(pipeline *PluginPipeline[T])

func WithTimeOut[T any](duration time.Duration) PipelineOption[T] {
	return func(pipeline *PluginPipeline[T]) {
		pipeline.timeout = duration
	}
}

type PluginPipeline[T any] struct {
	Plugins []PipelinePlugin[T]
	timeout time.Duration
}

func NewPluginPipeline[T any](opts ...PipelineOption[T]) *PluginPipeline[T] {
	pluginPipeline := &PluginPipeline[T]{
		Plugins: make([]PipelinePlugin[T], 0),
	}
	for _, opt := range opts {
		opt(pluginPipeline)
	}
	return pluginPipeline
}

func (p *PluginPipeline[T]) AddPlugin(plugin PipelinePlugin[T]) {
	p.Plugins = append(p.Plugins, plugin)
}

func (p *PluginPipeline[T]) Execute(ctx context.Context, data T) error {
	var copiedData T
	copiedData = *data
	var wg sync.WaitGroup
	newCtx := ctx
	if p.timeout > 0 {
		var cancel context.CancelFunc
		newCtx, cancel = context.WithTimeout(ctx, p.timeout)
		defer cancel()
	}
	for _, plugin := range p.Plugins {
		select {
		case <-newCtx.Done():
			return fmt.Errorf("the context was cancelled because of : %w", newCtx.Err())
		default:
		}
		wg.Add(1)
		go func() error {
			defer wg.Done()
			var criticalError internalErrors.CriticalError
			err := plugin.Execute(newCtx, copiedData)
			if errors.As(err, &criticalError) {
				fmt.Println("could not continue further as we have a critical error")
				return err
			}
			return nil
		}()

	}
	wg.Wait()
	*data = copiedData
	return nil
}
