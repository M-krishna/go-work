package main

import (
	"fmt"
	"math"
)

type Shape interface{ // creating a shape interface
	Area() float64
	Perimeter() float64
}

type Rect struct{
	width float64
	height float64
}

type Circle struct{
	radius float64
}

func (r Rect) Area() float64{
	return r.width * r.height
}

func (r Rect) Perimeter() float64{
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.radius
}

func main(){
	var s Shape
	fmt.Println("values of s is", s)
	fmt.Printf("type of s is %T\n", s)
	s = Rect{4.0, 3.0}
	r := Rect{4.0, 3.0}
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area of rectangle s", s.Area())
	fmt.Println("s == r is", s == r)
}
