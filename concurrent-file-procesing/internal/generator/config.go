package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

type CommandConfig struct {
	Rows   int
	Files  int
	OutDir string
}

func (c CommandConfig) Run() error {
	err := os.MkdirAll(c.OutDir, 0755)
	if err != nil {
		return err
	}
	generator := NewGenerator(c.Rows)
	for i := range c.Files {
		filePath := filepath.Join(c.OutDir, fmt.Sprintf("transactions_%d.csv", i+1))
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("could not create file: %w", err)
		}
		defer file.Close()
		data := generator.GenerateTransaction()
		var stringTransactionData []string
		for _, transaction := range data {
			stringData, err := transaction.ToCsv()
			if err != nil {
				return err
			}
			_, err = file.WriteString(stringData)
			if err != nil {
				return err
			}
			stringTransactionData = append(stringTransactionData, stringData)
		}

	}
	return nil
}
