package main

import (
	"errors"
	"fmt"
	"maps"
	"time"
)

type Resource interface {
	GetName() string
	GetKind() string
	GetGeneration() int64
}

type Comparable interface {
	Resource
	Equals(other Resource) bool
}

type Reconcilable interface {
	Comparable
	NeedsUpdate(desired Resource) (bool, error)
	Update(desired Resource) error
}

type ResourceMetadata struct {
	Name       string
	Namespace  string
	Generation int64
	Labels     map[string]string
}

func (r ResourceMetadata) GetName() string {
	return r.Name
}

type Phase string

const (
	PhasePending Phase = "Pending"
	PhaseRunning Phase = "Running"
	PhaseFailed  Phase = "Failed"
)

type ResourceStatus struct {
	Phase        Phase
	LastModified time.Time
	Message      string
}

type ServiceType string

const (
	ClusterIpServiceType    ServiceType = "ClusterIP"
	NodeportServiceType     ServiceType = "NodePort"
	LoadBalancerServiceType ServiceType = "LoadBalancer"
)

type ServiceResource struct {
	ResourceMetadata
	ResourceStatus
	Port       int
	TargetPort int
	Type       ServiceType
}

func (s *ServiceResource) GetKind() string {
	return "Service"
}
func (s *ServiceResource) GetGeneration() int64 {
	return s.Generation
}
func (s *ServiceResource) Equals(other Resource) bool {
	if o, ok := other.(*ServiceResource); ok {
		return s.Name == o.Name && s.Namespace == o.Namespace
	}
	return false
}
func (s *ServiceResource) NeedsUpdate(desired Resource) (bool, error) {
	if o, ok := desired.(*ServiceResource); !ok {
		return false, nil
	} else {
		if !s.Equals(desired) {
			return false, errors.New("The resources to be compared are not the same")
		}
		return s.Port != o.Port || s.TargetPort != o.TargetPort || s.Type != o.Type, nil

	}
}
func (s *ServiceResource) Update(desired Resource) error {
	d, ok := desired.(*ServiceResource)
	if ok {
		s.Namespace = d.Namespace
		s.Port = d.Port
		s.TargetPort = d.TargetPort
		s.Type = d.Type
		s.Generation += 1
		s.Phase = PhaseRunning
		s.Message = "New configuration applied"
		s.LastModified = time.Now()
		return nil
	}
	return errors.New("Could not update the pod resource")
}

type ConfigMapResource struct {
	ResourceMetadata
	ResourceStatus
	Data map[string]string
}

func (c *ConfigMapResource) GetKind() string {
	return "ConfigMap"
}
func (c *ConfigMapResource) GetGeneration() int64 {
	return c.Generation
}
func (c *ConfigMapResource) Equals(other Resource) bool {
	o, ok := other.(*ConfigMapResource)
	if ok {
		return c.Name == o.Name && c.Namespace == o.Namespace
	}
	return false
}

func (c *ConfigMapResource) NeedsUpdate(desired Resource) (bool, error) {
	o, ok := desired.(*ConfigMapResource)
	if !ok {
		return false, nil
	}
	if c.Name == o.Name && c.Namespace == o.Namespace {
		return !maps.Equal(c.Data, o.Data), nil
	}
	return false, nil
}

func (c *ConfigMapResource) Update(desired Resource) error {
	o, ok := desired.(*ConfigMapResource)
	if ok {
		c.Data = maps.Clone(o.Data)
		c.Generation++
		c.Phase = PhaseRunning
		c.Message = "Configuration updated"
		c.LastModified = time.Now()
		return nil
	}
	return errors.New("The resource to update should be a configmap resource")
}

type PodResource struct {
	ResourceMetadata
	ResourceStatus
	Image        string
	Replica      int
	RestartCount int
}

func (p *PodResource) GetKind() string {
	return "Pod"
}
func (p *PodResource) GetGeneration() int64 {
	return p.Generation
}
func (p *PodResource) Equals(other Resource) bool {
	if o, ok := other.(*PodResource); ok {
		return p.Name == o.Name && p.Namespace == o.Namespace
	}
	return false
}
func (p *PodResource) NeedsUpdate(desired Resource) (bool, error) {
	if o, ok := desired.(*PodResource); !ok {
		return false, nil
	} else {
		if !p.Equals(desired) {
			return false, nil
		}
		return p.Image != o.Image || p.Replica != o.Replica, nil
	}
}
func (p *PodResource) Update(desired Resource) error {
	d, ok := desired.(*PodResource)
	if ok {
		p.Namespace = d.Namespace
		p.Image = d.Image
		p.Replica = d.Replica
		p.Generation += 1
		p.Phase = PhaseRunning
		p.Message = "New configuration applied"
		p.LastModified = time.Now()
		return nil
	}
	return errors.New("Could not update the pod resource")
}

type Action string

const (
	CreatedAction  Action = "Created"
	UpdatedAction  Action = "Updated"
	NoChangeAction Action = "NoChange"
	FailedAction   Action = "Failed"
	DeletedAction  Action = "Deleted"
)

type ReconcilliationResult struct {
	ResourceName string
	Action       Action
	Error        error
	Duration     time.Duration
}

func ReconcileResources(desired []Reconcilable, actual []Reconcilable) []ReconcilliationResult {
	result := []ReconcilliationResult{}
	actualSeenResources := map[string]Reconcilable{}
	for _, resource := range actual {
		nameAndKind := resource.GetName() + ":" + resource.GetKind()
		actualSeenResources[nameAndKind] = resource
	}
	for _, desiredResource := range desired {
		nameAndKindDesired := desiredResource.GetName() + ":" + desiredResource.GetKind()
		actualResource, ok := actualSeenResources[nameAndKindDesired]
		if ok {
			var r ReconcilliationResult
			startTime := time.Now()
			ok, _ := actualResource.NeedsUpdate(desiredResource)
			if ok {
				err := actualResource.Update(desiredResource)
				endTime := time.Now()
				if err != nil {
					r = ReconcilliationResult{
						ResourceName: actualResource.GetName(),
						Action:       FailedAction,
						Error:        err,
						Duration:     endTime.Sub(startTime),
					}
				} else {
					r = ReconcilliationResult{
						ResourceName: actualResource.GetName(),
						Action:       UpdatedAction,
						Error:        err,
						Duration:     endTime.Sub(startTime),
					}
				}
			} else {
				endTime := time.Now()
				r = ReconcilliationResult{
					ResourceName: actualResource.GetName(),
					Action:       NoChangeAction,
					Error:        nil,
					Duration:     endTime.Sub(startTime),
				}
			}
			delete(actualSeenResources, nameAndKindDesired)
			result = append(result, r)
		} else {
			result = append(result, ReconcilliationResult{
				ResourceName: desiredResource.GetName(),
				Action:       CreatedAction,
				Error:        nil,
				Duration:     0,
			})
		}
	}
	for _, value := range actualSeenResources {
		r := ReconcilliationResult{
			ResourceName: value.GetName(),
			Action:       DeletedAction,
			Error:        nil,
			Duration:     0,
		}
		result = append(result, r)
	}
	return result
}


func main() {
	// Desired state
    desired := []Reconcilable{
        &PodResource{
            ResourceMetadata: ResourceMetadata{Name: "nginx", Namespace: "default"},
            Image: "nginx:1.21",
            Replica: 3,
        },
        &ServiceResource{
            ResourceMetadata: ResourceMetadata{Name: "api", Namespace: "default"},
            Port: 80,
        },
    }
    
    // Actual state (only nginx exists, different image)
    actual := []Reconcilable{
        &PodResource{
            ResourceMetadata: ResourceMetadata{Name: "nginx", Namespace: "default"},
            Image: "nginx:1.20",  // Old version
            Replica: 3,
        },
    }
    
    results := ReconcileResources(desired, actual)
    
    for _, r := range results {
        fmt.Printf("Resource: %s, Action: %s, Error: %v\n", 
            r.ResourceName, r.Action, r.Error)
    }
}
