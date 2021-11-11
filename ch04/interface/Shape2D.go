package main

import (
	"fmt"
	"math"
)

type Shape2D interface {
	Perimeter() float64
}

type circle struct {
	R float64
}

func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

func main() {
	c := circle{R: 1.5}
	fmt.Printf("R %.2f -> Perimeter %.3f \n", c.R, c.Perimeter())
	_, ok := interface{}(c).(Shape2D)
	if ok {
		fmt.Println("c is a Shape2D!")
	}
}
