package main

import (
	"fmt"
	"math"
)

type Describer interface { //interfaces should end in 'er' as they are verbs - they DO something.
	describe()
}
type Circle struct {
	radius float64
}

func (c Circle) describe() {
	area := math.Pi * math.Pow(c.radius, 2)
	circ := 2 * math.Pi * c.radius
	fmt.Printf("A circle with r=%v, a=%.2f, & c=%.2f\n", c.radius, area, circ)
}
func main() {
	circle := Circle{5}
	circle.describe()
}
