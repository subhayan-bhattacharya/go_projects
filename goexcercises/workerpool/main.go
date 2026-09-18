package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"workerpool/db"
	"workerpool/seeder"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := filepath.Join(dir, "users.db")
	// Initialize repository with dependency injection
	repo, err := db.NewBoltRepository(dbPath)
	must(err)
	defer repo.Close()
	seeder.SeedDb(repo)
	seededInputUsernames := seeder.SeedInput(repo)
	dataChanel := make(chan Data)
	//resultsChannel := make(chan Result[bool])
	go sendData(dataChanel, seededInputUsernames)
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
