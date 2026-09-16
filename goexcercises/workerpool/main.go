package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"workerpool/db"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := filepath.Join(dir, "users.db")
	fmt.Println(dbPath)
	// Initialize repository with dependency injection
	repo, err := db.NewBoltRepository(dbPath)
	must(err)
	defer repo.Close()
	for _ = range 100 {
		user := db.User{
			Username:  gofakeit.Username(),
			Email:     gofakeit.Email(),
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
		}
		_ = repo.AddUser(user)
	}
	//usernames, err := repo.AllUserNames()
	//if err != nil {
	//	panic("something went wrong , could not get usernames")
	//}
	//for _, username := range usernames {
	//	fmt.Println(username)
	//}

}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
