package errors

import "fmt"

type DataReadError struct {
	ErrorMessage string
}

func (d DataReadError) Error() string {
	return fmt.Sprintf("data file could not be read :%s\n", d.ErrorMessage)
}

type CriticalError struct {
	Message string
}

func (c CriticalError) Error() string {
	return fmt.Sprintf("CRITICAL FAILURE: %s", c.Message)
}
