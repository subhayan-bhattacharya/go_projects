package errors

import "fmt"

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	result := "Encountered a validation error " + v.Message
	return result
}

type QuotaExceedError struct {
	ResourceType string
	Current      int
	Limit        int
}

func (q QuotaExceedError) Error() string {
	message := fmt.Sprintf("For the resource type %s the limit is %d and that has been reached", q.ResourceType, q.Limit)
	return message
}

func (q QuotaExceedError) IsRetryable() bool {
	return false
}

type ResourceNotFoundError struct {
	ResourceType string
	ResourceName string
}

func (r ResourceNotFoundError) Error() string {
	message := fmt.Sprintf("For the resource type %s and name %s it is not found", r.ResourceType, r.ResourceName)
	return message
}

func (r ResourceNotFoundError) IsRetryable() bool {
	return false
}

type ResourceExistsError struct {
	ResourceName string
}

func (r ResourceExistsError) Error() string {
	message := fmt.Sprintf("The resource with name %s already exists", r.ResourceName)
	return message
}

func (r ResourceExistsError) IsRetryable() bool {
	return false
}