package main

import (
	"fmt"
	"htmlparser2"
	"io"
	"os"
)

func parseHtml() (io.Reader, error) {
	filepath := "data/file4.html"
	file, err := os.Open(filepath)
	return file, err
}

func main() {
	file, err := parseHtml()
	if err != nil {
		fmt.Println("could not read file.")
	}
	links, error := htmlparser2.Parse(file)
	if error != nil {
		fmt.Println("encountered error parsing links.")
	}
	fmt.Printf("%+v\n", links)
}
