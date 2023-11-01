package main

import (
	"fmt"
	"os"
	"strconv"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
// $ ./program x 92323232323345 i 2

func setIthBit(number, i int) int {
	bitMask := 1 << i
	number = number | bitMask
	return number
}

func main() {
	argsWithoutProg := os.Args[1:]
	theNum := argsWithoutProg[1]
	num64, _ := strconv.Atoi(theNum)
	numOfIdx, _ := strconv.Atoi(argsWithoutProg[3])

	result := setIthBit(num64, numOfIdx)
	fmt.Println(result) // Output: 14 (Binary: 1110)
	//fmt.Println("Was: ", strconv.FormatInt(int64(num64), 2))
	fmt.Printf("Was: %s Become: %s \n", strconv.FormatInt(int64(num64), 2), strconv.FormatInt(int64(result), 2))
}

// 	argsWithoutProg := os.Args[1:]

// 	theNum := argsWithoutProg[1]
// 	numOfIdx, _ := strconv.Atoi(argsWithoutProg[3])
// 	num64, _ := strconv.Atoi(theNum)
// 	_ = numOfIdx
// 	valOfBit := (num64 & (1 << (numOfIdx - 1))) >> (numOfIdx - 1)
// 	mask := 1 << numOfIdx
// 	temp := uint32(num64) & bits.Reverse32(uint32(mask))

// 	theVal := (1 << numOfIdx) | num64

// 	fmt.Println("Temp", temp)
// 	fmt.Println("Binary num: ", strconv.FormatInt(int64(num64), 2))
// 	fmt.Println("Bit: ", valOfBit)
// 	fmt.Println("Here: ", strconv.FormatInt(int64(theVal), 2))

// 	//fmt.Prinum64tln(argsWithoutProg)

// func reverse(x uint32, size uint32) uint32 {
// 	return bits.Reverse32(x) >> (32 - size)
// }
