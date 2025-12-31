package provisioner

import (
	customerrors "cloud-provisioner/errors"
	"fmt"
)


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
		return ResourceConfig{}, customerrors.ValidationError{
			Message: "key size is missing from the map",
		}
	}
	regionName, ok := config["region"]
	if !ok {
		return ResourceConfig{}, customerrors.ValidationError{
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
		return ResourceConfig{}, customerrors.ValidationError{
			Message: errorMessage,
		}
	}
}