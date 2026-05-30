package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func readFile(file string) (cyoa.Story, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	return cyoa.JsonStory(jsonFile)
}

func pathFn(r *http.Request) string {
	log.Printf("The path received is %s", r.URL.Path)
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	path = path[len("/story/"):]
	return path
}

func main() {
	port := flag.Int("port", 3000, "the port number in which the application runs")
	input := flag.String("input", "data.json", "The input json to the code.")
	story, err := readFile(*input)
	if err != nil {
		fmt.Println(err)
	}

	handler := cyoa.NewHandler(story, cyoa.WithTemplate(nil), cyoa.WithCustomPath(pathFn))
	mux := http.NewServeMux()
	mux.Handle("/story/", handler)
	fmt.Printf("Starting the server at port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
