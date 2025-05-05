package main

import (
	"fmt"
)

func main() {
	var accountNumber, beginningBalance, charges, credits, creditLimit, newBalance int

	fmt.Print("Enter account number: ")
	fmt.Scanln(&accountNumber)

	fmt.Print("Enter beginning balance: ")
	fmt.Scanln(&beginningBalance)

	fmt.Print("Enter total charges: ")
	fmt.Scanln(&charges)

	fmt.Print("Enter total credits: ")
	fmt.Scanln(&credits)

	fmt.Print("Enter allowed credit limit: ")
	fmt.Scanln(&creditLimit)

	newBalance = beginningBalance + charges - credits

	fmt.Println("Account Summary:")
	fmt.Printf("Account Number: %d\n", accountNumber)
	fmt.Printf("New Balance: %d\n", newBalance)

	if newBalance > creditLimit {
		fmt.Println("Credit limit exceeded!")
	} else {
		fmt.Println("Within credit limit.")
	}
}
