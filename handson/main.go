package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	UserIds      []int        `json:"user_ids"`
	CountryCodes []string     `json:"country_codes"`
	Permissions  []Permission `json:"permissions"`
}

type Permission struct {
	Resource string `json:"resource"`
	Level    int    `json:"level"`
}

func getData() (Data, error) {
	content, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file")
		return Data{}, err
	}
	var data Data
	if err := json.Unmarshal(content, &data); err != nil {
		fmt.Println("Error unmarshalling JSON")
		return Data{}, err
	}
	return data, nil
}

func main() {
	data, _ := getData()
	fmt.Println(data)
}
