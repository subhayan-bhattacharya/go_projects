package strategy

import (
	"fmt"
	"io"
	"os"
)

type PaymentStrategy interface {
	Pay(amount float64) bool
	SetWriter(writer io.Writer)
	SetLogger(logger io.Writer)
}

type WriterAndLogger struct {
	Writer io.Writer
	Logger io.Writer
}

func (w *WriterAndLogger) SetWriter(writer io.Writer) {
	w.Writer = writer
}

func (w *WriterAndLogger) SetLogger(logger io.Writer) {
	w.Logger = logger
}

type CreditCard struct {
	Name             string
	CreditCardNumber string
	WriterAndLogger
}

func (c *CreditCard) Pay(amount float64) bool {
	message := fmt.Sprintf("paying amount %f with the card number %s\n", amount, c.CreditCardNumber)
	c.Writer.Write([]byte(message))
	return true
}

type Paypal struct {
	Email string
	WriterAndLogger
}

func (c *Paypal) Pay(amount float64) bool {
	message := fmt.Sprintf("Paying with paypal with email %s\n", c.Email)
	c.Writer.Write([]byte(message))
	return true
}

type ShoppingCart struct {
	PaymentStrategy PaymentStrategy
	Total           float64
}

func (s *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	s.PaymentStrategy = strategy
	strategy.SetWriter(os.Stdout)
}

func (s *ShoppingCart) MakePayment() bool {
	return s.PaymentStrategy.Pay(s.Total)
}
