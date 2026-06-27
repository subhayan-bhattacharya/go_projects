package main

import (
	"composite"
	"fmt"
)

func main() {
	policy := composite.And(
		composite.Or(
			composite.HasPlan("enterprise"),
			composite.PermissionRule("analytics:write"),
		),
		composite.PermissionRule("app:login"),
	)
	users := []struct {
		Name string
		User composite.User
	}{
		{
			Name: "Free user with login only",
			User: composite.User{
				Plan:        "free",
				Permissions: []string{"app:login"},
			},
		},
		{
			Name: "Enterprise user with login",
			User: composite.User{
				Plan:        "enterprise",
				Permissions: []string{"app:login"},
			},
		},
		{
			Name: "Free user with analytics write and login",
			User: composite.User{
				Plan:        "free",
				Permissions: []string{"analytics:write", "app:login"},
			},
		},
		{
			Name: "Enterprise user without login",
			User: composite.User{
				Plan:        "enterprise",
				Permissions: []string{},
			},
		},
	}
	for _, item := range users {
		ctx := composite.Context{User: item.User}
		allowed := policy.Evaluate(ctx)
		fmt.Printf("%s => allowed: %v\n", item.Name, allowed)
	}

}
