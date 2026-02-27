package prices

import (
	"fmt"

	"github.com/bits-and-atoms/Price_Calculator/conversion"
	"github.com/bits-and-atoms/Price_Calculator/filemanager"
)

type TaxIncludedPriceJob struct{
	TaxRate float64
	InputPrice []float64
	TaxIncludedPrice map[string]float64
}

func (t * TaxIncludedPriceJob) Process(){
	result := make(map[string]string)
	t.LoadPrices()
	for _ , price := range t.InputPrice{
		temp := fmt.Sprintf("%.3f",price + price * t.TaxRate)
		result[fmt.Sprintf("%.3f",price)] = temp;
	}
	// t.TaxIncludedPrice = result
	fmt.Println(result)
}

func (t * TaxIncludedPriceJob) LoadPrices(){
	lines,err := filemanager.ReadLines("prices.txt")
	if err != nil{
		fmt.Println("error in reading price file")
		return
	}
	prices ,err:= conversion.StringsToFloats(lines)
	if err != nil{
		fmt.Println("error in converting to float in prices")
		return
	}
	t.InputPrice = prices
}
func NewTaxIncludedPriceJob(tax float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		TaxRate: tax,

	}
}