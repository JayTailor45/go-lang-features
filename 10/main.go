package main

import "fmt"

type Product struct {
	id    string
	Name  string
	Price float64
}

func main() {
	products := []Product{
		{id: "1", Name: "Laptop", Price: 999.99},
		{id: "2", Name: "Smartphone", Price: 499.99},
		{id: "3", Name: "Headphones", Price: 199.99},
		{id: "4", Name: "Smartwatch", Price: 299.99},
	}

	for _, product := range products {
		fmt.Println(product.Name, product.Price)
	}

	fmt.Println()

	fmt.Println(products)

	fmt.Println(products[0])

	fmt.Println()

	featuredProducts := products[1:3] // slice of products from index 2 to 3 (exclusive)
	// INFO: slices only provide a view into the underlying array, so modifying the slice will modify the original array
	fmt.Println(featuredProducts)

	fmt.Println()

	fmt.Println(products[0:])
	fmt.Println(products[:2])
	// fmt.Println(products[:26]) // This will cause a runtime error because the index is out of range

	fmt.Println()

	fmt.Println(len(featuredProducts)) // returns the number of elements in the slice
	fmt.Println(cap(featuredProducts)) // returns the capacity of the slice

	fmt.Println()

	// dynamic slice creation -> internally go creates an array and returns a slice that references it
	// INFO: the length of the slice is 0, but the capacity is 5, which means that we can add up to 5 elements to the slice before it needs to be resized
	prices := []float64{45.0}
	updatedPrices := append(prices, 10.5) // returns new array
	fmt.Println(prices)
	fmt.Println(updatedPrices)
	updatedPrices[0] = 50.0
	fmt.Println(prices)
	fmt.Println(updatedPrices)
}
