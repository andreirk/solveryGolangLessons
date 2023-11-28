package ex3

// Реализовать программу, которая принимает на вход файл с курсами валют относительно доллара в CSV формате
// Пример файла
// USD,1
// EUR,1.005
// RUR,100.0
// Реализовать логику сравнения, складывания, вычитания структуры type Money между собой, конвертации в нужную валюту (type Currency string)

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	_ "math"
	"os"
	"strconv"
	_ "strconv"
)

type Currency string

var currencyMap map[string]float64

type Money struct {
	amount   float64
	currency Currency
}

func (m *Money) add(m2 Money) float64 {
	result1 := m.amount * currencyMap[string(m.currency)]
	result2 := m2.amount * currencyMap[string(m2.currency)]
	return result1 + result2
}

/*
*
Usage example:

	go run ex3.go sample.csv 36 EUR  RUR

*
*/
func main() {
	// 1. parse file
	// 2. make map string[float64]
	fileName := flag.String("file", "", "File name")

	amount := *flag.Float64("amount", 0, "Amount")
	from := *flag.String("from", "", "From")
	to := *flag.String("to", "", "TO")

	currencyMap := make(map[string]float64)
	records := readCsvFile(*fileName)

	for _, record := range records {
		cur := record[0]
		val, _ := strconv.ParseFloat(record[1], 64)
		currencyMap[cur] = val
	}

	result := amount * currencyMap[from] * currencyMap[to]
	fmt.Printf("Result of conversion is: %2.3f \n", result)

}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	// csvReader.LazyQuotes = true
	// csvReader.FieldsPerRecord = -1
	// csvReader.Comma = ','
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
