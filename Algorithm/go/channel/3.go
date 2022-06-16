package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Cycle struct {
}

func (c Cycle) Area() float64 {
	return 1.0
}

func (c Cycle) Perimeter() float64 {
	return 1.0
}

func main() {
	var s Shape
	s = Cycle{}
	fmt.Println("value of s is", s)
	fmt.Printf("type of s is %T\n", s)

	c := math.Pow(2, 3)
	fmt.Println(c)
}
