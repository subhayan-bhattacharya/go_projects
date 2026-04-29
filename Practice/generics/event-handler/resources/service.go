package resources

import "fmt"

// ServiceType defines the type of service.
type ServiceType string

const (
	ClusterIP    ServiceType = "ClusterIP"
	NodePort     ServiceType = "NodePort"
	LoadBalancer ServiceType = "LoadBalancer"
)

func (s ServiceType) IsValid() (bool, error) {

	switch s {
	case ClusterIP, NodePort, LoadBalancer:
		return true, nil
	}
	return false, fmt.Errorf("The type of the service %s is not valid", s)
}

// ServicePort defines a port for a service.
type ServicePort struct {
	Port       int32
	TargetPort int32
	Protocol   string // e.g., "TCP", "UDP"
}

// ServiceSpec defines the desired state of Service.
type ServiceSpec struct {
	Type     ServiceType
	Ports    []ServicePort
	Selector map[string]string // Selects pods that are part of this service
}

// Service is a resource that exposes an application running on a set of Pods as a network service.
type Service struct {
	Metadata ObjectMetadata
	Spec     ServiceSpec
}

func (s Service) Log() string {
	return fmt.Sprintf("Service: %s, inside namespace %s, Type: %s", s.Metadata.Name, s.Metadata.Namespace, s.Spec.Type)
}

func (s Service) IsValid() (bool, error) {
	_, namespaceError := s.Metadata.Namespace.IsValid()
	if namespaceError != nil {
		return false, namespaceError
	}
	_, typeError := s.Spec.Type.IsValid()
	if typeError != nil {
		return false, typeError
	}
	portsValid := len(s.Spec.Ports) > 0
	if !portsValid {
		return false, fmt.Errorf("The number of ports %d is not valid", len(s.Spec.Ports))

	}
	return true, nil
}
