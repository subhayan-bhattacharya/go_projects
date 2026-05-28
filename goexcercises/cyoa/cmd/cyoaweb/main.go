package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func readFile(file string) (cyoa.Story, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	return cyoa.JsonStory(jsonFile)
}

func main() {
	port := flag.Int("port", 3000, "the port number in which the application runs")
	input := flag.String("input", "data.json", "The input json to the code.")
	story, err := readFile(*input)
	if err != nil {
		fmt.Println(err)
	}
	handler := cyoa.NewHandler(story)
	fmt.Printf("Starting the server at port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}
