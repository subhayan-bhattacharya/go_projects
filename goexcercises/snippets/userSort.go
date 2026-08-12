package snippets

import (
	"cmp"
	"slices"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
	Role string `json:"Role"`
}

func SortUsers(users []User) {
	// Reorders the list directly in place
	slices.SortFunc(users, func(a, b User) int {
		roleCmp := cmp.Compare(a.Role, b.Role)
		if roleCmp == 0 {
			ageCmp := cmp.Compare(b.Age, a.Age)
			if ageCmp == 0 {
				return cmp.Compare(a.Name, b.Name)
			}
			return ageCmp
		}
		return roleCmp
	})
}
