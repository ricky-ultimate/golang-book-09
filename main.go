package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return c.r + c.x //this is absolute nonsense
}

type Rectangle struct {
	w, l float64
}

func (r *Rectangle) area() float64 {
	return r.l * r.w
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.l + r.w)
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalPerimeter(shape ...Shape) float64 {
	var total float64

	for _, s := range shape {
		total += s.perimeter()
	}
	return total
}

func main() {

	c := Circle{1.0, 2.0, 3.0}
	r := Rectangle{1.0, 2.0}

	fmt.Println(totalPerimeter(&c, &r))
	fmt.Printf(" circleperimeter : %v, rect_per : %v", c.perimeter(), r.perimeter())
}
