package main

import (
	customerrors "cloud-provisioner/errors" // Import your errors library
	"cloud-provisioner/provisioner"
	"errors"
	"fmt"
	"strconv"
	// Import your provisioner library
)


func main() {
	p := provisioner.NewProvisioner(2, 2, 2)
	vmConfig := map[string]string{
		"size":   "large",
		"region": "us-east",
	}
	for i := 0; i <= 2; i++ {
		vmName := "Vm-" + strconv.Itoa(i)
		_, err := p.ProvisionVM(vmName, vmConfig)
		if err != nil {
			var validationError customerrors.ValidationError
			var quotaExceedError customerrors.QuotaExceedError
			var resourceExceedError customerrors.ResourceExistsError
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