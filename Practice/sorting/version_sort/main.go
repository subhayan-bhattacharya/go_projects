package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Software struct {
	Name string
	Version string
}

func parseVersion(version string) []int {
	digits := strings.Split(version, ".")
	result := make([]int, 0, len(digits))
	for _, digit := range digits {
		converted, _ := strconv.Atoi(digit)
		result = append(result, converted)
	}
	if len(result) < 3 {
		result = append(result, 0)
	}
	return result
}

func compareSoftware(software1, software2 Software) int {
	digits1 := parseVersion(software1.Version)
	digits2 := parseVersion(software2.Version)

	if digits1[0] < digits2[0] {
		return -1
	} else if digits1[0] > digits2[0] {
		return 1
	} else {
		if digits1[1] < digits2[1] {
			return -1
		} else if digits1[1] > digits2[1] {
			return 1
		} else {
				return cmp.Compare(digits1[2], digits2[2])
		}
	}
}


func getContents(fileName string)([]Software, error) {
	softwares := []Software{}
	contents, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Could not read file")
	}
	err = json.Unmarshal(contents, &softwares)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshall")
	}
	return softwares, nil
}

func main() {
	contents , err := getContents("softwares.json")
	if err != nil {
		fmt.Println("We fucked up")
	}
	slices.SortFunc(contents, compareSoftware)
	for _, software := range contents {
		fmt.Println(software)
	}
}
