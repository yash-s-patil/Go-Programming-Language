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
	}

	fmt.Println("Your Choice:", choice)
}
