package main

import "fmt"

func main() {
	age := 32

	var agePointerExplicitType *int = &age // explicit type
	agePointer := &age                     // short hand

	fmt.Println(age)
	fmt.Println("-----------")
	fmt.Println(agePointerExplicitType)
	fmt.Println(agePointer)
	fmt.Println("-----------")
	fmt.Println(agePointer)  // prints address
	fmt.Println(*agePointer) // prints actual value
	fmt.Println(&agePointer) // prints address

	fmt.Println("-----------")
	fmt.Println(getAdultYearsByValue(age))      // call by value
	fmt.Println(getAdultYearsByReference(&age)) // call by reference
}

// pass by value: function gets it's own copy of variable
func getAdultYearsByValue(age int) int {
	return age - 18
}

// pass by address (using pointer): function recieves same reference of a variable
func getAdultYearsByReference(age *int) int {
	// *age = *age - 18 // <--- update to reference variable is also possible
	return *age - 18
}
