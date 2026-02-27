package main

import (
	"fmt"

	"github.com/bits-and-atoms/Price_Calculator/filemanager"
	"github.com/bits-and-atoms/Price_Calculator/prices"
)
func main(){
	taxRates := []float64{0,0.07}
	for _,tax := range taxRates{
		fm := filemanager.New("prices.txt",fmt.Sprintf("resutl_%.0f.json",tax*100))
		t := prices.NewTaxIncludedPriceJob(fm,tax)
		t.Process()
	}
}
