package events

import "time"

type UserSignedUp struct {
	UserName string
	Email    string
}

type WelcomeEmailSent struct {
	Recipient string
	TimeStamp time.Time
}
