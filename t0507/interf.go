package main

import (
	"math"
	"fmt"
)

type Shape interface {
	area() float64
}

type Circles struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (circle Circles) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rect Rectangle) area() float64 {
	return rect.height * rect.width
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circles{x: 0, y: 0, radius: 5}
	rect := Rectangle{width: 10, height: 5}
	fmt.Println(getArea(circle))
	fmt.Println(getArea(rect))

	i := 12
	fmt.Println(float32(i) + 3.45)

}
