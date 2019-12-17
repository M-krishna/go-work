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
	
	var c Shape
	c = Circle{5}
	fmt.Printf("type of c is %T\n", c)
	fmt.Printf("value of c is %v\n", c)
	fmt.Println("Area of Circle c is", c.Area())
	fmt.Println("Perimeter of Circle c is", c.Perimeter())

	// in order to successfully implement an interface
	// you need to implement all the methods declared by the interface
	// in the above code, if we remove perimeter method
	// go will throw us as an error stating that the perimeter method is missing

	//calling the explain function
	type MyString string
	ms := MyString("Hello World")
	explain(ms)
	rr := Rect{3.0, 5.0}
	explain(rr)
}

// empty interface
// when an interface has zero methods, it is called empty interface.
// its represented by interface{}

func explain(i interface{}){
	fmt.Printf("value given to the explain function is of type %T with value %v", i, i)
}
