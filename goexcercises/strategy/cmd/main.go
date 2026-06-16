package main

import (
	"flag"
	"strategy"
)

func main() {
	paymentMethod := flag.String("method", "creditcard", "either creditcard or paypal.")
	cardNumber := flag.String("card", "1234-5678-9012-3456", "Credit card number")
	email := flag.String("email", "user@example.com", "PayPal email")
	amount := flag.Float64("amount", 100.0, "Amount to pay")
	flag.Parse()
	paymentStrategy, _ := strategy.Factory(*paymentMethod)
	switch *paymentMethod {
	case strategy.CREDITCARD_STRATEGY:
		cc := paymentStrategy.(*strategy.CreditCard)
		cc.Name = "Subhayan"
		cc.CreditCardNumber = *cardNumber
	case strategy.PAYPAL_STRATEGY:
		pp := paymentStrategy.(*strategy.Paypal)
		pp.Email = *email
	}
	cart := strategy.ShoppingCart{
		Total: *amount,
	}
	cart.SetPaymentStrategy(paymentStrategy)
	cart.MakePayment()
}
