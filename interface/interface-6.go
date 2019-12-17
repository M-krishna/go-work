package main

import "fmt"

/*
	Type Assertion.
	It is used to extract the underlying value of the interface.
	i.(T) is the syntax which is used to get the underlying value of the interface i of concrete type T.

*/

func assert(i interface{}){
	// s := i.(int) // get the underlying int value from i
	// fmt.Println(s)

	v, ok := i.(int)
	fmt.Println(v, ok)
}

func main(){
	var s interface{} = 10 // if we convert this to string, it will throw as an error.
	assert(s)

	// To work with other types we can use this syntax (v, ok := i.(T))
	var v interface{} = "Krishna"
	assert(v)
}