package main

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Send() error
	GetStatus() string
	GetRecipient() string
	MarkAsRead()
	GetAttempts() int
}

type NotificationBase struct {
	Status string
	Attempts int
	IsRead bool
}


func (b *NotificationBase) MarkAsRead() {
	b.IsRead = true
}

func (b *NotificationBase) GetAttempts() int {
	return b.Attempts
}

func (b *NotificationBase) GetStatus() string {
	return b.Status
}


type EmailNotification struct {
	To string
	Subject string
	Body string
	NotificationBase
}

func (e *EmailNotification) Send() error {
	if e.Body == "" {
		return errors.New("Cannot send an email with an empty body\n")
	}
	fmt.Printf("Sending email to %s\n", e.To)
	e.Attempts += 1
	e.Status = "sent"
	return nil
}


func (e *EmailNotification) GetRecipient() string {
	return e.To
}


type SmsNotification struct {
	PhoneNumber string
	Message string
	NotificationBase
}


func (s *SmsNotification) Send() error {
	if s.Message == "" {
		s.Status = "failed"
		return errors.New("Cannot send a message with an empty body\n")
	} else if len(s.Message) > 160 {
		s.Status = "failed"
		return errors.New("Message with more than 160 characters is not supported\n")
	}
	fmt.Printf("Sending Sms to %s\n", s.PhoneNumber)
	s.Attempts += 1
	s.Status = "sent"
	return nil
}

func (s *SmsNotification) GetRecipient() string {
	return s.PhoneNumber
}

type PushNotification struct {
	DeviceId string
	Title string
	Body string
	NotificationBase
	IsDeviceActive bool 
}

func (p *PushNotification) Send() error {
	if !p.IsDeviceActive {
		return errors.New("Cannot send to an inactive device\n")
	}
	p.Attempts += 1
	p.Status = "Sent"
	return nil
}

func (p *PushNotification) GetRecipient() string {
	return p.DeviceId
}

func ProcessNotifications(devices []Notifier) {
	for _, device := range devices {
		err := device.Send()
		if err != nil {
			if _, ok := device.(*EmailNotification); ok {
				fmt.Printf("The email could not be sent %v\n", err)
			}
			if _, ok := device.(*PushNotification); ok {
				fmt.Printf("The push notification could not be sent %v\n", err)
			}
			if _, ok := device.(*SmsNotification); ok {
				fmt.Printf("The sms notification could not be sent %v\n", err)
			}
		}
		fmt.Printf("After successfull send the value of attempts is %d and status is %s\n", device.GetAttempts(), device.GetStatus())
	}
}

func main() {
	emailNotification := EmailNotification{
		To: "subhayan.here@gmail.com",
		Subject: "test email",
		Body: "This is just a test email",
		NotificationBase: NotificationBase{
			Status: "NotSent",
			Attempts: 0,
			IsRead: false,
		},
	}
	smsNotification := SmsNotification{
		PhoneNumber: "01234",
		Message: "test sms notification",
		NotificationBase: NotificationBase{
			Status: "NotSent",
			Attempts: 0,
			IsRead: false,
		},
	}
	pushNotification := PushNotification{
		DeviceId: "1234",
		Title: "Test notification",
		Body: "test notification",
		IsDeviceActive: true,
		NotificationBase: NotificationBase{
			Status: "NotSent",
			Attempts: 0,
			IsRead: false,
		},
	}
	notifiers := []Notifier{&pushNotification, &smsNotification, &emailNotification}
	ProcessNotifications(notifiers)
}

