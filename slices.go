package main

import "fmt"


func main(){
	a := []int{1,2,3} // this is a slice, similar to an array. But this one is dynamic.
	fmt.Printf("%v \n", len(a)) // we can use the len function for slices.
	fmt.Printf("%v \n", cap(a)) // this is a capacity function. which returns the capacity of a slice.
	// here its gonna return 3, because the underlying 

	diffSlice() // function call
	makeFunc()
}


// Different ways of slice
func diffSlice(){
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:] // starting from 3rd index
	d := a[:6] // upto 6th but excluding 6th element
	e := a[3:6] // first value is inclusive and second value is exclusive
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)

	// we can also use slices with an array also.
	// a := [...]int{1,2,3,4,5,6,7,8,9,10}
}


// Make function

func makeFunc(){
	// we can use make function to make slices, arrays etc.
	// Make function takes upto 3 arguments, but minimum 2 args.
	a := make([]int, 3) // where the 1st argument is the type of the data you want to make, 
			    // and the 2nd argument is the length of the slice, 3rd argument is the capacity
			    // because here we are creating a slice object
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", len(a))
	fmt.Printf("%v \n", cap(a))
}
