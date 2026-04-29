package main

import (
	"errors"
	"fmt"
)

type Entity interface {
	GetId() string
	Describe() string
	Validate() (bool, error)
}

type User struct {
	Id    string
	Email string
}

func (u User) GetId() string {
	return u.Id
}

func (u User) Describe() string {
	return fmt.Sprintf("id %s and email %s", u.Id, u.Email)
}

func (u User) Validate() (bool, error) {
	if u.Email == "" {
		return false, errors.New("The email cannot be empty")
	}
	return true, nil
}

type DataMapper[T Entity] struct {
	Database map[string]T
}

func NewDataMapper[T Entity]() *DataMapper[T] {
	database := make(map[string]T)
	return &DataMapper[T]{
		Database: database,
	}
}

func (e *DataMapper[T]) Save(item T) error {
	validated, _ := item.Validate()
	if !validated {
		return fmt.Errorf("The item %v could not be validated", item)
	}
	e.Database[item.GetId()] = item
	return nil
}

func (e *DataMapper[T]) Get(id string) (T, error) {
	var zero T
	item, ok := e.Database[id]
	if !ok {
		return zero, fmt.Errorf("No item with the id %s could be found", id)
	}
	return item, nil
}

func (e *DataMapper[T]) Filter(predicate func(T) bool) []T {
	results := make([]T, 0)
	for _, item := range e.Database {
		if predicate(item) {
			results = append(results, item)
		}
	}
	return results
}

func main() {
	subhayan := User{
		Id:    "123",
		Email: "subhayan.here@gmail.com",
	}

	dataMapper := NewDataMapper[User]()
	error := dataMapper.Save(subhayan)
	if error != nil {
		fmt.Println("Could not save the user")
	}
}
