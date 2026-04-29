package main

import "time"

type UserId string
type UserStatus string

const (
	Active   UserStatus = "Active"
	Deleted  UserStatus = "Deleted"
	Disabled UserStatus = "Disabled"
)

func (s UserStatus) isValid(status string) bool {
	switch s {
	case Active, Deleted, Disabled:
		return true
	default:
		return false
	}
}

type User struct {
	UserId    UserId
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    UserStatus
}

func main() {

}
