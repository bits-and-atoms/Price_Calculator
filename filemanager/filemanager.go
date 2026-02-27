package filemanager

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string,error){
	fl,err := os.Open(path)
	if err != nil{
		return nil,err
	}

	scanner := bufio.NewScanner(fl)
	var lines []string
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	fl.Close()
	if err != nil{
		return nil,err
	}
	return lines,nil
}