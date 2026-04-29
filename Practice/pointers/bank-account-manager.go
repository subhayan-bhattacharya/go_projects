package main

import "fmt"

type BankAccount struct {
	AccountNumber string
	Balance float64
	Owner string
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Amount %f has to be more than 0\n", amount)
	}
	b.Balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount float64) (bool, error) {
	if amount <= 0 {
		return false, fmt.Errorf("You cannot withdraw 0 or less than 0 amount \n")
	}
	if b.Balance >= amount {
		b.Balance -= amount
		return true, nil
	}
	return false, fmt.Errorf("The amount requested %f is more than what is available \n", amount)
}

func (b BankAccount) GetBalance() float64 {
	return b.Balance
}


func main() {
	bankAccount := BankAccount{
		AccountNumber: "1234",
		 Balance: 100,
		 Owner: "Subhayan",
	}
	fmt.Printf("Initial balance: $%.2f\n", bankAccount.GetBalance())
	bankAccount.Deposit(20)
	fmt.Printf("After depositing $20: $%.2f\n", bankAccount.GetBalance())
	fmt.Printf("Current balance: $%.2f\n", bankAccount.GetBalance())
	bankAccount.Deposit(30)
	fmt.Printf("After depositing $30: $%.2f\n", bankAccount.GetBalance())
	ok, err := bankAccount.Withdraw(45)
	if ok {
		fmt.Println("Withdraw successfull")
	} else {
		fmt.Println("Withdraw unsuccessfull")
		fmt.Printf("%v", err)
	}
	ok, err = bankAccount.Withdraw(-20)
	if ok {
		fmt.Println("Withdraw successfull")
	} else {
		fmt.Println("Withdraw unsuccessfull")
		fmt.Printf("%v", err)
	}
}
