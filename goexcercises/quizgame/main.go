package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	csvFilename := flag.String("csv", "inputs.csv", "csv file name to process")
	timeLimit := flag.Int("limit", 5, "time limit for the quiz in seconds")
	flag.Parse()

	problems, err := extractCsv(*csvFilename)
	if err != nil {
		fmt.Printf("Error: could not extract contents from %s: %v\n", *csvFilename, err)
		os.Exit(1)
	}

	score, wrong := playGame(*timeLimit, problems)
	fmt.Printf("\nFinal score: %d\n", score)
	fmt.Printf("Mistakes: %d\n", wrong)
}

func playGame(timeLimit int, problems []Problem) (int, int) {
	score, wrong := 0, 0
	reader := bufio.NewReader(os.Stdin)
	for _, problem := range problems {
		fmt.Printf("What is the answer to %s: ", problem.Question)
		answerChan := make(chan string, 1)
		timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
		go func() {
			<-timer.C
			answerChan <- "timeout"
		}()
		input, _ := reader.ReadString('\n')
		userAnswer := strings.TrimSpace(input)
		timer.Stop()
		select {
		case answer := <-answerChan:
			if answer == "timeout" {
				fmt.Println("time is up!")
				wrong++
			}
		default:
			if userAnswer == problem.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				wrong++
			}
		}
	}
	return score, wrong
}

func extractCsv(fileName string) ([]Problem, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	problems := make([]Problem, 0, len(data))
	for _, line := range data {
		if len(line) >= 2 {
			problems = append(problems, Problem{
				Question: line[0],
				Answer:   line[1],
			})
		}
	}
	return problems, nil
}
