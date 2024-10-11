package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// create a variable of own custom struct type
	// var appUser user

	// value to variable should be the instance of the user struct type
	// instance of the struct
	var appUser *user.User
	appUser, err := user.New(userfirstName, userlastName, userbirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// calling a method of user struct
	appUser.OutputUserDetails()

	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
