package study

import (
	"fmt"
	"math"
)

func TestInterface() {
	rectangle := Rectangle{}
	var shape Shape = &rectangle
	rectangle.width = 1
	rectangle.length = &rectangle.width
	fmt.Printf("rectangle.length: %v\n", rectangle.length)
	fmt.Printf("rectangle.area(): %.2f\n", rectangle.area())
	fmt.Printf("rectangle.length: %v\n", rectangle.length)
	fmt.Printf("shape.area(): %.2f\n", shape.area())
	fmt.Printf("width = %.2f lenth = %.2f\n", rectangle.width, *rectangle.length)
	shape = Circle{1}
	fmt.Printf("%.2f\n", shape.area())
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	length *float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r *Rectangle) area() float64 {
	i := 2.0
	r.length = &i
	return r.width * (*r.length)
}
