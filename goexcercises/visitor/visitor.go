package visitor

type Expr interface {
	Accept(v Visitor) error
}

type NumberLiteral struct {
	Value float64
}

func (n *NumberLiteral) Accept(v Visitor) error {
	return v.VisitNumberLiteral(n)
}

type StringLiteral struct {
	Value string
}

func (s *StringLiteral) Accept(v Visitor) error {
	return v.VisitStringLiteral(s)
}

type Variable struct {
	Name string
}

func (ve *Variable) Accept(v Visitor) error {
	return v.VisitVariable(ve)
}

type BinaryExpression struct {
	Left     Expr
	Operator string
	Right    Expr
}

func (b *BinaryExpression) Accept(v Visitor) error {
	return v.VisitBinaryExpression(b)
}

type FunctionCall struct {
	Name string
	Args []Expr
}

func (f *FunctionCall) Accept(v Visitor) error {
	return v.VistFunctionCall(f)
}

type Visitor interface {
	VisitNumberLiteral(n *NumberLiteral) error
	VisitStringLiteral(s *StringLiteral) error
	VisitVariable(vr *Variable) error
	VisitBinaryExpression(b *BinaryExpression) error
	VistFunctionCall(f *FunctionCall) error
}

type VariableCollectorVisitor struct {
	Variables map[string]bool
}

func NewVariableCollectorVisitor() *VariableCollectorVisitor {
	return &VariableCollectorVisitor{Variables: make(map[string]bool)}
}

func (v *VariableCollectorVisitor) VisitNumberLiteral(n *NumberLiteral) error {
	return nil
}

func (v *VariableCollectorVisitor) VisitStringLiteral(s *StringLiteral) error {
	return nil
}

func (v *VariableCollectorVisitor) VisitVariable(vr *Variable) error {
	v.Variables[vr.Name] = true
	return nil
}

func (v *VariableCollectorVisitor) VisitBinaryExpression(b *BinaryExpression) error {
	if err := b.Left.Accept(v); err != nil {
		return err
	}
	if err := b.Right.Accept(v); err != nil {
		return err
	}
	return nil
}

func (v *VariableCollectorVisitor) VistFunctionCall(f *FunctionCall) error {
	for _, arg := range f.Args {
		if err := arg.Accept(v); err != nil {
			return err
		}
	}
	return nil
}
