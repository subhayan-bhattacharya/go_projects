package main

import "time"

type DbConfig struct {
	Host     string
	Port     int
	Password string
}

type CacheConfig struct {
	Cache      map[string]any
	TimeToLive time.Duration
}

type ConfigStore[C any] struct {
	Store map[string]C
}

func NewConfigStore[C any]() *ConfigStore[C] {
	store := make(map[string]C)
	return &ConfigStore[C]{
		Store: store,
	}
}

func (c *ConfigStore[C]) StoreConfig(name string, config C) error {
	c.Store[name] = config
	return nil
}

func main() {
	dbConfig := DbConfig{
		Host:     "localhost",
		Port:     9000,
		Password: "admin",
	}

	store := NewConfigStore[DbConfig]()
	_ = store.StoreConfig("database", dbConfig)

	cacheConfig := CacheConfig{
		Cache:      make(map[string]any),
		TimeToLive: 5 * time.Minute,
	}

	cache := NewConfigStore[CacheConfig]()
	_ = cache.StoreConfig("database", cacheConfig)

}
