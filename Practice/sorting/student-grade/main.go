package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type StudentGrade struct {
	Name string `json:"name"`
	Grade int `json:"grade"`
	Attendance float64 `json:"attendance"`
	Submitted bool `json:"submitted"`
}

type Students struct {
	Students []StudentGrade `json:"students"`
}

func didStudentSubmit(grade StudentGrade) int {
	if grade.Submitted {
		return 0
	}
	return 1
}

func compareStudent(student1 , student2 StudentGrade) int {
	return cmp.Or(
		cmp.Compare(didStudentSubmit(student1), didStudentSubmit(student2)),
		cmp.Compare(student2.Grade, student1.Grade),
		cmp.Compare(student2.Attendance, student1.Attendance),
		cmp.Compare(student1.Name, student2.Name),
	)
}

func readFileContents(jsonFile string) (Students, error) {
	var students Students
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		return students, fmt.Errorf("Could not read file %w", err)
	}
	err = json.Unmarshal(data, &students)
	if err != nil {
		return students, fmt.Errorf("Could not unmarshall data %w", err)
	}
	return students, nil
}

func main() {
	contents, err := readFileContents("student-grades.json")
	if err != nil {
		fmt.Printf("We fucked up %v", err)
	}
	slices.SortFunc(contents.Students, compareStudent)
	for _, grade := range contents.Students {
		fmt.Println(grade)
	}
}