package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

type Story map[string]Chapter

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

func readFile(file string) Story {
	
}

func main() {
	input := flag.String("input", "data.json", "The input json to the code.")
	jsonFile, err := os.Open(*input)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var story Story
	err = json.Unmarshal(byteValue, &story)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	for key, value := range story {
		fmt.Println(key)
		fmt.Println(value)
	}
	defer jsonFile.Close()
}
