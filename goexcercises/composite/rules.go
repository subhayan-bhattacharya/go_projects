package composite

type User struct {
	Plan        string
	Permissions []string
}

type Context struct {
	User User
}

type Rule interface {
	Evaluate(ctx Context) bool
}

type HasPermissionRule struct {
	Permission string
}

func PermissionRule(permission string) Rule {
	return HasPermissionRule{
		Permission: permission,
	}
}

func (h HasPermissionRule) Evaluate(ctx Context) bool {
	for _, permission := range ctx.User.Permissions {
		if permission == h.Permission {
			return true
		}
	}
	return false
}

type OrRule struct {
	Rules []Rule
}

func (o OrRule) Evaluate(ctx Context) bool {
	for _, rule := range o.Rules {
		if rule.Evaluate(ctx) {
			return true
		}
	}
	return false
}

func Or(rules ...Rule) Rule {
	return OrRule{Rules: rules}
}

type HasPlanRule struct {
	Plan string
}

func (h HasPlanRule) Evaluate(ctx Context) bool {
	if ctx.User.Plan == h.Plan {
		return true
	}
	return false
}

func HasPlan(plan string) Rule {
	return HasPlanRule{Plan: plan}
}

type AndRule struct {
	Rules []Rule
}

func (a AndRule) Evaluate(ctx Context) bool {
	for _, rule := range a.Rules {
		if !rule.Evaluate(ctx) {
			return false
		}
	}
	return true
}

func And(rules ...Rule) Rule {
	return AndRule{Rules: rules}
}
