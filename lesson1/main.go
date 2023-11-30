package main

import (
	"flag"
	"fmt"
	"lesson1/ex1"
	_ "os"
	"strconv"
)

func main() {
	numPtr := flag.Int("num", 0, "Number")
	idxPtr := flag.Int("i", 0, "Index")
	toBitPtr := flag.Int("toBit", 0, "To bit")

	flag.Parse()
	numOfIdx := *idxPtr
	theNum := *numPtr
	toBit := *toBitPtr
	result := ex1.SetIthBit(theNum, numOfIdx, toBit)

	fmt.Println("num", theNum)
	fmt.Println("idx", numOfIdx)
	fmt.Println("Result num", result)
	fmt.Printf("Was: %s Become: %s \n", strconv.FormatInt(int64(theNum), 2), strconv.FormatInt(int64(result), 2))
}
