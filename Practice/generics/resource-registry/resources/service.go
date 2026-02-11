package resources

import (
	"errors"
	"fmt"
)

type ServiceType string

const (
	ClusterIp    ServiceType = "ClusterIp"
	NodePort     ServiceType = "NodePort"
	LoadBalancer ServiceType = "LoadBalancer"
)

func (s ServiceType) isValidService() bool {
	switch s {
	case ClusterIp, NodePort, LoadBalancer:
		return true
	default:
		return false
	}
}

type Service struct {
	Id          string
	Name        string
	Namespace   string
	ServiceType ServiceType
}

func NewService(id string, name string, namespace string, sType string) (*Service, error) {
	serviceType := ServiceType(sType)
	if !serviceType.isValidService() {
		return &Service{}, errors.New("Invalid service provided.")
	}
	return &Service{
		Id:          id,
		Name:        name,
		Namespace:   namespace,
		ServiceType: ServiceType(serviceType),
	}, nil
}

func (s Service) GetResourceId() string {
	return s.Id
}

func (s Service) GetName() string {
	return s.Name
}

func (s Service) String() string {
	return fmt.Sprintf("Service : name %s and namespace %s and type %s", s.Name, s.Namespace, s.ServiceType)
}
