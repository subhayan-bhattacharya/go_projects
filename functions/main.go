package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Marshaller interface {
	Marshall() ([]byte, error)
}

type UnMarshaller interface {
	UnMarshall([]byte) error
}

func Load[T UnMarshaller](filename string, data T) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}
	unmarshallError := data.UnMarshall(bytes)
	if unmarshallError != nil {
		return fmt.Errorf("could not unmarshall: %w", err)
	}
	return nil
}

func Save[T Marshaller](filename string, data T) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create a file %w", err)
	}
	defer file.Close()
	dataToWrite, err := data.Marshall()
	if err != nil {
		return fmt.Errorf("something went wrong with Marshalling..%w", err)
	}
	_, err = file.Write(dataToWrite)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}

type User struct {
	Name string
	Age  int
}

func (u *User) Marshall() ([]byte, error) {
	data, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		newError := fmt.Errorf("we have a marshall error %w", err)
		return []byte{}, newError
	}
	return data, nil
}

func (u *User) UnMarshall(data []byte) error {
	err := json.Unmarshal(data, u)
	if err != nil {
		return fmt.Errorf("there is an unmarshalling error: %w", err)
	}
	return nil
}

func main() {
	data := User{
		Name: "Subhayan",
		Age:  41,
	}
	err := Save("data.json", &data)
	if err != nil {
		fmt.Println("A problem occured")
		fmt.Println("The actual issue is : ", errors.Unwrap(err))
	}
}
