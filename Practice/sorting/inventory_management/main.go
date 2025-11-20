package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type Item struct {
	SKU string
	Price float64
	InStock bool
}

func compareItem(item1, item2 Item) int {
	if item1.InStock && ! item2.InStock {
		return -1
	} else if ! item1.InStock && item2.InStock {
		return 1
	}
	return cmp.Compare(item2.Price, item1.Price)
}

func readContents(jsonFile string) ([]Item, error) {
	items := []Item{}
	contents, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("Could not read file")
	}
	err = json.Unmarshal(contents, &items)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshall contents")
	}
	return items, nil
}


func main() {
	contents, err := readContents("inventory.json")
	if err != nil {
		fmt.Println("Fucked up!")
	}
	slices.SortFunc(contents, compareItem)
	for _, item := range contents {
		fmt.Println(item)
	}
}
