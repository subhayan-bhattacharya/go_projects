package main

import "fmt"

type PaymentMethod interface {
	Process(amount float64) error
}

type CreditCard struct {
	CardNumber string
	TransactionCount int
}

func (c *CreditCard) Process(amount float64) error {
	fmt.Printf("Processing payment of %f amount for creditcard\n", amount)
	c.TransactionCount += 1
	return nil
}

type PayPal struct {
	Email string
	TransactionCount int
}

func (c *PayPal) Process(amount float64) error {
	fmt.Printf("Processing payment of %f amount for paypal\n", amount)
	c.TransactionCount += 1
	return nil
}

type Crypto struct {
	WalletAddress string
}

func (c Crypto) Process(amount float64) error {
	fmt.Printf("Processing payment of %f amount for crypto\n", amount)
	return nil
}

func ProcessMultiplePayments(methods []PaymentMethod, amount float64) {
	for _, method := range methods {
		method.Process(amount)
	}
}

func main() {
	creditCard := CreditCard{
		CardNumber: "123",
		TransactionCount: 34,
	}
	payPal := PayPal{
		Email: "subhayan.here@gmail.com",
		TransactionCount: 23,
	}
	crypto := Crypto{
		WalletAddress: "address",
	}
	methods := []PaymentMethod{&creditCard, &payPal, crypto}
	fmt.Println("Before processing:")
	fmt.Printf("CreditCard transactions: %d\n", creditCard.TransactionCount)
	fmt.Printf("PayPal transactions: %d\n", payPal.TransactionCount)
	
	ProcessMultiplePayments(methods, 123.34)
	
	fmt.Println("\nAfter processing:")
	fmt.Printf("CreditCard transactions: %d\n", creditCard.TransactionCount)
	fmt.Printf("PayPal transactions: %d\n", payPal.TransactionCount)
}
