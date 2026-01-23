package main

import (
	"fmt"
	resources "generics/resources"
)

type Resource interface {
	GetResourceId() string
	GetName() string
}

func main() {
	pod, err := resources.NewPod(
		"123",
		"nginx",
		"starops",
		"running",
	)
	if err != nil {
		fmt.Printf("Error encountered %s", err)
	} else {
		fmt.Printf("The pod details are: %s\n", pod)
	}

	service, err := resources.NewService(
		"234",
		"nginxService",
		"starops",
		"ClusterIp",
	)
	if err != nil {
		fmt.Printf("Error encountered %s", err)
	} else {
		fmt.Printf("The service details are: %s\n", service)
	}
}
