package command

import (
	"errors"
	"fmt"
)

type InventoryService struct{}

func (s *InventoryService) ReserveStock(item string, qty int) error {
	fmt.Printf("[Inventory] Reserved %d units of '%s'\n", qty, item)
	return nil
}

func (s *InventoryService) ReleaseStock(item string, qty int) {
	fmt.Printf("[Inventory Rollback] Released %d units of '%s'\n", qty, item)
}

type PaymentGateway struct{}

func (p *PaymentGateway) Charge(amount float64) (string, error) {
	if amount > 500.0 {
		return "", errors.New("[Payment Error] Fraud protection triggered: Amount too high")
	}
	fmt.Printf("[Payment] Successfully charged $%.2f\n", amount)
	return "TXN_998877", nil
}

func (p *PaymentGateway) Refund(transactionID string) {
	fmt.Printf("[Payment Rollback] Refunded transaction %s\n", transactionID)
}
