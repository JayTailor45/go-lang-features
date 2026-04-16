package list

import "fmt"

func unpackExample() {
	prices := []int{1, 2, 3, 4, 5}

	prices = append(prices, 6)

	prices = prices[1:]

	fmt.Println(prices)

	discountedPrices := []int{10, 20}

	prices = append(prices, discountedPrices...) // <-- unpacking the discountedPrices slice

	//-------------------------------------
}
