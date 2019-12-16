package main

import (
	"fmt"
	"math"
)

/*
	The method is just a function with special receiver type between the func keyword and the 
	method name. The receiver can either be a struct type or non-struct type.

	> syntax of method declaration.
	func (t Type) methodname(parameter list){

	}

	The above snippet creates a method named methodname with receiver type Type.
	t is called as the receiver and can be accessed inside the function.
*/

// simple Employee struct
type Employee struct{
	name string
	salary int
	currency string
}

/*
	displaySalary() method has Employee as a receiver type.
*/
func (e Employee) displaySalary(){
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

// Methods vs Functions
func dispSalary(e Employee){
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}


type Rectangle struct {
	length int
	width int
}

type Circle struct {
	radius float64
}

/*
	Methods with same names for different types
*/

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main()  {
	emp := Employee{"krishna", 20000, "$"}
	emp.displaySalary()

	/* 
		Methods vs Functions
	*/
	dispSalary(emp)


	/*
		So why do we have methods when we can write the same program using functions.

		Go is not a pure object-oriented programming language and it does not support classes.
		Hence methods on types are a way to achieve behavior similar to classes.
		Methods allows a logical grouping of behavior related to a type similar to classes.
		Methods with same name can be defined on different types whereas functions with same
		name are not allowed.

	*/
	fmt.Println("Methods for different types")
	r := Rectangle{10, 10}
	fmt.Printf("Area of the rectangle is %d\n", r.Area())

	c := Circle{10.5}
	fmt.Printf("Area of the circle is %.2f\n", c.Area())
}