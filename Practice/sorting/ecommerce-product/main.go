package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type Product struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	Rating float64 `json:"rating"`
	Stock int `json:"stock"`
}

type ProductData struct {
	Products []Product
}

func compareStockStatus(p Product) int {
	if p.Stock > 0 {
		return 0
	}
	return 1
}

func compareProduct(p1 , p2 Product) int {
	return cmp.Or(
		cmp.Compare(compareStockStatus(p1), compareStockStatus(p2)),
		cmp.Compare(p1.Price, p2.Price),
	)
}


func readJsonContents (jsonFile string) (ProductData, error) {
	var contents ProductData
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		return contents, fmt.Errorf("Could not read file")
	}
	err = json.Unmarshal(data, &contents)
	if err != nil {
		return contents, fmt.Errorf("Could not unmarshall contents")
	}
	return contents, nil
}

func main() {
	data, _ := readJsonContents("products.json")
	slices.SortFunc(data.Products, compareProduct)
	for _, data := range data.Products {
		fmt.Println(data)
	}
}