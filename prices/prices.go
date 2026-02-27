package prices

import (
	"fmt"

	"github.com/bits-and-atoms/Price_Calculator/conversion"
	"github.com/bits-and-atoms/Price_Calculator/filemanager"
)

type TaxIncludedPriceJob struct{
	IOManager filemanager.FileManager `json:"-"` //excluded
	TaxRate float64 `json:"tax_rate"`
	InputPrice []float64 `json:"input_prices"`
	TaxIncludedPrice map[string]string `json:"tax_included_prices"`
}

func (t * TaxIncludedPriceJob) Process(){
	result := make(map[string]string)
	t.LoadPrices()
	for _ , price := range t.InputPrice{
		temp := fmt.Sprintf("%.3f",price + price * t.TaxRate)
		result[fmt.Sprintf("%.3f",price)] = temp;
	}
	t.TaxIncludedPrice = result
	// fmt.Println(result)
	err := t.IOManager.WriteJson(t)
	if err != nil{
		fmt.Println("unable to write resutl json")
		return
	}
}

func (t * TaxIncludedPriceJob) LoadPrices(){
	lines,err := t.IOManager.ReadLines()
	if err != nil{
		fmt.Println("error in reading price file")
		return
	}
	prices ,err:= conversion.StringsToFloats(lines)
	if err != nil{
		fmt.Println("error in converting to float in prices",err)
		return
	}
	t.InputPrice = prices
}
func NewTaxIncludedPriceJob(ioHandler *filemanager.FileManager, tax float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		TaxRate: tax,
		IOManager: *ioHandler,
	}
}