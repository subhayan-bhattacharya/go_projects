package main

import "fmt"

type ObjectMeta struct {
	Name      string
	Namespace string
	Labels    map[string]string
}

type MetadataCarrier interface {
	GetMetadata() ObjectMeta
}

type MetadataAccessor struct {
	Metadata ObjectMeta
}

func (m MetadataAccessor) GetMetadata() ObjectMeta {
	return m.Metadata
}

type Pod struct {
	MetadataAccessor
	// anonymous struct
	Spec struct {
		NodeName string
	}
}

type Deployment struct {
	MetadataAccessor
	Spec struct {
		Replicas int
	}
}

type FilterFunc[T MetadataCarrier] func(T) bool

func FilterResources[T MetadataCarrier](resources []T, predicate FilterFunc[T]) []T {
	result := make([]T, 0)
	for _, v := range resources {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Create some test data
	pods := []Pod{
		{MetadataAccessor: MetadataAccessor{Metadata: ObjectMeta{Name: "pod1", Namespace: "default"}}, Spec: struct{ NodeName string }{NodeName: "node-1"}},
		{MetadataAccessor: MetadataAccessor{Metadata: ObjectMeta{Name: "pod2", Namespace: "default"}}, Spec: struct{ NodeName string }{NodeName: "node-2"}},
		{MetadataAccessor: MetadataAccessor{Metadata: ObjectMeta{Name: "pod3", Namespace: "kube-system"}}, Spec: struct{ NodeName string }{NodeName: "node-1"}},
	}

	// Try your filter!
	filtered := FilterResources(pods, func(p Pod) bool {
		return p.Spec.NodeName == "node-1"
	})

	fmt.Printf("Found %d pods on node-1\n", len(filtered))
	for _, p := range filtered {
		fmt.Printf("  - %s/%s\n", p.Metadata.Namespace, p.Metadata.Name)
	}
}
