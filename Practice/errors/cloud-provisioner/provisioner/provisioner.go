package provisioner

import (
	customerrors "cloud-provisioner/errors"
	"fmt"
)

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
		return Resource{}, customerrors.ValidationError{
			Message: "Cannot provision a Vm with an empty name",
		}
	}
	_, ok := p.Resources[name]
	if ok {
		return Resource{}, customerrors.ResourceExistsError{
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
		return Resource{}, customerrors.QuotaExceedError{
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