package datasources

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type DataSource[T any] interface {
	ReadAll() ([]T, error)
}

func Ingest[T any](source DataSource[T]) []T {
	output, _ := source.ReadAll()
	return output
}

func readData(filename string) ([][]string, error) {
	var data [][]string
	file, err := os.Open(filename)
	if err != nil {
		return data, fmt.Errorf("unable to open file for reading %w", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return data, fmt.Errorf("could not read csv file %w", err)
		}

		data = append(data, row)
	}
	return data, nil
}

type CsvReader struct {
	rows [][]string
}

func NewCsvReader(fileName string) (*CsvReader, error) {
	fileData, err := readData(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not get data from file %w", err)
	}
	return &CsvReader{
		rows: fileData,
	}, nil
}

func (c *CsvReader) FetchRows() [][]string {
	return c.rows
}

type CsvDataSourceAdapter[T any] struct {
	reader    *CsvReader
	converter func([]string) (T, error)
}

func NewCsvDataSourceAdapter[T any](reader *CsvReader, converter func([]string) (T, error)) CsvDataSourceAdapter[T] {
	return CsvDataSourceAdapter[T]{
		reader:    reader,
		converter: converter,
	}
}

func (c CsvDataSourceAdapter[T]) ReadAll() ([]T, error) {
	var output []T
	for _, row := range c.reader.FetchRows() {
		convertedRow, err := c.converter(row)
		if err != nil {
			fmt.Printf("error converting row %s", err)
			continue
		}
		output = append(output, convertedRow)
	}
	return output, nil
}
