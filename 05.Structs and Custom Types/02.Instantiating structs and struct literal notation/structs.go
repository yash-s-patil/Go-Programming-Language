package main

import (
	"fmt"
	"time"
)

// custom struct - blueprint
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// create a variable of own custom struct type
	var appUser user

	// value to variable should be the instance of the user struct type
	// instance of the struct
	appUser = user{
		firstName: userfirstName,
		lastName:  userlastName,
		birthDate: userbirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
