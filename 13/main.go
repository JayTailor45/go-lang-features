package main

import "fmt"

// typealias
type floatMap map[string]float64

// typealias with custom methods
func (m floatMap) display() {
	fmt.Println(m)
}

func main() {
	// make([]T, length, capacity)
	userName := make([]string, 4, 5)

	userName[0] = "John"

	// userName[1] is null

	userName = append(userName, "David")
	userName = append(userName, "Max")

	fmt.Println(userName)

	courseRatings := make(map[string]int, 3)

	courseRatings["React"] = 5
	courseRatings["Java"] = 4
	courseRatings["Node"] = 3
	courseRatings["Angular"] = 3
	courseRatings["Javascript"] = 3

	fmt.Println(courseRatings)

	marks := make(floatMap, 3)
	marks["John"] = 90.5
	marks["David"] = 85.0
	marks["Max"] = 92.0

	marks.display()

	// for loop
	for index, value := range userName {
		fmt.Println("index : ", index)
		fmt.Println("value : ", value)
	}

	fmt.Println()

	for key, value := range courseRatings {
		fmt.Println("key : ", key)
		fmt.Println("value : ", value)
	}
}
