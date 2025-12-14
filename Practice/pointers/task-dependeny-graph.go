package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Task struct {
    ID           string
    Description  string
    Status       string // "pending", "in-progress", "completed"
    Dependencies [] string // Tasks that must complete before this one
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
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

func DetectCircularDependencies(tasks []Task) ([]string, error) {
    taskMap := make(map[string]Task)
    for _, task := range tasks {
        taskMap[task.ID] = task
    }
    
    state := make(map[string]int)
    var cyclePath []string
    
    // Define a separate helper function
    var exploreDependencies func(taskID string, path []string) bool
    
    exploreDependencies = func(taskID string, path []string) bool {
        // Check if already visiting (cycle!)
        if state[taskID] == 1 {
            cyclePath = path
            return true
        }
        
        // Check if already visited (done)
        if state[taskID] == 2 {
            return false
        }
        
        // Mark as visiting
        state[taskID] = 1
        path = append(path, taskID)
        
        // Explore all dependencies
        task := taskMap[taskID]
        for _, depID := range task.Dependencies {
            if exploreDependencies(depID, path) {  // Recursive call!
                return true
            }
        }
        
        // Mark as visited
        state[taskID] = 2
        return false
    }
    
    // Check each task
    for _, task := range tasks {
        if state[task.ID] == 0 {
            if exploreDependencies(task.ID, []string{}) {
                return cyclePath, errors.New("circular dependency detected")
            }
        }
    }
    
    return nil, nil
}

func ValidateDependencies(tasks []Task) ([]string, error) {
	results := []string {}
	taskMap := map[string]bool {}
	inValidTaskMap := map[string]bool {}
	for _, task := range tasks {
		taskMap[task.ID] = true
	}
	for _, task := range tasks {
		for _, dependentTask := range task.Dependencies {
			_, exists := taskMap[dependentTask]
			_, seen := inValidTaskMap[dependentTask]
			if !exists && !seen {
				results = append(results, dependentTask)
				inValidTaskMap[dependentTask] = true
			} 
		}
	}
	if len(results) > 0 {
		return results, errors.New("Invalid dependencies are present")
	}
	return results, nil
}

func main() {
	tasks, _ := getData("task-dependency-graph-data.json")
	validationResult, _ := ValidateDependencies(tasks.Tasks)
	for _, data := range validationResult {
		fmt.Println(data)
	}
}
