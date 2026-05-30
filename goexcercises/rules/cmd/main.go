package main

import (
	"fmt"
	"rules"
)

type Order struct {
	Amount    float64
	Country   string
	IsPremium bool
	HasCoupon bool
}

var highValue = rules.Condition[Order]("Order amount is more than 100", func(input Order) bool {
	return input.Amount > 100
})
var isPremium = rules.Condition[Order]("Is the order a premium order?", func(input Order) bool {
	return input.IsPremium
})
var isFromGermany = rules.Condition[Order]("is the order from Germany?", func(input Order) bool {
	return input.Country == "Germany"
})

func main() {
	order := Order{
		Amount:    100,
		Country:   "France",
		IsPremium: true,
		HasCoupon: false,
	}
	fmt.Printf("The description of the rule is %s\n", highValue.Description())
	allRules := []rules.Rule[Order]{highValue, isFromGermany, isPremium}
	andCondition := rules.And[Order](allRules...)
	orCondition := rules.Or[Order](allRules...)
	fmt.Println(andCondition.Evaluate(order))
	fmt.Println(orCondition.Evaluate(order))
}
