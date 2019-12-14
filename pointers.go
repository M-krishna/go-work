package main

import "fmt"

/*
	Pointer.
	A Pointer is a variable which stores address of another variable.
	Eg. b = 156(memory address of b is 0x12091adf) -> a(0x12091adf)
	Now a is said to point b.
*/

// Passing pointer to a function
func change(val *int){
	*val = 55 // changing the value using dereference
}

// Returning a pointer from a function
func returnPointer() *int {
	i := 5
	return &i
}

// passing pointer to an array as a function argument
func modify(arr *[3]int){
	arr[0] = 2000 // short hand syntax for (*arr)[0]
}

// instead of passing array to a pointer.
// we can use slice.
func modifyUsingSlice(sls []int){
	sls[0] = 500
}



func main(){

	// Declaring Pointers.
	// *T is the type of a pointer variable which points to a value of type T.
	fmt.Println("Declaring Pointers")
	b := 255 // initializing value for b
	var a *int = &b // creating a pointer variable and assigning the address of b(the & operator is used to get the address of the variable)
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("Memory address of b is", a)

	// Zero value of a pointer is nil.
	fmt.Println("Zero value of a pointer.")
	c := 25
	var d *int
	if d == nil{
		fmt.Println("d is ", d)
		d = &c
		fmt.Println("d after initialization ", d)		
	}

	/* 
		Creating pointers using new function 
		The new function takes a type as an argument and returns a pointer to a newly allocated zero
		value of the type passed as an argument.
	*/

	fmt.Println("Creating Pointers using new function")
	size := new(int) // size is a pointer variable
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 23
	fmt.Printf("New value of size is %d\n", *size)

	/*
		Dereferencing a pointer.
		This means accessing the value of a variable which the pointer points to.
		(Eg.) *a is the syntax to dereference a.
	*/
	fmt.Println("Dereferencing a pointer")
	bb := 255
	aa := &bb
	fmt.Println("address of b is ", aa)
	fmt.Println("value of b is ", *aa)

	// lets write another program to change the value by using pointer.
	fmt.Println("Changing value by using pointer.")
	cc := 300
	dd := &cc
	fmt.Println("Value of cc is ", *dd)
	*dd++
	fmt.Println("Changed value of cc is ", *dd)

	// Passing pointer to a function
	fmt.Println("Passing pointer to a function.")
	e := 10
	fmt.Println("Before function call ", e)
	f := &e
	change(f)
	fmt.Println("After function call ", e)

	// Returning pointer from a function
	fmt.Println("Returning a Pointer from a function")
	g := returnPointer()
	fmt.Println("The value which is returned from the function is ", *g)

	/*
		Note:
		Do not pass a pointer to an array as a argument to a function. 
		Use slice instead.
	*/
	fmt.Println("Passing pointer to an array as a function argument")
	h := [3]int{1,2,3}
	fmt.Println("Before passing pointer array to the function ", h)
	modify(&h)
	fmt.Println("After passing pointer array to the function ", h)

	// But this is not idiomatic way of achieving this in Go.
	// we can use slices for this.
	fmt.Println("Using slice to achieve the above thing.")
	i := []int{1,2,3,4,5}
	fmt.Println("Before passing the slice to the function", i)
	modifyUsingSlice(i[:])
	fmt.Println("After passing the slice to the function", i)

	/*
		Go does not support pointer arithmetic
	*/
}