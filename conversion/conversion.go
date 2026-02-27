package conversion

import "strconv"

func StringsToFloats(arr []string) ([]float64,error){
	result := make([]float64,len(arr))

	for ind, val := range arr{
		temp , err := strconv.ParseFloat(val,64)
		if err != nil{
			return nil,err
		}
		result[ind] = temp
	}
	return result,nil
}