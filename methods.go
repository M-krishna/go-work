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

// simple address struct
type Address struct{
	city, state string
}

// method for Address struct
func (a Address) getFullAddress(){
	fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}

// simple Employee struct
type Employee struct{
	name string
	salary int
	currency string
	Address // anonymous struct field(itself is a structure)
}

/*
	displaySalary() method has Employee as a receiver type.
*/
func (e Employee) displaySalary(){
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

// Methods vs Functions
func displaySalary(e Employee){
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

/*
	Difference Between value receiver and pointer receiver
*/
func (e Employee) changeName(newName string){
	e.name = newName
}

func (e *Employee) changeSalary(newSalary int){
	e.salary = newSalary
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

/*
	Pointer receivers in methods vs Pointer arguments in functions.
*/
func (r *Rectangle) Perimeter() {
	fmt.Println("Perimeter method output: ", 2 * (r.length + r.width))
}

func Perimeter(r *Rectangle){
	fmt.Println("Perimeter function output: ", 2 * (r.length + r.width))
}


// Methods with non-struct receivers.
type MyInt int
func (a MyInt) add(b MyInt) MyInt{
	return a + b
}


func main()  {
	emp := Employee{name: "krishna", salary: 20000, currency: "$"}
	emp.displaySalary()

	/* 
		Methods vs Functions
	*/
	displaySalary(emp)


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

	/*
		Pointer Receivers vs Value Receivers.
		It is possible to create methods with pointer receivers.
		The difference between value receiver and pointer receiver is, changes made inside a method 
		with pointer receiver is visible to the caller whereas this is not the case in value receiver.

		When to use pointer receiver and when to use value receiver.
		
		Generally pointer receivers can be used when the changes made to the receiver 
		inside the method should be visible to the caller.

		Pointer receivers can also be used in places where its expensive to copy a data structure.
		For Eg. if a struct has so many fields and passing it as a value receiver to a method will
		need the entire struct to be copied which will be expensive. In case if a pointer receiver is 
		used, the struct will not be copied and only a pointer to it will be used in the method.
	*/
	fmt.Println("Pointer Receivers vs Value Receivers")
	emp1 := Employee{name: "Krishna", salary: 20000, currency: "$"}
	fmt.Printf("Employee name before change %s\n", emp1.name)
	emp1.changeName("Micheal")
	fmt.Printf("Employee name after change %s\n", emp1.name)

	fmt.Printf("Employee salary before change %d\n", emp1.salary)
	(&emp1).changeSalary(40000) // can be written as emp1.changeSalary(40000)
	fmt.Printf("Employee salary after change %d\n", emp1.salary)

	/*
		Methods of anonymous struct fields.
		Methods belonging to anonymous fields of a struct can be called as if they belong to a 
		structure where the anonymous field is defined.
	*/
	fmt.Println("Methods of anonymous struct fields.")
	emp2 := Employee{
		name: "krishna",
		salary: 34000,
		currency: "$",
		Address: Address{"Chennai", "Tamilnadu"},
	}
	emp2.getFullAddress()


	/* 
		Value receivers in methods vs Value arguments in functions 

		when a function has only value arguments, it accepts only a value argument.

		when a method has a value receiver, it will accept both pointer and a value receiver.
	*/
	fmt.Println("Value receiver in methods vs Value arguments in functions.")
	emp3 := Employee{
		name: "Krishna",
		salary: 12000,
		currency: "$",
	}
	emp3.displaySalary()
	displaySalary(emp3)

	p := &emp3
	p.displaySalary() // The reason is that the line p.displaySalary(), for convenience will be interpreted by Go as (*p).displaySalary() since displaySalary has a value receiver.
	// displaySalary(p) // throws an error because displaySalary is a function which takes values as an argument.

	/*
		Pointer receivers in methods vs Pointer arguments in functions.

		Functions with pointer arguments will only accept pointers.

		Methods with pointer receivers will accept both pointer and value receivers.
	*/
	fmt.Println("Pointer receivers in methods vs Pointer arguments in functions.")
	r1 := Rectangle{10, 10}
	r1.Perimeter() // calling by value receiver
	// Perimeter(r1) // throws an error because this accepts a pointer.
	// so we can do this.
	p1 := &r1
	p1.Perimeter() // calling by pointer receiver.
	Perimeter(p1)

	/*
		Methods with non-struct receivers.

		so far we have defined methods only on struct types. It is also possible to define methods
		on non-struct types but there is a catch.

		To define a method on a type, the definition of the receiver type and the definition of the 
		method should present in the same package.
	*/
	num1 := MyInt(1)
	num2 := MyInt(2)
	fmt.Println("The sum is ",num1.add(num2))
	fmt.Println("And vice versa num2.add(num1)", num2.add(num1))
}