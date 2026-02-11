package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache[T comparable, V any] struct {
	mu   sync.RWMutex
	data map[T]V
}

func (c *Cache[T, V]) Set(key T, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.data == nil {
		return fmt.Errorf("The cache has not been initialized")
	}
	c.data[key] = value
	return nil
}

func (c *Cache[T, V]) Get(key T) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func (c *Cache[T, V]) Delete(key T) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.data[key]
	if !ok {
		return fmt.Errorf("The key %v is not there in the cache", key)
	}
	delete(c.data, key)
	return nil
}

func (c *Cache[T, V]) Has(key T) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.data[key]
	return ok
}

// it is important to return a pointer and not a copy of the struct
func NewCache[T comparable, V any]() *Cache[T, V] {
	return &Cache[T, V]{
		data: make(map[T]V),
	}
}

type User struct {
	Name     string
	LastName string
	Age      int
}

type Product struct {
	Id    int
	Name  string
	Price int
}

func addUser(c *Cache[string, User], key string, u User) error {
	return c.Set(key, u)
}

func addProduct(c *Cache[int, Product], key int, p Product) error {
	return c.Set(key, p)
}

func main() {
	stringCache := NewCache[string, User]()
	go func() {
		if err := addUser(stringCache, "u1", User{Name: "Sam", LastName: "Lee", Age: 30}); err != nil {
			fmt.Println("set error:", err)
		}
	}()
	go func() {
		if err := addUser(stringCache, "u2", User{Name: "Priya", LastName: "Shah", Age: 27}); err != nil {
			fmt.Println("set error:", err)
		}
	}()

	if v, ok := stringCache.Get("u1"); ok {
		fmt.Printf("u1: %+v\n", v)
	} else {
		fmt.Println("u1 not found")
	}

	fmt.Println("has u2:", stringCache.Has("u2"))

	if err := stringCache.Delete("u2"); err != nil {
		fmt.Println("delete error:", err)
	}
	fmt.Println("has u2 after delete:", stringCache.Has("u2"))

	if err := stringCache.Delete("missing"); err != nil {
		fmt.Println("delete error:", err)
	}

	productCache := NewCache[int, Product]()
	go func() {
		if err := addProduct(productCache, 1, Product{Id: 123, Name: "Pen", Price: 30}); err != nil {
			fmt.Println("set error:", err)
		}
	}()
	go func() {
		if err := addProduct(productCache, 2, Product{Id: 456, Name: "Book", Price: 50}); err != nil {
			fmt.Println("set error:", err)
		}
	}()
	go func() {
		if err := addProduct(productCache, 3, Product{Id: 999, Name: "Laptop", Price: 150}); err != nil {
			fmt.Println("set error:", err)
		}
	}()
	if product, ok := productCache.Get(67); !ok {
		fmt.Println("No such product in cache with key: ", 67)
	} else {
		fmt.Printf("%+v", product)
	}

	time.Sleep(100 * time.Millisecond)
}
