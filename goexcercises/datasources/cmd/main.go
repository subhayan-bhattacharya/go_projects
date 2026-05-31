package main

import (
	"datasources"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	ID        string
	Name      string
	Amount    float64
	TimeStamp time.Time
}

func rowToTransaction(row []string) (Transaction, error) {
	amount, err := strconv.ParseFloat(strings.TrimSpace(row[2]), 64)
	if err != nil {
		fmt.Printf("could not convert amount of transaction of id %s into float", row[0])
		return Transaction{}, err
	}
	timeStamp, err := time.Parse(time.RFC3339, row[3])
	if err != nil {
		fmt.Printf("could not convert time of transaction of id %s", row[0])
		return Transaction{}, err
	}
	return Transaction{
		ID:        row[0],
		Name:      row[1],
		Amount:    amount,
		TimeStamp: timeStamp,
	}, nil
}

func main() {
	csvFile := flag.String("csvFileName", "data.csv", "csv file for ingesting")
	flag.Parse()
	newCsvReader, _ := datasources.NewCsvReader(*csvFile)
	csvAdapter := datasources.NewCsvDataSourceAdapter[Transaction](newCsvReader, rowToTransaction)
	for _, data := range datasources.Ingest(csvAdapter) {
		fmt.Println(data)
	}
}
