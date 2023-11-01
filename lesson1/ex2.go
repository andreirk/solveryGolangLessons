package main

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
// $ ./program a=23,45 b=1,3.43

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	x float64
	y float64
}

func (p1 *Point) Add(p2 *Point) *Point {
	return &Point{x: p1.x + p2.x, y: p1.y + p2.y}
}

func (p1 *Point) calcDistance(p2 *Point) float64 {
	// AB = √(xb - xa)2 + (yb - ya)2
	return math.Sqrt((p1.y-p1.x)*2 + (p2.y-p2.x)*2)
}

func main() {

	argsWithoutProg := os.Args[0:]
	point1x, _ := strconv.ParseFloat(argsWithoutProg[1], 64)
	point1y, _ := strconv.ParseFloat(argsWithoutProg[2], 64)

	point2x, _ := strconv.ParseFloat(argsWithoutProg[3], 64)
	point2y, _ := strconv.ParseFloat(argsWithoutProg[4], 64)

	point1 := &Point{x: point1x, y: point1y}
	point2 := &Point{x: point2x, y: point2y}

	distance := point1.calcDistance(point2)

	fmt.Printf("Distance is: %s \n", strconv.FormatFloat(distance, 'f', 4, 64))
}
