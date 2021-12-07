package main

import (
	"fmt"
	"math"
)

type square struct {
	lenght float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {

	return s.lenght * s.lenght

}
func (c circle) area() float64 {

	return math.Pi * c.radius * c.radius

}

type shape interface {
	area() float64
}

func info(s shape) {

	fmt.Println(s.area())

}

func main() {
	sq := square{33}
	cr := circle{18.2}
	info(sq)
	info(cr)

}
