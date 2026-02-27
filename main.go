package main

import "github.com/bits-and-atoms/Price_Calculator/prices"
func main(){
	taxRates := []float64{0,0.07}
	for _,tax := range taxRates{
		t := prices.NewTaxIncludedPriceJob(tax)
		t.Process()
	}
}
