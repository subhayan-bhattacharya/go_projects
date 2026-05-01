package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	UserIds      []int        `json:"user_ids"`
	CountryCodes []string     `json:"country_codes"`
	Permissions  []Permission `json:"permissions"`
}

type Permission struct {
	Resource string `json:"resource"`
	Level    int    `json:"level"`
}

// Set is a generic set backed by map[T]struct{}. Duplicate elements are silently ignored.
type Set[T comparable] struct {
	Items map[T]struct{}
}

// NewSet returns an empty Set with an initialised backing map.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		Items: make(map[T]struct{}),
	}
}

// Add inserts item into the Set. No-ops if item is already present.
func (s *Set[T]) Add(item T) {
	if _, ok := s.Items[item]; !ok {
		s.Items[item] = struct{}{}
	}
}

// Difference returns a new Set with elements in s that are not in other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.Items {
		_, ok := other.Items[item]
		if !ok {
			result.Add(item)
		}
	}
	return result
}

// Union returns a new Set containing all elements from both s and other.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.Items {
		result.Add(item)
	}
	for item := range other.Items {
		result.Add(item)
	}
	return result
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.Items[item]
	return ok
}

func (s *Set[T]) Remove(item T) {
	delete(s.Items, item)
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.Items {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

func (s *Set[T]) Len() int {
	return len(s.Items)
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Set[T]) Clear() {
	clear(s.Items)
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for item := range s.Items {
		_, ok := other.Items[item]
		if !ok {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	for item := range other.Items {
		_, ok := s.Items[item]
		if !ok {
			return false
		}
	}
	return true
}

func FromSlice[T comparable](input []T) *Set[T] {
	result := NewSet[T]()
	return Reduce(input, result, func(data T, result *Set[T]) *Set[T] {
		result.Add(data)
		return result
	})
}

func ToSlice[T comparable](s *Set[T]) []T {
	result := make([]T, 0)
	for item := range s.Items {
		result = append(result, item)
	}
	return result
}

func Reduce[T, U any](data []T, initial U, reducer func(T, U) U) U {
	result := initial
	for _, item := range data {
		initial = reducer(item, initial)
	}
	return result
}

func getData() (Data, error) {
	content, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file")
		return Data{}, err
	}
	var data Data
	if err := json.Unmarshal(content, &data); err != nil {
		fmt.Println("Error unmarshalling JSON")
		return Data{}, err
	}
	return data, nil
}

func main() {
	data, _ := getData()
	deDuplicatedUsers := FromSlice(data.UserIds)
	fmt.Println(ToSlice(deDuplicatedUsers))
	deDuplicatedCountryCodes := FromSlice(data.CountryCodes)
	fmt.Println(ToSlice(deDuplicatedCountryCodes))
	deDuplicatedPermissions := FromSlice(data.Permissions)
	fmt.Println(ToSlice(deDuplicatedPermissions))
	itemsCount := len(data.UserIds)
	left := FromSlice(data.UserIds[:itemsCount/2])
	right := FromSlice(data.UserIds[itemsCount/2:])
	fmt.Println("Left: ", ToSlice(left))
	fmt.Println("Right: ", ToSlice(right))
	fmt.Println(ToSlice(left.Intersect(right)))
	fmt.Println(ToSlice(left.Union(right)))
	fmt.Println(left.IsSubset(right))
	fmt.Println(left.IsSuperset(right))
}
