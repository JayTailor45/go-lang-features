package main

import (
	"fmt"

	"jaytailor.com/price-calculator/filemanager"
	"jaytailor.com/price-calculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)

		fm := filemanager.New("prices.txt", fmt.Sprintf("output_%.0f.json", taxRate*100))
		// cm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index])
		// if err != nil {
		// 	fmt.Println("Could not process this job")
		// 	fmt.Println(err)
		// }
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}
