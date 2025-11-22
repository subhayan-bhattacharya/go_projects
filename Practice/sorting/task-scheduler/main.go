package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"time"
)

type Task struct {
	ID string `json:"id"`
	Priority int `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	Deadline time.Time `json:"deadline"`
	EstimatedHours int `json:"estimated_hrs"`
	Dependencies []string `json:"dependencies"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

func compareTasksByPriority(task1, task2 Task) int {
	return cmp.Compare(task2.Priority, task1.Priority)
}

func compareTasksByDeadline(task1 , task2 Task) int {
	if task1.Deadline.IsZero() && task2.Deadline.IsZero() {
		return 0
	} 
	if task2.Deadline.IsZero() {
		return -1
	}
	if task1.Deadline.IsZero() {
		return 1
	} 
	return task1.Deadline.Compare(task2.Deadline)
}

func isTaskOverdue(task Task) bool {
	return task.Deadline.Before(time.Now())
}

func compareTasks(task1, task2 Task) int {
	if isTaskOverdue(task1) && isTaskOverdue(task2) {
		return cmp.Or(
			compareTasksByPriority(task1, task2),
			compareTasksByDeadline(task1, task2),
		)
	}
	if isTaskOverdue(task1) {
		return -1
	}
	if isTaskOverdue(task2) {
		return 1
	}
	return cmp.Or(
		compareTasksByPriority(task1, task2),
		compareTasksByDeadline(task1, task2),
	)
}

func getData(inputFile string) (Tasks, error) {
	var tasks Tasks
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return tasks, fmt.Errorf("Could not read file %w", err)
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return tasks, fmt.Errorf("Could not unmarshall contents %w", err)
	}
	return tasks, nil
}

func main() {
	data, err := getData("data.json")
	if err != nil {
		fmt.Printf("We could not read data %v", err)
	}
	slices.SortFunc(data.Tasks, compareTasks)
	for _, task := range data.Tasks {
		fmt.Println(task)
	}
}