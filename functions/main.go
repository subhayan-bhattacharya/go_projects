package main

import "fmt"

type SayHappyBirthday interface {
	SayHappyBirthday()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHappyBirthday() {
	fmt.Println("Happy Birthday", p.Name)
	p.Age++
	fmt.Println("You are now", p.Age)
}

type PersonRepository struct {
	people []SayHappyBirthday
}

func (pr *PersonRepository) AddPerson(p Person) {
	pr.people = append(pr.people, p)
}

func (pr *PersonRepository) GetPersonByName(name string) *Person {
	for _, p := range pr.people {
		person, ok := p.(Person) // This is necessary to convert SayHappyBirthday to Person
		if ok && person.Name == name {
			return &person
		}
	}
	return nil
}

func main() {
	repo := &PersonRepository{}

	repo.AddPerson(Person{Name: "Subhayan", Age: 40})
	repo.AddPerson(Person{Name: "Dimpu", Age: 37})
	repo.AddPerson(Person{Name: "Shaayan", Age: 36})

	person := repo.GetPersonByName("Shaayan")
	if person != nil {
		fmt.Println("Person found:", person.Name, "Age:", person.Age)
		person.SayHappyBirthday()
	} else {
		fmt.Println("Person not found")
	}
}
