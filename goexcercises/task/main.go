package main

import (
	"fmt"
	"os"
	"path/filepath"
	"task/cmd"
	"task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	homeDir, _ := homedir.Dir()
	dbPath := filepath.Join(homeDir, "tasks.db")

	// Initialize repository with dependency injection
	repo, err := db.NewBoltRepository(dbPath)
	must(err)
	defer repo.Close()

	// Set the repository in the command context
	cmd.SetRepository(repo)

	must(cmd.RootCommand.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
