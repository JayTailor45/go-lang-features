package main

import (
	"fmt"

	"jaytailor.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name")
	lastName := getUserData("Please enter your last name")
	birthday := getUserData("Please enter your birthday")

	var appUser *user.User
	// appUser = User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	// birthday:  birthday, // <-- ommited keys will have nil value set
	// 	createdAt: time.Now(),
	// }

	// using constructor function to create struct instance
	appUser, err := user.New(firstName, lastName, birthday)
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails() // accessing struct method
	appUser.ClearUserName()     // accessing struct method
	appUser.OutputUserDetails() // accessing struct method

	admin := user.NewAdmin("test@test.com", "password")
	admin.OutputUserDetails() // accessing struct method from embedded struct
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
