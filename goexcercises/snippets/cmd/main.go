package main

import (
	"encoding/json"
	"fmt"
	"os"
	"snippets"
)

func main() {
	//input := []string{"tea", "eat", "ate", "listen", "silent", "⌘⌥", "⌥⌘", "⌘⌘⌥", "café", "facé", "hello"}
	//check := snippets.GroupAnagrams(input)
	//for _, group := range check {
	//	fmt.Println(group)
	//}
	var users snippets.Users
	file, err := os.ReadFile("data/users.json")
	if err != nil {
		panic("encountered an error")
	}
	if err := json.Unmarshal(file, &users); err != nil {
		panic(err)
	}
	snippets.SortUsers(users.Users)
	for _, user := range users.Users {
		fmt.Printf("%+v\n", user)
	}

}
