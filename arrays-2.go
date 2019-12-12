package main

import "fmt"

// same as arrays.go, but slight different and formatting

func changeLocal(num [5]int){
	num[0] = 55
	fmt.Println("Local Change Function ", num)
}

// printing multidimensional array
func multidimensionalArray(a [3][2]string){
	for _, v1 := range a{
		for _, v2 := range v1{
			fmt.Printf("%s ", v2)
		}
		fmt.Println()
	}
}

func main(){

	var a [3]int // int array with length 3
	/* 	
		the output will be [0 0 0]
		we have only declared an array
		all elements in the array are automatically assigned 
		to 0 value
	*/
	fmt.Println(a)

	// lets assign some values to the above array
	a[0] = 1
	a[1] = 2
	a[2] = 3
	
	fmt.Println(a) // [1 2 3]

	// lets create the same array using short hand syntax
	// using different variable
	var b = [3]int{1, 2, 3}
	fmt.Println("Short hand syntax")
	fmt.Println(b) // [1 2 3]
	/* 	its not necessary to have all elements in the array to be 
		assigned a value during short hand syntax
		see the below example
	*/
	var c = [3]int{1}
	fmt.Println("Not assigning all the values")
	fmt.Println(c) // [1 0 0]

	/* 	we can also use a literal(...) to let the compiler specify 
		the length for us
	*/
	var d = [...]int{1}
	fmt.Println("Using array literal")
	fmt.Println(d) // [1]

	/*
		the size of the array is a part of the type.
		hence [5]int and [25]int are distinct types.
		because of this arrays cannot be resized.
	*/

	// e := [5]int{1, 2, 3}
	// var f [3]int
	// f = e // not possible since e & f are distinct types

	/*
		Arrays are "value types" and not "reference types"
		if they are assigned to a new variable
		a copy of the original array is assigned to a new variable
		if changes are made to a new variable, it will not affect the
		original array
	*/

	g := [...]string{"India", "USA", "Germany", "Finland", "France"}
	h := g
	h[0] = "Tamilnadu"
	fmt.Println("Original Array", g) // Original Array [India USA Germany Finland France]
	fmt.Println("New Array", h) // New Array [Tamilnadu USA Germany Finland France]
	fmt.Println("Original Array", g) // Original Array [India USA Germany Finland France]

	/*
		similarly when passing an array to a function, 
		they are passed by value and the original array is
		unchanged
	*/
	var aa = [5]int{1,2,3,4,5}
	changeLocal(aa)
	fmt.Println("Original array after function call " , aa) 

	/* 
		Length of an array can be found using len() function 
	*/

	fmt.Println("Length of the above array ", len(aa))

	// iterating arrays using range
	// normal iteration
	var bb = [5]int{1,2,3,4,5}
	for i := 0; i < len(bb); i++ {
		fmt.Printf("%d th element of a is %d\n", i, bb[i])
	}

	/*
		Now using range
		range returns both the index and value.
	*/
	fmt.Println("Using Range")
	for i, v := range bb{
		fmt.Printf("%d th element of a is %d\n", i, v)
	}

	// multidimensional arrays
	multiarray1 := [3][2]string{
		{"lion", "tiger"},
		{"dog", "cat"},
		{"cow", "goat"},
	}
	multidimensionalArray(multiarray1)

	var multiarray2 [3][2]string
	multiarray2[0][0] = "John"
	multiarray2[0][1] = "Doe"
	multiarray2[1][0] = "Foe"
	multiarray2[1][1] = "Bar"
	multiarray2[2][0] = "William"
	multiarray2[2][1] = "Son"

	multidimensionalArray(multiarray2)

}