package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type Player struct {
	Name string
	Score int
	LastActive int64
}

func comparePlayer(p1 Player, p2 Player) int {
	c := cmp.Compare(p2.Score, p1.Score)
	if c != 0 {
		return c
	}
	return cmp.Compare(p1.LastActive, p2.LastActive)
}


func getContents(jsonFile string) ([]Player, error) {
	players := []Player{}
	fileContent, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %v", err)
	}
	err = json.Unmarshal(fileContent, &players)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshall %v", err)
	}
	return players, nil
}

func main() {
	contents, _ := getContents("scoreboard.json")
	slices.SortFunc(contents, comparePlayer)
  for _, content := range contents {
		fmt.Println(content)
	}
}