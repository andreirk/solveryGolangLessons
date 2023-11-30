package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"lesson2/jsonToCsv"
	"log"
	"os"
	"strconv"
	"strings"
)

func Printfln(str string, data ...interface{}) {
	fmt.Println(fmt.Printf(str, data...))
}

var defaultFilePath string = "./lesson2/jsonToCsv/sample.json"

var FileContent jsonToCsv.TagsContainer

func LoadJson(filePath string) (err error) {
	file, err := os.Open(filePath)
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&FileContent)
	}
	return
}

func printAsCsvToConsole(records [][]string) {
	for _, row := range records {
		fmt.Printf("%s\n", strings.Join(row, ","))
	}
}

func printToCsvFile(records [][]string, filePath string) {
	f, csvError := os.Create(filePath)
	defer f.Close()
	if csvError != nil {
		log.Fatalln("failed to open file", csvError)
	}
	defer f.Close()
	w := csv.NewWriter(f)

	defer w.Flush()
	for _, record := range records {
		if err := w.Write([]string(record)); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}

func main() {
	jsonPathPtr := flag.String("jsonFile", "", "Json file path")
	outPutPathFilePtr := flag.String("output", "", "Output")

	flag.Parse()

	jsonPath := *jsonPathPtr
	outPutPathFile := *outPutPathFilePtr

	if jsonPath == "" {
		fmt.Println("Please provide --jsonFile=[your File Path] and --output=[path to csv ouput] (optional)")
	}

	if jsonPath == "" {
		jsonPath = defaultFilePath
	}

	err := LoadJson(jsonPath)
	if err != nil {
		Printfln("Error Loading FileContent: %v", err.Error())
	} else {
		Printfln("Id: %v", FileContent.Id)
		Printfln("Latitude: %v", FileContent.Latitude)
		Printfln("Longitude: %v", FileContent.Longitude)
		Printfln("Tags: %v", FileContent.Tags)

		records := [][]string{}
		tags := []string{}
		vals := []string{}

		for key, val := range FileContent.Tags {
			tags = append(tags, key)
			switch v := val.(type) {
			case int:
				val, _ = val.(int)
				vals = append(vals, val.(string))
				fmt.Printf("Integer: %v \n", v)
			case float64:
				val, _ = val.(float64)
				var s string = strconv.FormatFloat(val.(float64), 'f', -1, 64)
				vals = append(vals, s)
				fmt.Printf("Float64: %v \n", v)
			case string:
				vals = append(vals, val.(string))
				fmt.Printf("String: %v \n", v)
			case bool:
				vals = append(vals, fmt.Sprint(v))
				fmt.Printf("String: %v \n", v)
			default:
				fmt.Printf("Usupported type %v \n", v)
			}
		}

		records = append(records, tags)
		records = append(records, vals)

		if outPutPathFile != "" {
			printToCsvFile(records, outPutPathFile)
		} else {
			printAsCsvToConsole(records)
		}
	}

}
