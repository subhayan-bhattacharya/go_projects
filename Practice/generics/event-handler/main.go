package main

import (
	"event-handler/handlers"
	"event-handler/resources"
	"fmt" // Import fmt for printing
)

func main() {
	podLoggingHandler := handlers.LoggingHandler[resources.Pod]{}
	podValidationHandler := handlers.ValidationHandler[resources.Pod]{}
	podEventHandlerChain := handlers.NewEventHandlerChain[resources.Pod]()
	podEventHandlerChain.Add(&podLoggingHandler) // Pass address
	podEventHandlerChain.Add(&podValidationHandler) // Pass address

	pod := resources.Pod{
		Metadata: resources.ObjectMetadata{
			Name:      "my-pod",
			Namespace: resources.StarOps,
			Labels:    map[string]string{"app": "myapp"},
		},
		Status: resources.PodRunning,
		Spec: resources.PodSpec{
			Containers: []resources.Container{
				{
					Name:  "myapp",
					Image: "nginx:latest",
					Ports: []int{80},
				},
			},
		},
	}

	fmt.Println("Attempting to add a valid pod:")
	if err := podEventHandlerChain.OnAdd(pod); err == nil { // Check for nil error
		fmt.Println("Pod added successfully by the event chain.")
	} else {
		fmt.Printf("Failed to add pod by the event chain: %v\n", err)
	}

	// Example of an invalid pod to demonstrate validation
	invalidPod := resources.Pod{
		Metadata: resources.ObjectMetadata{
			Name:      "invalid-pod",
			Namespace: "invalid-namespace", // This namespace is not in StarOps, Hitzler, Fichtner
			Labels:    map[string]string{"app": "badapp"},
		},
		Status: resources.PodUnknown, // This status is valid
		Spec: resources.PodSpec{
			Containers: []resources.Container{}, // No containers, making it invalid
		},
	}

	fmt.Println("\nAttempting to add an invalid pod:")
	if err := podEventHandlerChain.OnAdd(invalidPod); err == nil { // Check for nil error
		fmt.Println("Invalid pod added successfully by the event chain (this should not happen).")
	} else {
		fmt.Printf("Failed to add invalid pod by the event chain (as expected): %v\n", err)
	}

	// Demonstrate OnUpdate
	updatedPod := pod
	updatedPod.Status = resources.PodSucceeded
	fmt.Println("\nAttempting to update a pod:")
	if err := podEventHandlerChain.OnUpdate(pod, updatedPod); err == nil { // Check for nil error
		fmt.Println("Pod updated successfully by the event chain.")
	} else {
		fmt.Printf("Failed to update pod by the event chain: %v\n", err)
	}

	// Demonstrate OnDelete
	fmt.Println("\nAttempting to delete a pod:")
	if err := podEventHandlerChain.OnDelete(pod); err == nil { // Check for nil error
		fmt.Println("Pod deleted successfully by the event chain.")
	} else {
		fmt.Printf("Failed to delete pod by the event chain: %v\n", err)
	}
}
