package main

import "fmt"

/*
	Zero Value of an interface.
	Zero value of an interface is nil.
	A nil interface has both its underlying value and as well as concrete type as nil.
*/

type Describer interface{
	Describe()
}

func main(){
	var d1 Describer
	if d1 == nil{
		fmt.Printf("d1 is nil and has a type of %T and value of %v\n", d1, d1)
	}

	// If we try to call a method on the nil interface, the program will panic since the nil interface
	// is neither has a underlying value nor a concrete type.
	// d1.Describe() // throws an error.
}