// The marriage between generic data structures and generic functions
// this time we use maps
// Keys have to be comparable in Go , this is a Go rule

package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MapStore[K comparable, V any] map[K]V

func (m MapStore[K, V]) Get(k K) (V, bool) {
	value, ok := m[k]
	if ok {
		return value, true
	} else {
		var zero V
		return zero, false
	}
}

func (m MapStore[K, V]) Set(k K, v V) {
	m[k] = v
}

func (m MapStore[K, V]) Keys() []K {
	mapLength := len(m)
	fmt.Println(mapLength)
	keys := []K{}
	for key, value := range m {
		fmt.Println("Getting value: ", value)
		keys = append(keys, key)
	}
	return keys
}

func (m MapStore[K, V]) Values() []V {
	values := []V{}
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
func FindKeysWithMaxValue[K comparable, V constraints.Ordered](m MapStore[K, V]) []K {
	allValues := m.Values()
	maxValue := allValues[0]
	for _, value := range allValues {
		if value > maxValue {
			maxValue = value
		}
	}
	keysWithMaxValue := []K{}
	for key, value := range m {
		if value == maxValue {
			keysWithMaxValue = append(keysWithMaxValue, key)
		}
	}
	return keysWithMaxValue
}

func main() {
	NewMapStore := MapStore[string, int]{
		"Skoda":    1,
		"Mercedes": 3,
		"Porsche":  1,
		"Hyundai":  1,
	}
	fmt.Println("getting key Merecedes: ", NewMapStore["Mercedes"])
	fmt.Println(NewMapStore.Keys())
	NewMapStore.Set("Volkswagen", 3)
	fmt.Println("Keys which have the max value in the store is : ", FindKeysWithMaxValue(NewMapStore))
}
