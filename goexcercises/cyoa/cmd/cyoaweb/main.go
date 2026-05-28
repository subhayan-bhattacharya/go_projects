package main

import (
	"cyoa"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func readFile(file string) (cyoa.Story, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var story cyoa.Story
	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&story); err != nil {
		fmt.Println(err)
	}
	return story, nil
}

func main() {
	input := flag.String("input", "data.json", "The input json to the code.")
	story, err := readFile(*input)
	if err != nil {
		fmt.Println(err)
	}
	for key, _ := range story {
		fmt.Println(key)
	}
}
