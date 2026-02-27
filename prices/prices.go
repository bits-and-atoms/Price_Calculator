package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	fl,err := os.Open("prices.txt")
	if err != nil{
		fmt.Println("error to get price file")
		return
	}

	scanner := bufio.NewScanner(fl)
	var lines []string
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	fl.Close()
	if err != nil{
		fmt.Println("could not read price file")
		return
	}
	prices := make([]float64,len(lines))
	for ind,val := range lines{
		fltval ,err := strconv.ParseFloat(val,64)
		if err != nil{
			fmt.Println("cant able to convert price to float")
			return
		}
		prices[ind] = fltval
	}
	t.InputPrice = prices
}
func NewTaxIncludedPriceJob(tax float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		TaxRate: tax,

	}
}