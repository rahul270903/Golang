// Interface
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %0.2f\n", s.Area())
}

func main() {
	var circleRadius float64
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&circleRadius)
	circle := Circle{radius: circleRadius}

	var rectWidth, rectHeight float64
	fmt.Print("Enter the width and height of the rectangle: ")
	fmt.Scan(&rectWidth, &rectHeight)
	rectangle := Rectangle{width: rectWidth, height: rectHeight}

	fmt.Println("Circle:")
	PrintArea(circle)

	fmt.Println("Rectangle:")
	PrintArea(rectangle)
}
