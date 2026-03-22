package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Println("Please enter your investment amount:")
	fmt.Scan(&investmentAmount)

	fmt.Println("Please enter your investment horizon in years:")
	fmt.Scan(&years)

	fmt.Println("Please enter your expected rate of return:")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value of your investments would be ", futureValue)
	fmt.Println("Real (inflation adjusted) future value of your investments would be ", futureRealValue)
}
