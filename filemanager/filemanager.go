package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	fl, err := os.Open(fm.InputFilePath)
	if err != nil {
		fl.Close()
		return nil, err
	}

	scanner := bufio.NewScanner(fl)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	fl.Close()
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func (fm FileManager) WriteJson(data any) error {
	//any or interface{}
	fl, err := os.Create(fm.OutputFilePath)
	if err != nil {
		fl.Close()
		return err
	}
	encoder := json.NewEncoder(fl)
	err = encoder.Encode(data)
	fl.Close()
	if err != nil {
		return err
	}
	return nil
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
