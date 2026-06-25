package main

import (
	"fmt"
	"interpreter"
)

func main() {
	ctx := interpreter.Context{
		"x": 420,
	}
	x := interpreter.NewVariableExpression("x")
	left := interpreter.NewNumberExpression(10)
	right := interpreter.NewNumberExpression(20)
	third := interpreter.NewNumberExpression(2)
	add := interpreter.NewAddExpression(left, right)
	subtract := interpreter.NewSubtractExpression(add, third)
	fmt.Println(add.Evaluate(ctx))
	fmt.Println(subtract.Evaluate(ctx))
	newAdd := interpreter.NewAddExpression(x, subtract)
	fmt.Println(newAdd.Evaluate(ctx))
}
