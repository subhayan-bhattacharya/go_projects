package main

import (
	"errors"
	"fmt"
	"time"
)

type DbConfig struct {
	Host     string
	Port     int
	Password string
}

type CacheConfig struct {
	Cache      map[string]any
	TimeToLive time.Duration
}

var configStore = ConfigStore{
	Store: make(map[string]any),
}

type ConfigStore struct {
	Store map[string]any
}

func RegisterConfig[C any](name string, config C) error {
	_, ok := configStore.Store[name]
	if ok {
		return errors.New("the config already exists")
	}
	configStore.Store[name] = config
	return nil
}

func GetConfig[C any](name string) (C, error) {
	var zeroC C
	c, ok := configStore.Store[name]
	if !ok {
		return zeroC, fmt.Errorf("no such config with name %s", name)
	}
	config, ok := c.(C)
	if !ok {
		return zeroC, errors.New("the types do not match.")
	}
	return config, nil
}

func main() {
	dbConfig := DbConfig{
		Host:     "localhost",
		Port:     9000,
		Password: "admin",
	}

	cacheConfig := CacheConfig{
		Cache:      make(map[string]any),
		TimeToLive: 5 * time.Minute,
	}

	_ = RegisterConfig[DbConfig]("database", dbConfig)
	_ = RegisterConfig[CacheConfig]("cache", cacheConfig)
	config, _ := GetConfig[DbConfig]("database")
	fmt.Printf("what is the value of host %s \n", config.Host)
}
