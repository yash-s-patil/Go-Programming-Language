package user

import (
	"errors"
	"fmt"
	"time"
)

// custom struct - blueprint
type User struct {
	// we dont want to expose fields outside the user package
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// build a new struct based on existing struct - struct embedding
type Admin struct {
	email    string
	password string
	User
}

// method of user struct
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

// mutation method - change the data in struct
// we want to change the original struct instance and not the copy of it. so we use pointer to it
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// utility function - helps creating struct (constructor)
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}
