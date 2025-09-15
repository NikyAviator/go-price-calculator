package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

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
