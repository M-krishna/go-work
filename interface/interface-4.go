package main

import "fmt"

/*
	Interface Internal Representation.

	An interface can be thought of as being represented internally by a tuple (type, value).
	type is the underlying concrete type of the interface and value holds the value of the concrete type.
*/

type Tester interface{
	Test()
}

type MyFloat float64

func (m MyFloat) Test() {
	fmt.Println(m)
}

func describe(t Tester){
	fmt.Printf("T is of type %T, and of value %v\n", t, t)
}


func main(){
	f := MyFloat(23.5)
	var t Tester
	t = f
	describe(t)
	t.Test()
}