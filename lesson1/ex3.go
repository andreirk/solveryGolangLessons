package main

// Реализовать программу, которая принимает на вход файл с курсами валют относительно доллара в CSV формате
// Пример файла
// USD,1
// EUR,1.005
// RUR,100.0
// Реализовать логику сравнения, складывания, вычитания структуры type Money между собой, конвертации в нужную валюту (type Currency string)

import (
	"encoding/csv"
	"fmt"
	"log"
	_ "math"
	"os"
	"strconv"
	_ "strconv"
)

type Currency struct {
	name      string
	rateToUsd float64
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
	argsWithoutProg := os.Args[0:]
	fileName := argsWithoutProg[1]
	amount, _ := strconv.ParseFloat(argsWithoutProg[2], 64)
	from := argsWithoutProg[3]
	to := argsWithoutProg[4]

	// readCsv2(fileName)

	currencyMap := make(map[string]float64)
	records := readCsvFile(fileName)

	for _, record := range records {
		cur := record[0]
		val, _ := strconv.ParseFloat(record[1], 64)
		// fmt.Printf("%d %s - %s", idx+1, record[0], record[1])
		currencyMap[cur] = val
	}
	// fmt.Println(currencyMap)

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
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func readCsv2(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	reader.Comma = ','

	// for range []int{1, 2} {
	// 	row, err := reader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal("-- here is err: ", err)
	// 	}
	// 	fmt.Println("--Here", row[1])
	// 	// fmt.Println(row)
	// }

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}

	for _, eachrecord := range records {
		fmt.Printf("Currency: %s = %s \n", eachrecord[0], eachrecord[1])
	}
}
