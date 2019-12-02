package main

import "fmt"

func main(){
	/*
	different types of data types
	> Boolean
	> Integer
	> Float
	> complex numbers
	> Text
	*/

	// boolean
	var a bool // declaring a bool variable
	fmt.Printf("%v, %T", a, a)

	// Integer
	/*
		different types of integers
		int8, int16, int32, int64
	*/
	var b int = 42
	fmt.Printf("%v, %T", b, b)

	aa := 10 // 1010
	bb := 3 // 1100
	fmt.Println(aa + bb)
	fmt.Println(aa - bb)
	fmt.Println(aa * bb)
	fmt.Println(aa / bb) // here it returns an integer value.
	fmt.Println(aa % bb)

	// Bitwise operators
	fmt.Println(aa | bb)
	fmt.Println(aa ^ bb)
	fmt.Println(aa & bb)
	fmt.Println(aa &^ bb)

	// shifting(left & Right)
	aaa := 8
	fmt.Println(aaa << 3)
	fmt.Println(aaa >> 3)

	// Floating points
	var n float64 = 3.14
	fmt.Printf("%v, %T \n\n", n, n)

	// Two types in Text
	// 1st one is string
	s := "I am Krishna"
	fmt.Printf("%v, %T \n\n", s[0], s[0])

	// string concatenation
	s1 := "Hello"
	s2 := "World"
	fmt.Printf("%v, %T \n\n", s1 + s2, s1 + s2)

	// string is an alias for byte
	sr := "Krishna"
	kr := []byte(sr)
	fmt.Printf("%v, %T \n\n", kr, kr)

	// Another datatype called rune
	var r rune = 'a'
	fmt.Printf("%v, %T \n\n", r, r)
}
