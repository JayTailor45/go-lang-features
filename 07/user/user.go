package user

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
func (u User) OutputUserDetails() { // (u User) is called receiver argument,
	fmt.Print(u)
	fmt.Println(u.firstName, u.lastName, u.birthday)
}

// method attached to a struct
// using *User so function will recieve reference of the struct
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthday string) (*User, error) {
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
