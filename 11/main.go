package main

import "fmt"

func main() {
	fmt.Println("1. Create an array with 3 hobbies")
	hobbies := [3]string{"Coding", "Learning", "Gaming"}
	fmt.Println(hobbies)
	fmt.Println()

	fmt.Println("2. Print first hobby and then create slice of remaining hobbies")
	first := hobbies[0]
	second := hobbies[1:3]
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println()

	fmt.Println("3. Create a slice of first two hobbies")
	firstTwo := hobbies[:2]
	fmt.Println(firstTwo)
	fmt.Println()

	fmt.Println("4. expand the slice firstTwo to include all hobbies")
	fmt.Println(len(firstTwo))
	fmt.Println(cap(firstTwo))
	allHobbies := firstTwo[:3]
	fmt.Println(allHobbies)
	fmt.Println()

	fmt.Println("5. create dynamic array")
	courseGoals := []string{"Learn Go", "Build Projects"}
	fmt.Println(courseGoals)
	fmt.Println()

	fmt.Println("6. Update second goal and insert third goal")
	courseGoals[1] = "Build Real-World Projects"
	courseGoals = append(courseGoals, "Contribute to Open Source")
	fmt.Println(courseGoals)
	fmt.Println()

	fmt.Println("7. Create dynamic list with product struct with two products and then add a third product")
	type Product struct {
		Name  string
		Price float64
	}

	products := []Product{
		{Name: "Laptop", Price: 999.99},
		{Name: "Smartphone", Price: 499.99},
	}

	products = append(products, Product{Name: "Tablet", Price: 299.99})
	fmt.Println(products)
	fmt.Println()

}
