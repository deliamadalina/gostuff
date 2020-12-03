package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{}
	s := square{}

	s.sideLength = 2.5
	t.height = 1.0
	t.base = 2

	printArea(t)
	printArea(s)

}

func (t triangle) getArea() float64 {
	return 0.35 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(sp shape) {
	fmt.Println(sp.getArea())
}
