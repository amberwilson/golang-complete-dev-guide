package main

import "fmt"

type Triangle struct {
	height float64
	base   float64
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

type Square struct {
	sideLength float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type Shape interface {
	getArea() float64
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {
	triangle := Triangle{base: 2, height: 10}
	square := Square{sideLength: 10}

	fmt.Println("Triangle area for a 2 x 10 is = ")
	printArea(triangle)

	fmt.Println("Square area for a 10 x 10 is = ")
	printArea(square)
}
