// Simple practice code for defining custom errors in Golang

package main

import "fmt"

type OddNumberError struct {
	number int
}

func (n *OddNumberError) Error() string {
	result := fmt.Sprintf("I hate odd numbers: %d", n.number)
	return result
}

func CheckPositiveInt(number int) (bool, error) {
	if number%2 != 0 {
		return false, &OddNumberError{number: number}
	}
	return true, nil
}

func main() {
	for i := 0; i < 10; i++ {
		_, err := CheckPositiveInt(i)
		if err != nil {
			fmt.Println(err)
		}
	}
}
