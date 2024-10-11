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
	// var appUser user

	// value to variable should be the instance of the user struct type
	// instance of the struct
	var appUser user = user{
		firstName: userfirstName,
		lastName:  userlastName,
		birthDate: userbirthdate,
		createdAt: time.Now(),
	}

	outputUserDetails(&appUser)
}

func outputUserDetails(u *user) {
	// shortcut
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
	// deference it first - but use shortcut provided by GO
	fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate, (*u).createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
