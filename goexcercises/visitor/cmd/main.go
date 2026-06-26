package main

import (
	"fmt"
	"visitor"
)

//max(price * quantity, discount + 100)

func main() {
	price := &visitor.Variable{Name: "price"}
	quantity := &visitor.Variable{Name: "quantity"}
	discount := &visitor.Variable{Name: "discount"}
	hundred := &visitor.NumberLiteral{Value: 100}
	priceTimesQuantity := &visitor.BinaryExpression{
		Left:     price,
		Operator: "*",
		Right:    quantity,
	}
	discountPlusHundred := &visitor.BinaryExpression{
		Left:     discount,
		Operator: "+",
		Right:    hundred,
	}
	expr := &visitor.FunctionCall{
		Name: "max",
		Args: []visitor.Expr{priceTimesQuantity, discountPlusHundred},
	}
	collector := visitor.NewVariableCollectorVisitor()
	err := expr.Accept(collector)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(collector.Variables)
}
