package seeder

import (
	"workerpool/db"

	"github.com/brianvoe/gofakeit/v7"
)

func SeedInput(repo *db.BoltRepository) []string {
	var data []string
	seededData, _ := repo.AllUserNames(5)
	data = append(data, seededData...)
	newFaker := gofakeit.New(11)
	for _ = range 5 {
		data = append(data, newFaker.Username())
	}
	return data
}

func SeedDb(repo *db.BoltRepository) {
	faker := gofakeit.New(10)
	for _ = range 100 {
		user := db.User{
			Username:  faker.Username(),
			Email:     faker.Email(),
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
		}
		_ = repo.AddUser(user)
	}
}
