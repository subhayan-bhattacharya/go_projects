package main

import (
	"errors"
	"fmt"
	"strconv"
)

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

type RegionName string

const (
	UsEast     RegionName = "us-east"
	EuropeWest RegionName = "eu-west"
	AsiaEast   RegionName = "asia-east"
)

type ResourceConfig struct {
	Size   string
	Region RegionName
}

func CreateResourceConfig(config map[string]string) (ResourceConfig, error) {
	size, ok := config["size"]
	if !ok {
		return ResourceConfig{}, ValidationError{
			Message: "key size is missing from the map",
		}
	}
	regionName, ok := config["region"]
	if !ok {
		return ResourceConfig{}, ValidationError{
			Message: "key region is missing from the map",
		}
	}
	region := RegionName(regionName)
	switch RegionName(region) {
	case UsEast, EuropeWest, AsiaEast:
		return ResourceConfig{
			Size:   size,
			Region: region,
		}, nil
	default:
		errorMessage := fmt.Sprintf("The region name %s does not exist", regionName)
		return ResourceConfig{}, ValidationError{
			Message: errorMessage,
		}
	}
}

type Resource struct {
	Type   string
	Name   string
	Config ResourceConfig
}

type Provisioner struct {
	Resources          map[string]Resource
	ResourceTypeLimits map[string]int
}

func NewProvisioner(maxVMs, maxDBs, maxLBs int) *Provisioner {
	resources := make(map[string]Resource, 3)
	resourceLimits := make(map[string]int, 3)
	resourceLimits["VM"] = maxVMs
	resourceLimits["DB"] = maxDBs
	resourceLimits["LB"] = maxLBs
	return &Provisioner{
		Resources:          resources,
		ResourceTypeLimits: resourceLimits,
	}
}

func (p *Provisioner) ProvisionVM(name string, config map[string]string) (Resource, error) {
	if name == "" {
		return Resource{}, ValidationError{
			Message: "Cannot provision a Vm with an empty name",
		}
	}
	_, ok := p.Resources[name]
	if ok {
		return Resource{}, ResourceExistsError{
			ResourceName: name,
		}
	}
	count := 0
	for _, resource := range p.Resources {
		if resource.Type == "VM" {
			count += 1
		}
	}
	if count >= p.ResourceTypeLimits["VM"] {
		return Resource{}, QuotaExceedError{
			ResourceType: "VM",
			Current:      count,
			Limit:        p.ResourceTypeLimits["VM"],
		}
	}
	resourceConfig, err := CreateResourceConfig(config)
	if err != nil {
		return Resource{}, fmt.Errorf("An Validation error is encountered %w", err)
	}
	newVm := Resource{
		Name:   name,
		Type:   "VM",
		Config: resourceConfig,
	}
	p.Resources[name] = newVm
	return newVm, nil
}

func main() {
	p := NewProvisioner(2, 2, 2)
	vmConfig := map[string]string{
		"size":   "large",
		"region": "us-east",
	}
	for i := 0; i <= 2; i++ {
		vmName := "Vm-" + strconv.Itoa(i)
		_, err := p.ProvisionVM(vmName, vmConfig)
		if err != nil {
			var validationError ValidationError
			var quotaExceedError QuotaExceedError
			var resourceExceedError ResourceExistsError
			if errors.As(err, &validationError) {
				fmt.Printf("A validation error occured %s\n", validationError.Message)
			} else if errors.As(err, &quotaExceedError) {
				fmt.Printf("Quota exceeded for %s with limit %d and current %d\n", quotaExceedError.ResourceType, quotaExceedError.Current, quotaExceedError.Limit)
			} else if errors.As(err, &resourceExceedError) {
				fmt.Printf("Resource exceed error for resource name %s", resourceExceedError.ResourceName)
			}
		} else {
			fmt.Println("Resources provisioned successfully")
		}

	}
}
