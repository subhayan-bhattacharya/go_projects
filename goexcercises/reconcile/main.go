package main

import (
	"cmp"
	"slices"
)

type PodUpdate struct {
	Namespace string
	Name      string
	Revision  int
	Deleted   bool
}

func ActivePods(updates []PodUpdate) []PodUpdate {
	unique := map[string]PodUpdate{}
	for _, podUpdate := range updates {
		key := podUpdate.Namespace + podUpdate.Name
		_, ok := unique[key]
		if !ok {
			unique[key] = podUpdate
		} else {
			if unique[key].Revision <= podUpdate.Revision {
				unique[key] = podUpdate
			}
		}
	}

	var podsNotDeleted []PodUpdate
	for _, pod := range unique {
		if !pod.Deleted {
			podsNotDeleted = append(podsNotDeleted, pod)
		}
	}
	slices.SortFunc(
		podsNotDeleted, func(a, b PodUpdate) int {
			return cmp.Or(
				cmp.Compare(a.Namespace, b.Namespace),
				cmp.Compare(a.Name, b.Name),
			)
		},
	)
	return podsNotDeleted
}
