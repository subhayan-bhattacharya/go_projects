package customerrors

import "fmt"

type BaseError struct {
	Message string
	Handler string
}

type ResourceAddError struct {
	BaseError
}

func (r ResourceAddError) Error() string {
	return fmt.Sprintf("Resource add error raised by handler %s and message %s", r.Handler, r.Message)
}

type ResourceDeleteError struct {
	BaseError
}

func (r ResourceDeleteError) Error() string {
	return fmt.Sprintf("Resource delete error raised by handler %s and message %s", r.Handler, r.Message)
}

type ResourceUpdateError struct {
	BaseError
}

func (r ResourceUpdateError) Error() string {
	return fmt.Sprintf("Resource update error raised by handler %s and message %s", r.Handler, r.Message)
}
