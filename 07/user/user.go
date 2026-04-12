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

type Admin struct {
	email    string
	password string
	User     // embedding User struct to Admin struct, Admin struct will have access to all fields and methods of User struct
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

func NewAdmin(email, password string) *Admin {
	return &Admin{
		User: User{
			firstName: "admin",
			lastName:  "admin",
			birthday:  "12-12-2000",
			createdAt: time.Now(),
		},
		email:    email,
		password: password,
	}
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
