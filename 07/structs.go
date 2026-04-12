package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

// method attached to a struct
// using User not (*User) so function will recieve copy of struct
func (u User) outputUserDetails() { // (u User) is called receiver argument,
	fmt.Print(u)
	fmt.Println(u.firstName, u.lastName, u.birthday)
}

// method attached to a struct
// using *User so function will recieve reference of the struct
func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthday string) (*User, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("all fields are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}, nil
}

func main() {
	firstName := getUserData("Please enter your first name")
	lastName := getUserData("Please enter your last name")
	birthday := getUserData("Please enter your birthday")

	var appUser *User
	// appUser = User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	// birthday:  birthday, // <-- ommited keys will have nil value set
	// 	createdAt: time.Now(),
	// }

	// short hand notation
	appUser = &User{
		firstName,
		lastName,
		birthday,
		time.Now(),
	}

	appUser.outputUserDetails() // accessing struct method
	appUser.clearUserName()     // accessing struct method
	appUser.outputUserDetails() // accessing struct method

	// using constructor function to create struct instance
	appUser, err := newUser(firstName, lastName, birthday)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
