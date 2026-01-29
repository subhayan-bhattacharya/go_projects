package main

import (
	"fmt"
	resources "generics/resources"
)

func main() {
	podRegistry := NewRegistry[resources.Pod]()
	pod1 := resources.Pod{Id: "1", Name: "nginx", Namespace: "default", Status: resources.Status("running")}
	pod2 := resources.Pod{Id: "2", Name: "redis", Namespace: "default", Status: resources.Status("pending")}
	pod3 := resources.Pod{Id: "3", Name: "postgres", Namespace: "prod", Status: resources.Status("running")}
	pod4 := resources.Pod{Id: "4", Name: "mysql", Namespace: "prod", Status: resources.Status("running")}
	pod5 := resources.Pod{Id: "5", Name: "mongodb", Namespace: "prod", Status: resources.Status("failed")}
	podRegistry.Add(pod1)
	podRegistry.Add(pod2)
	podRegistry.Add(pod3)
	podRegistry.Add(pod4)
	podRegistry.Add(pod5)

	allPods := podRegistry.List()
	fmt.Println("Getting all pods list")
	for _, pod := range allPods {
		fmt.Printf("Pod : %s\n", pod)
	}

	runningPods := podRegistry.Filter(func(p resources.Pod) bool {
		return p.Status == resources.Status("running")
	})
	fmt.Printf("Total running pods %d\n", len(runningPods))
	fmt.Println("Getting all running pods..")
	for _, p := range runningPods {
		fmt.Printf("%s\n", p)
	}

	failedPods := podRegistry.Filter(func(p resources.Pod) bool {
		return p.Status == resources.Status("failed")
	})
	fmt.Println("The number of failed pods are %d", len(failedPods))
	if len(failedPods) > 0 {
		fmt.Println("Deleting failed pods...")
		for _, p := range failedPods {
			podRegistry.Remove(p.Id)
		}
	}
	fmt.Println("After deleting failed pods the pods remaining ...")
	for _, p := range podRegistry.List() {
		fmt.Println(p)
	}
}
