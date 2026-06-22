package command

import (
	"fmt"
	"slices"
)

type Command interface {
	execute() error
	rollback() error
}

type InventoryCommand struct {
	Handler InventoryService
	Item    string
	Qty     int
}

func (i InventoryCommand) execute() error {
	return i.Handler.ReserveStock(i.Item, i.Qty)
}

func (i InventoryCommand) rollback() error {
	i.Handler.ReleaseStock(i.Item, i.Qty)
	return nil
}

type PaymentCommand struct {
	Handler       PaymentGateway
	Amount        float64
	TransactionId string
}

func (p *PaymentCommand) execute() error {
	id, err := p.Handler.Charge(p.Amount)
	p.TransactionId = id
	return err
}

func (p *PaymentCommand) rollback() error {
	p.Handler.Refund(p.TransactionId)
	return nil
}

func Invoke(commands []Command) {
	var invokedCommands []Command
	for _, command := range commands {
		err := command.execute()
		if err != nil {
			fmt.Printf("command errored out ... performing rollback")
			slices.Reverse(invokedCommands)
			for _, invokedCommand := range invokedCommands {
				_ = invokedCommand.rollback()
			}
			return
		}
		invokedCommands = append(invokedCommands, command)
	}
}
