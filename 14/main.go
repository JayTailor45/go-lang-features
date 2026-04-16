package main

import "fmt"

type TransformerFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// functions are first class citizens in Go, we can pass them as arguments to other functions
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	// annonymous function example
	quadrupled := transformNumbers(&numbers, func(i int) int { return i * 4 })

	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(quadrupled)

	// closure example
	counter := genereateCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	// recursion example
	fmt.Println(factorial(5)) // 120

	// variadic function example
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8))

	sumUp(1, 2, 3, 4)
	// numers... expand list operator
	sumUp(1, numbers...)
}

func transformNumbers(nums *[]int, transformer TransformerFn) []int {
	updated := []int{}

	for _, val := range *nums {
		updated = append(updated, transformer(val))
	}
	return updated
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

// possible to return a function from another function
func getTransformerFn() TransformerFn {
	return double
}

// closure example - a function that captures the variables from its surrounding scope
func genereateCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// factorial - example of recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// variadic function example - a function that accepts a variable number of arguments
// ...int collect all operator
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func sumUp(startingValue int, nums ...int) {
	fmt.Println("startingValue", startingValue)
	fmt.Println("nums", nums)
}
