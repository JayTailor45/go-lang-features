package main

import (
	"fmt"
	"math"
)

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	const inflationRate = 2.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue
}

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Please enter your investment amount: \n")
	fmt.Scan(&investmentAmount)

	outputText("Please enter your investment horizon in years: \n")
	fmt.Scan(&years)

	outputText("Please enter your expected rate of return: \n")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value of your investments would be %.2f \n", futureValue)
	formattedFRV := fmt.Sprintf("Real (inflation adjusted) future value of your investments would be %.2f", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}
