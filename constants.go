package main

import "fmt"

//const a int = 20 // package level constant variable
func main(){
	const a int = 45
	const b = 10
	var aa int16 = 27
	const bb = 10
	fmt.Printf("%v, %T \n\n", a + b, a + b)
	fmt.Printf("%v, %T \n\n", aa + bb, aa + bb) // here while adding a constant with a variable of type int16, the compiler changes the type of constant to int16.

	sample() //calling the sample function
	sample_two() // function call
	byteExample()
	rolesiota()
}

// iota start

const (
	_ = iota
	a
	b
	c
)

func sample(){
	var specialVariable int
	fmt.Printf("%v \n", a == specialVariable)
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
}

const (
	_ = iota + 5
	catSpecial
	dogSpecial
	lionSpecial
)

func sample_two(){
	fmt.Printf("%v \n", catSpecial)
}


const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func byteExample(){
	fileSize := 40000000000.
	fmt.Printf("%v \n", KB)
	fmt.Printf("%v \n", MB)
	fmt.Printf("%v \n", GB)
	fmt.Printf("%v \n", TB)
	fmt.Printf("%.2fGB", fileSize/GB)
}


const (
	_ = iota
	isAdmin = 1 << iota
	isTenant
	isVendor
	isPropertyManager
)

func rolesiota(){
	var a byte = 1 << 1
	fmt.Printf("%b, %v \n", a, a)
	var roles byte = isAdmin | isTenant | isVendor
	fmt.Printf("%b \n", roles)
	fmt.Println( isAdmin & roles )
	fmt.Printf("Is Admin? %v \n", isAdmin & roles == isAdmin)
}


// iota end
