package generator

import (
	"concurrent-file-procesing/internal/domain"
	"fmt"
	rand "math/rand/v2"
	"time"
)

type Generator struct {
	rng  *rand.Rand
	rows int
}

func NewGenerator(rows int) Generator {
	return Generator{
		rand.New(rand.NewPCG(42, 0)),
		rows,
	}
}

func (g Generator) GenerateTransaction() []domain.Transaction {
	result := make([]domain.Transaction, g.rows)
	currencies := []string{"EUR", "GBP", "USD", "INR", "CHF"}
	for _ = range g.rows {
		transaction := domain.Transaction{
			TransactionId: fmt.Sprintf("TXN-%s", GenerateRandomString(6, g.rng)),
			CustomerId:    fmt.Sprintf("CUST-%s", GenerateRandomString(6, g.rng)),
			MerchantId:    fmt.Sprintf("MERC-%s", GenerateRandomString(6, g.rng)),
			Amount:        g.rng.Float64() * 1000,
			Currency:      currencies[g.rng.IntN(len(currencies))],
			TimeStamp:     time.Time{},
		}
		result = append(result, transaction)
	}
	return result
}

func GenerateRandomString(length int, rng *rand.Rand) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range length {
		result[i] = charset[rng.IntN(len(charset))]
	}
	return string(result)
}
