package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Wrapper struct {
	UserIds      []int        `json:"user_ids"`
	CountryCodes []string     `json:"country_codes"`
	Permissions  []Permission `json:"permissions"`
}

type Permission struct {
	Resource string `json:"resource"`
	Level    int    `json:"level"`
}

// FromSlice builds a Set from a slice, silently skipping duplicate elements.
func FromSlice[T comparable](data []T) *Set[T] {
	initial := NewSet[T]()
	return Reduce(data, initial, func(s *Set[T], data T) *Set[T] {
		_ = s.Add(data)
		return s
	})
}

// ToSlice returns the Set elements as a slice. Order is not guaranteed (map iteration).
func ToSlice[T comparable](input *Set[T]) []T {
	var result []T
	for key, _ := range input.data {
		result = append(result, key)
	}
	return result
}

// Reduce folds data into an accumulator of a different type U, applying reducer on each element.
func Reduce[T, U any](data []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, singleData := range data {
		result = reducer(result, singleData)
	}
	return result
}

// Filter returns a new slice containing only elements for which filterer returns true.
func Filter[T comparable](input []T, filterer func(T) bool) []T {
	var result []T
	for _, value := range input {
		if filterer(value) {
			result = append(result, value)
		}
	}
	return result
}

// struct{} as the value type uses zero memory — we only care about key presence.
type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet returns an empty Set with an initialised backing map.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}

// Add returns an error instead of silently overwriting so callers can detect duplicates.
func (s *Set[T]) Add(element T) error {
	_, ok := s.data[element]
	if ok {
		return errors.New("the element to add exists")
	}
	s.data[element] = struct{}{}
	return nil
}

// Remove deletes element from the Set. Returns an error if the element is not present.
func (s *Set[T]) Remove(element T) error {
	_, ok := s.data[element]
	if !ok {
		return errors.New("the element to remove does not exist")
	}
	delete(s.data, element)
	return nil
}

// Contains reports whether element is present in the Set.
func (s *Set[T]) Contains(element T) bool {
	_, ok := s.data[element]
	return ok
}

// Intersect reuses ToSlice+Filter to avoid duplicating iteration logic.
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	sliceS := ToSlice[T](s)
	intersected := Filter[T](sliceS, func(data T) bool {
		return other.Contains(data)
	})
	for _, i := range intersected {
		_ = result.Add(i)
	}
	return result
}

// Union adds both sets; duplicate Add errors are intentionally ignored since overlap is expected.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range other.data {
		_ = result.Add(key)
	}
	for key := range s.data {
		_ = result.Add(key)
	}
	return result
}

func readData() (Wrapper, error) {
	var results Wrapper
	data, err := os.ReadFile("data.json")
	if err != nil {
		return results, fmt.Errorf("could not read data %v", err)
	}
	err = json.Unmarshal(data, &results)
	if err != nil {
		return results, fmt.Errorf("could not unmarshall %v", err)
	}
	return results, nil
}

func main() {
	data, err := readData()
	if err != nil {
		fmt.Println(err)
	}
	userIdSet := FromSlice[int](data.UserIds)
	countryCodesSet := FromSlice[string](data.CountryCodes)
	permissionSet := FromSlice[Permission](data.Permissions)

	fmt.Println(ToSlice[int](userIdSet))
	fmt.Println(ToSlice[string](countryCodesSet))
	fmt.Println(ToSlice[Permission](permissionSet))
}
