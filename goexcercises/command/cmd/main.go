package main

import "command"

func main() {
	iCommandHandler := command.InventoryService{}
	iCommand := command.InventoryCommand{
		Handler: iCommandHandler,
		Item:    "computer",
		Qty:     500,
	}
	paymentCommandHandler := command.PaymentGateway{}
	pCommand := command.PaymentCommand{
		Handler:       paymentCommandHandler,
		Amount:        1900,
		TransactionId: "",
	}
	commands := []command.Command{iCommand, &pCommand}
	command.Invoke(commands)
}
