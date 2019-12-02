package main

import "fmt"

// creating a Saiyan Structure with Name and Power keys
type Saiyan struct{
	Name string
	Power int
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
}

func Super(s Saiyan){
	s.Power += 1000
}

func SuperOne(s *Saiyan){
	s.Power += 1000
}
