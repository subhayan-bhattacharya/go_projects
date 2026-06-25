package interpreter

type Context map[string]int

type Expression interface {
	Evaluate(ctx Context) int
}

type VariableExpression struct {
	name string
}

func NewVariableExpression(name string) VariableExpression {
	return VariableExpression{
		name: name,
	}
}

func (v VariableExpression) Evaluate(ctx Context) int {
	return ctx[v.name]
}

type NumberExpression struct {
	value int
}

func NewNumberExpression(value int) NumberExpression {
	return NumberExpression{
		value: value,
	}
}

func (n NumberExpression) Evaluate(ctx Context) int {
	return n.value
}

type AddExpression struct {
	left  Expression
	right Expression
}

func NewAddExpression(left, right Expression) AddExpression {
	return AddExpression{
		left:  left,
		right: right,
	}
}

func (a AddExpression) Evaluate(ctx Context) int {
	return a.left.Evaluate(ctx) + a.right.Evaluate(ctx)
}

type SubtractExpression struct {
	left  Expression
	right Expression
}

func NewSubtractExpression(left, right Expression) SubtractExpression {
	return SubtractExpression{
		left:  left,
		right: right,
	}
}

func (s SubtractExpression) Evaluate(ctx Context) int {
	return s.left.Evaluate(ctx) - s.right.Evaluate(ctx)
}
