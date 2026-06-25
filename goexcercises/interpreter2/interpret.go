package interpreter2

//age >= 18 AND country == "DE"

type Context struct {
	Age     int
	Country string
}

type Rule interface {
	Evaluate(ctx Context) bool
}

type AgeRule struct {
	age int
}

func (a AgeRule) Evaluate(ctx Context) bool {
	return a.age >= ctx.Age
}

type CountryRule struct {
	country string
}

func (c CountryRule) Evaluate(ctx Context) bool {
	return c.country == ctx.Country
}

type AndRule struct {
	left  Rule
	right Rule
}

func (a AndRule) Evaluate(ctx Context) bool {
	return a.left.Evaluate(ctx) && a.right.Evaluate(ctx)
}
