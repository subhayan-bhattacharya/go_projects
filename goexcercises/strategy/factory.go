package strategy

import (
	"fmt"
	"os"
)

const (
	CREDITCARD_STRATEGY = "creditcard"
	PAYPAL_STRATEGY     = "paypal"
)

func Factory(strategy string) (PaymentStrategy, error) {
	switch strategy {
	case CREDITCARD_STRATEGY:
		return &CreditCard{
			WriterAndLogger: WriterAndLogger{
				Logger: os.Stdout,
			},
		}, nil
	case PAYPAL_STRATEGY:
		return &Paypal{
			WriterAndLogger: WriterAndLogger{
				Logger: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("strategy '%s' not found", strategy)
	}
}
