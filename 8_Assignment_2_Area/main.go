package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	newTriangle := triangle{height: 3.2, base: 2.5}
	newSquare := square{sideLength: 4}
	printArea(newTriangle)
	printArea(newSquare)
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
