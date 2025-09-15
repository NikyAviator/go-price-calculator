package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("an Error occured while opening the file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		// lines will hold all the prices in the prices.txt file
		lines = append(lines, scanner.Text())
	}
	scannerErr := scanner.Err()

	if scannerErr != nil {
		file.Close()
		return nil, errors.New("failed to read the line in file")
	}
	file.Close()
	return lines, nil
}

func (fm FileManager) WriteJSONToFile(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create the file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to convert data to json")
	}
	file.Close()
	return nil

}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
