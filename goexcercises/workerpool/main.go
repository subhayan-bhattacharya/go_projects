package main

import (
	"fmt"
	"os"
	"path/filepath"
	"workerpool/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	homeDir, _ := homedir.Dir()
	dbPath := filepath.Join(homeDir, "users.db")

	// Initialize repository with dependency injection
	repo, err := db.NewBoltRepository(dbPath)
	must(err)
	defer repo.Close()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
