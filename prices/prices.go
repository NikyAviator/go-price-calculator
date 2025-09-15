package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices/prices.txt")

	if err != nil {
		fmt.Println("An Error occured!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		// lines will hold all the prices in the prices.txt file
		lines = append(lines, scanner.Text())
	}
	scannerErr := scanner.Err()

	if scannerErr != nil {
		fmt.Println("An Error occured while scanning the file!")
		fmt.Println(scannerErr)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("An Error occured while converting string to float!")
			fmt.Println(err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
	}
	job.InputPrices = prices
	file.Close()
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}

}
