package main

import (
	"jaytailor.com/price-calculator/cmdmanager"
	"jaytailor.com/price-calculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("output_%.0f.json", taxRate*100))
		cm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cm, taxRate)
		priceJob.Process()
	}
}
