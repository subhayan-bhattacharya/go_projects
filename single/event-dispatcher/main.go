package main

import (
	dispatcher "event-dispatcher/dispatcher"
)
import handlers "event-dispatcher/handlers"
import events "event-dispatcher/events"

func main() {
	signedUser1 := events.UserSignedUp{
		UserName: "subhayan",
		Email:    "subhayan.here@gmail.com",
	}
	signedUser2 := events.UserSignedUp{
		UserName: "admin",
		Email:    "admin.here@gmail.com",
	}
	dispatcherSignedUp := dispatcher.NewDispatcher[events.UserSignedUp]()
	_ = dispatcherSignedUp.Subscribe(handlers.LoggerHandler)
	_ = dispatcherSignedUp.Subscribe(handlers.DatabaseHandler)
	_ = dispatcherSignedUp.Subscribe(handlers.RiskyHandler)
	_ = dispatcherSignedUp.Publish(signedUser1)
	_ = dispatcherSignedUp.Publish(signedUser2)
}
