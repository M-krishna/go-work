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

	// using factory function 
	vegeta := NewSaiyan("Vegeta", 5000)
	fmt.Println(vegeta)
	piccolo := NewSaiyanOne("piccolo", 3000)
	fmt.Println(piccolo)

	// fields of a structure
	sprsaiyan()
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

// structures dont have constructors
// instead we can create a factory function
// that returns the instance of desired type
// here im creating a factory function for Saiyan type
func NewSaiyan(name string, power int) Saiyan{
	return Saiyan{
		Name: name,
		Power: power,
	}
}

// we can also return pointers
func NewSaiyanOne(name string, power int) *Saiyan{
	return &Saiyan{
		Name: name,
		Power: power,
	}
}

// Fields of a Structure
// Fields can be of any type - including other structures and types
// like arrays, maps, interface and functions.

type SuperSaiyan struct {
	Name string
	Power int
	Father *SuperSaiyan
}

func sprsaiyan() {
	gohan := &SuperSaiyan{
		Name: "Gohan",
		Power: 8000,
		Father: &SuperSaiyan{
			Name: "Goku",
			Power: 9000,
			Father: nil,
		},
	}
	fmt.Println(gohan)
	fmt.Println(gohan.Name, gohan.Power, gohan.Father)
}
