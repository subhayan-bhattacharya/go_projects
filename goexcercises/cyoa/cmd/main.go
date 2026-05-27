package main

import (
	"cyoa"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func readFile(file string) (cyoa.Story, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return cyoa.Story{}, err
	}
	var story cyoa.Story
	err = json.Unmarshal(byteValue, &story)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return cyoa.Story{}, err
	}
	defer jsonFile.Close()
	return story, nil
}

func main() {
	input := flag.String("input", "data.json", "The input json to the code.")
	story, err := readFile(*input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(story)
}
