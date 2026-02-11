package resources

import "fmt"

type PodStatus string

const (
	PodPending   PodStatus = "Pending"
	PodRunning   PodStatus = "Running"
	PodSucceeded PodStatus = "Succeeded"
	PodFailed    PodStatus = "Failed"
	PodUnknown   PodStatus = "Unknowm"
)

func (s PodStatus) IsValid() (bool, error) {
	switch s {
	case PodPending, PodRunning, PodSucceeded, PodFailed, PodUnknown:
		return true, nil
	}
	return false, fmt.Errorf("The status of the pod %s is not valid", s)
}

type Container struct {
	Name  string
	Image string
	Ports []int
}

type PodSpec struct {
	Containers []Container
}

type Pod struct {
	Metadata ObjectMetadata
	Status   PodStatus
	Spec     PodSpec
}

func (p Pod) Log() string {
	return fmt.Sprintf("Pod: %s, inside namespace %s, Status: %s", p.Metadata.Name, p.Metadata.Namespace, p.Status)
}

func (p Pod) IsValid() (bool, error) {

	_, namespaceError := p.Metadata.Namespace.IsValid()
	if namespaceError != nil {
		return false, namespaceError
	}
	_, statusError := p.Status.IsValid()
	if statusError != nil {
		return false, statusError
	}
	containersValid := len(p.Spec.Containers) > 0
	if !containersValid {
		return false, fmt.Errorf("The number of containers %d is not valid", len(p.Spec.Containers))
	}
	return true, nil
}
