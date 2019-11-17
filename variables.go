package main

import "fmt"

var i int = 34 // Package level variable declaration

func main(){
	fmt.Println(i)
	var i int = 42 //one way of declaring variable
	fmt.Println(i)
	var j int //another way of declaring variable
	j = 10
	fmt.Println(j)
	k := 20 // Dynamic type variable declaration
	fmt.Println(k)
	fmt.Printf("%v %T \n\n", i, i)

	typeCasting() // function call
}

func typeCasting(){
	var i int = 42
	fmt.Printf("%v, %T \n\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T \n\n", j, j)
}
