package main

import (
	"fmt"

	"jaytailor.com/price-calculator/filemanager"
	"jaytailor.com/price-calculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("output_%.0f.json", taxRate*100))
		// cm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])
		// if err != nil {
		// 	fmt.Println("Could not process this job")
		// 	fmt.Println(err)
		// }
	}

	// awaiting for each go routines completion
	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }

	// error can occurs only in some cases, so can't use following to listen for errors
	// for _, errorChan := range errorChans {
	// 	<-errorChan
	// }

	for index := range taxRates {
		select { // waits for multiple channel operations and executes whichever is ready first
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			{
				fmt.Println("Done")
			}
		}
	}
}
