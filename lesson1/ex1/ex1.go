package ex1

// import (
// 	"flag"
// 	"fmt"
// 	_ "os"
// 	"strconv"
// )

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
// $ ./program x 92323232323345 i 2

func SetIthBit(number, i, theBit int) int {
	var newNumber, bitMask int
	bitMask = 1 << i
	if theBit == 0 {
		newNumber = number & ^bitMask
	} else {
		newNumber = number | bitMask
	}

	return newNumber
}

// func main() {

// 	numPtr := flag.Int("num", 0, "Number")
// 	idxPtr := flag.Int("i", 0, "Index")
// 	toBitPtr := flag.Int("toBit", 0, "To bit")

// 	flag.Parse()
// 	numOfIdx := *idxPtr
// 	theNum := *numPtr
// 	toBit := *toBitPtr
// 	result := SetIthBit(theNum, numOfIdx, toBit)

// 	fmt.Println("num", theNum)
// 	fmt.Println("idx", numOfIdx)
// 	fmt.Println("Result num", result)
// 	fmt.Printf("Was: %s Become: %s \n", strconv.FormatInt(int64(theNum), 2), strconv.FormatInt(int64(result), 2))
// }
