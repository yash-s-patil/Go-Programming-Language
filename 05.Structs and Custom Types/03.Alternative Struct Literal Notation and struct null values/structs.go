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

	// we can omit the key and just specify the value (but the order should be same as the blueprint)
	appUser = user{
		userfirstName,
		userlastName,
		userbirthdate,
		time.Now(),
	}

	// all field will have null value
	appUser = user{}

	// ... do something awesome with that gathered data!

	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
