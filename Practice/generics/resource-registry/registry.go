// File to store registry information

package main

type Resource interface {
	// interface for defining how should a resource look like
	GetResourceId() string
	GetName() string
}

type Registry[T Resource] struct {
	Resources map[string]T
}

func NewRegistry[T Resource]() *Registry[T] {
	m := make(map[string]T)
	return &Registry[T]{
		Resources: m,
	}
}

func (r *Registry[T]) Add(resource T) {
	r.Resources[resource.GetResourceId()] = resource
}

func (r *Registry[T]) Get(id string) (T, bool) {
	resource, ok := r.Resources[id]
	if !ok {
		var zero T
		return zero, false
	}
	return resource, true
}

func (r *Registry[T]) List() []T {
	result := make([]T, 0, len(r.Resources))
	for _, v := range r.Resources {
		result = append(result, v)
	}
	return result
}

func (r *Registry[T]) Remove(id string) bool {
	_, ok := r.Resources[id]
	if !ok {
		return false
	}
	delete(r.Resources, id)
	return true
}

type FilterFunc[T Resource] func(T) bool

func (r *Registry[T]) Filter(predicate FilterFunc[T]) []T {
	result := make([]T, 0)
	for _, v := range r.Resources {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
