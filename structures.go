package main

import "fmt"

// creating a Saiyan Structure with Name and Power keys
type Saiyan struct{
	Name string
	Power int
}

// creating a Rectangle Type Structure
type Rectangle struct{
	length, width int
}

func (r Rectangle) Area() int{ // Area by Value
	return r.length * r.width
}

// adding a Perimeter method to the Rectangle Type
func (r Rectangle) Perimeter() int{
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Area_by_reference() int { // Area by Reference
	return r.length * r.width
}

func main(){
	// inferring Saiyan structure to goku variable
	goku := Saiyan{"Goku", 9000}
	Super(goku) // here we are calling the super function to check whether it changes the value of power.
	// but it wont change the value
	// because Go passes values to a function as copies.
	// if you really want to change the value, we need to use pointers.
	SuperOne(&goku)
	fmt.Print(goku)
	goku.Hello()
	fmt.Println(goku)
	// Instance for Rectangle Type
	r1 := Rectangle{4, 3}
	fmt.Println("Rectangle is: ", r1)
	fmt.Println("Rectangle Area is", r1.Area())
	fmt.Println("Rectangle Area by Reference: ", (&r1).Area())
	fmt.Println("Rectangle Perimeter is: ", r1.Perimeter())
}

func Super(s Saiyan){
	s.Power += 1000
}

func SuperOne(s *Saiyan){
	s.Power += 1000
}

// functions on structures
func (s *Saiyan) Hello(){
	s.Power += 100
}

