package main

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
)

type Resource struct {
	CPU    int
	Memory int
}

type Node struct {
	Name          string
	Capacity      Resource
	Ready         bool
	Unschedulable bool
}

type Pod struct {
	Namespace string
	PodName   string
	NodeName  string
	Request   Resource
}

type Placement struct {
	Namespace string
	PodName   string
	NodeName  string
	Reason    string
}

func schedulePods(
	nodes []Node, runningPods []Pod, pendingPods []Pod,
	quotas map[string]Resource,
) []Placement {
	var placements []Placement
	nodeResourceConsumption := map[string]Resource{}
	clonedQuotas := maps.Clone(quotas)
	for _, pod := range runningPods {
		podNamespace := pod.Namespace
		podNodeName := pod.NodeName
		quota, ok := clonedQuotas[podNamespace]
		if ok {
			quota.CPU -= pod.Request.CPU
			quota.Memory -= pod.Request.Memory
			clonedQuotas[podNamespace] = quota
		}
		nodeConsumption, ok := nodeResourceConsumption[podNodeName]
		if ok {
			nodeConsumption.CPU += pod.Request.CPU
			nodeConsumption.Memory += pod.Request.Memory
			nodeResourceConsumption[podNodeName] = nodeConsumption
		} else {
			nodeResourceConsumption[podNodeName] = pod.Request
		}
	}

	for _, pendingPod := range pendingPods {
		namespaceQuota, ok := clonedQuotas[pendingPod.Namespace]
		if ok {
			if namespaceQuota.Memory < pendingPod.Request.
				Memory || namespaceQuota.CPU < pendingPod.Request.CPU {
				placement := Placement{
					Namespace: pendingPod.Namespace,
					PodName:   pendingPod.PodName,
					NodeName:  "",
					Reason:    "QuotaExceeded",
				}
				placements = append(placements, placement)
			}
		}
		var eligibleNodes []Node
		for _, node := range nodes {
			if node.Ready && !node.Unschedulable {
				if node.Capacity.CPU >= pendingPod.Request.CPU && node.
					Capacity.Memory >= pendingPod.Request.Memory {
					eligibleNodes = append(eligibleNodes, node)
				}
			}
		}
		var placement Placement
		if len(eligibleNodes) == 0 {
			placement = Placement{
				Namespace: pendingPod.Namespace,
				PodName:   pendingPod.PodName,
				NodeName:  "",
				Reason:    "NoEligibleNode",
			}
		} else {
			slices.SortFunc(
				eligibleNodes, func(nodeA, nodeB Node) int {
					return cmp.Or(
						cmp.Compare(
							nodeResourceConsumption[nodeA.Name].CPU,
							nodeResourceConsumption[nodeB.Name].CPU,
						),
						cmp.Compare(
							nodeResourceConsumption[nodeA.Name].Memory,
							nodeResourceConsumption[nodeB.Name].Memory,
						),
						cmp.Compare(nodeA.Name, nodeB.Name),
					)
				},
			)
			placement = Placement{
				Namespace: pendingPod.Namespace,
				PodName:   pendingPod.PodName,
				NodeName:  eligibleNodes[0].Name,
				Reason:    "",
			}
		}
		placements = append(placements, placement)
	}
	return placements
}

func main() {
	tests := []struct {
		name        string
		nodes       []Node
		runningPods []Pod
		pendingPods []Pod
		quotas      map[string]Resource
		expected    string
	}{
		{
			name: "quota exceeded must produce one rejection",
			nodes: []Node{
				{
					Name: "node-a", Capacity: Resource{CPU: 8, Memory: 8},
					Ready: true,
				},
			},
			pendingPods: []Pod{
				{
					Namespace: "team-a", PodName: "pod-1",
					Request: Resource{CPU: 3, Memory: 3},
				},
			},
			quotas:   map[string]Resource{"team-a": {CPU: 2, Memory: 2}},
			expected: "[{Namespace:team-a PodName:pod-1 NodeName: Reason:QuotaExceeded}]",
		},
		{
			name: "running pods must reduce available node capacity",
			nodes: []Node{
				{
					Name: "node-a", Capacity: Resource{CPU: 4, Memory: 4},
					Ready: true,
				},
			},
			runningPods: []Pod{
				{
					Namespace: "team-a", PodName: "running", NodeName: "node-a",
					Request: Resource{CPU: 3, Memory: 3},
				},
			},
			pendingPods: []Pod{
				{
					Namespace: "team-a", PodName: "pod-1",
					Request: Resource{CPU: 2, Memory: 2},
				},
			},
			quotas:   map[string]Resource{"team-a": {CPU: 10, Memory: 10}},
			expected: "[{Namespace:team-a PodName:pod-1 NodeName: Reason:NoEligibleNode}]",
		},
		{
			name: "pending pods must be accounted for sequentially",
			nodes: []Node{
				{
					Name: "node-a", Capacity: Resource{CPU: 4, Memory: 4},
					Ready: true,
				},
			},
			pendingPods: []Pod{
				{
					Namespace: "team-a", PodName: "pod-1",
					Request: Resource{CPU: 3, Memory: 3},
				},
				{
					Namespace: "team-a", PodName: "pod-2",
					Request: Resource{CPU: 3, Memory: 3},
				},
			},
			quotas:   map[string]Resource{"team-a": {CPU: 10, Memory: 10}},
			expected: "pod-1 on node-a, then pod-2 rejected with NoEligibleNode",
		},
		{
			name: "best fit must minimize remaining CPU",
			nodes: []Node{
				{
					Name: "node-a", Capacity: Resource{CPU: 10, Memory: 10},
					Ready: true,
				},
				{
					Name: "node-b", Capacity: Resource{CPU: 6, Memory: 10},
					Ready: true,
				},
			},
			pendingPods: []Pod{
				{
					Namespace: "team-a", PodName: "pod-1",
					Request: Resource{CPU: 3, Memory: 3},
				},
			},
			quotas:   map[string]Resource{"team-a": {CPU: 10, Memory: 10}},
			expected: "pod-1 on node-b",
		},
	}

	for _, test := range tests {
		actual := schedulePods(
			test.nodes, test.runningPods, test.pendingPods, test.quotas,
		)
		fmt.Printf(
			"\n%s\nexpected: %s\nactual:   %+v\n", test.name, test.expected,
			actual,
		)
	}
}
