package main

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
// $ ./program x 92323232323345 i 2

import (
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

func main() {

	argsWithoutProg := os.Args[1:]

	theNum := argsWithoutProg[1]
	numOfIdx, _ := strconv.Atoi(argsWithoutProg[3])
	num64, _ := strconv.Atoi(theNum)
	_ = numOfIdx
	valOfBit := (num64 & (1 << (numOfIdx - 1))) >> (numOfIdx - 1)
	mask := 1 << numOfIdx
	temp := uint32(num64) & bits.Reverse32(uint32(mask))

	theVal := (1 << numOfIdx) | num64

	fmt.Println("Temp", temp)
	fmt.Println("Binary num: ", strconv.FormatInt(int64(num64), 2))
	fmt.Println("Bit: ", valOfBit)
	fmt.Println("Here: ", strconv.FormatInt(int64(theVal), 2))

	//fmt.Prinum64tln(argsWithoutProg)

}

func reverse(x uint32, size uint32) uint32 {
	return bits.Reverse32(x) >> (32 - size)
}
