package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	csvFilename := flag.String("csv", "inputs.csv", "csv file name to process")
	flag.Parse()
	err, problems := extractCsv(*csvFilename)
	if err != nil {
		fmt.Printf("could not extract contents from %s", *csvFilename)
		os.Exit(1)
	}
	score := 0
	wrong := 0
	for _, problem := range problems {
		var userAnswer int
		question := problem.Question
		answer, _ := strconv.Atoi(strings.TrimSpace(problem.Answer))
		fmt.Printf("What is the answer to %s : ", question)
		fmt.Scan("%d", &userAnswer)
		if userAnswer == answer {
			fmt.Println("Correct ! moving on !")
			score += 1
		} else {
			fmt.Printf("Expected %d got %d\n", answer, userAnswer)
			wrong += 1
		}
	}
	fmt.Printf("Final score for you is : %d\n", score)
	fmt.Printf("you made %d mistakes\n", wrong)
	if err != nil {
		panic(err)
	}
}

func extractCsv(fileName string) (error, []Problem) {
	file, err := os.Open(fileName)
	if err != nil {
		return err, nil
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	if err != nil {
		return err, nil
	}
	problems := make([]Problem, 0, len(data))
	for _, line := range data {
		problem := Problem{
			Question: line[0],
			Answer:   line[1],
		}
		problems = append(problems, problem)
	}
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return err, nil
	}
	return err, problems
}
