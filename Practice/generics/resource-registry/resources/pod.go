package resources

import (
	"errors"
	"fmt"
)

type Status string

const (
	Running Status = "running"
	Failed  Status = "failed"
	Pending Status = "pending"
)

func (s Status) IsValidStatus() bool {
	switch s {
	case Running, Pending, Failed:
		return true
	default:
		return false
	}
}

type Pod struct {
	Id        string
	Name      string
	Namespace string
	Status    Status
}

func (p Pod) GetResourceId() string {
	return p.Id
}

func (p Pod) GetName() string {
	return p.Name
}

func (p Pod) String() string {
	return fmt.Sprintf("Pod: name %s and namespace %s and status %s", p.Name, p.Namespace, p.Status)

}

func NewPod(id string, name string, namespace string, status string) (*Pod, error) {
	podStatus := Status(status)
	if !podStatus.IsValidStatus() {
		return &Pod{}, errors.New("Invalid pod status")
	}
	return &Pod{
		Id:        id,
		Name:      name,
		Namespace: namespace,
		Status:    podStatus,
	}, nil
}
