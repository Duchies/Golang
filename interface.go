package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main started...")

	circleObj := circle{
		radius: 10,
	}

	squareObj := square{
		length: 5,
	}
	// without interface
	fmt.Println("Circle Area", circleObj.getArea())
	fmt.Println("square Area", squareObj.getArea())

	// with interface
	arr := []shape{&circleObj, &squareObj}

	for _, obj := range arr {
		fmt.Println("Shape Area:", obj.getArea())
	}

	fmt.Println("Main finished...")
}

type circle struct {
	radius int
}

type square struct {
	length int
}

func (c *circle) getArea() int {
	return 2 * c.radius
}

func (s *square) getArea() int {
	return s.length * s.length
}

type shape interface {
	getArea() int
}
