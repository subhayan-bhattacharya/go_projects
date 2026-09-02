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
	must(db.Init(dbPath))
	must(cmd.RootCommand.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
