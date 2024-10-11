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

// method of user struct
func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

// mutation method - change the data in struct
// we want to change the original struct instance and not the copy of it. so we use pointer to it
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
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

	// calling a method of user struct
	appUser.outputUserDetails()

	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
