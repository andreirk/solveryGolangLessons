package ex2

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
// $ ./program a=23,45 b=1,3.43

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

type Point struct {
	x float64
	y float64
}

func (p1 *Point) Add(p2 *Point) *Point {
	return &Point{x: p1.x + p2.x, y: p1.y + p2.y}
}

func (p1 *Point) CalcDistance(p2 *Point) float64 {
	// AB = √(xb - xa)2 + (yb - ya)2
	return math.Sqrt((p1.y-p1.x)*2 + (p2.y-p2.x)*2)
}

func main() {

	point1xPtr := flag.Float64("p1x", 0, "First point x")
	point1yPtr := flag.Float64("p1y", 0, "First point y")

	point2xPtr := flag.Float64("p2x", 0, "Second point x")
	point2yPtr := flag.Float64("p2y", 0, "Second point y")

	flag.Parse()

	point1x := *point1xPtr
	point1y := *point1yPtr
	point2x := *point2xPtr
	point2y := *point2yPtr

	point1 := &Point{x: point1x, y: point1y}
	point2 := &Point{x: point2x, y: point2y}

	distance := point1.CalcDistance(point2)
	fmt.Printf("Distance is: %s \n", strconv.FormatFloat(distance, 'f', 4, 64))
}
