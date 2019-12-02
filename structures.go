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
	fmt.Print(goku)
}
