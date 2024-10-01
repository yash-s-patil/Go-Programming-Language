package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your Choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your Balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your Deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		if depositAmount <= 0 {
			fmt.Println("Invalid Amount")
			return
		}
		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Withdrawl amount: ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)
		if withdrawalAmount <= 0 {
			fmt.Println("Invalid Amount")
			return
		}
		if withdrawalAmount > accountBalance {
			fmt.Println("Invalid Amount. You can't withdraw more than you have")
			return
		}
		accountBalance -= withdrawalAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else {
		fmt.Println("Thankyou Do visit again!")
	}

	fmt.Println("Your Choice:", choice)
}
