package domain

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"time"
)

type Transaction struct {
	TransactionId string
	CustomerId    string
	MerchantId    string
	Amount        float64
	Currency      string
	TimeStamp     time.Time
}

func (t Transaction) ToCsv() (string, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	row := []string{
		t.TransactionId,
		t.CustomerId,
		t.MerchantId,
		fmt.Sprintf("%.2f", t.Amount),
		t.Currency,
		t.TimeStamp.Format(time.RFC3339),
	}
}
