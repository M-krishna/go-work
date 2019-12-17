package main

import "fmt"

/*
	Empty Interface.
	An interface which has zero methods is called empty interface.
	It is represented as interface{}.
	"Since empty interface has zero methods, all types implement the empty interface."
*/

func describe(i interface{}){
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main(){
	s := "Hello World"
	describe(s)
	i := 6
	describe(i)
	strt := struct{
		name string
	}{
		"krishna",
	}
	describe(strt)
}