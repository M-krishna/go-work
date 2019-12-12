package main

import "fmt"

// same as slices.go, but slightly different and formatted

/*
	A slice is a convenient, flexible and powerfull wrapper on top of an array.
	Slices do not own any data on their own. They are the just references to existing arrays.
*/


// passing a slice to a function
func subtractOne(n []int){
	for i := range n{
		n[i] -= 2
	}
}

// memory optimization
func countriesneeded() []string{
	countries := []string{"India", "USA", "Germany", "Russia", "France", "Italy"}
	needed_countries := countries[:len(countries)-2]
	countries_copy := make([]string, len(needed_countries))
	copy(countries_copy, needed_countries)
	return countries_copy
}


func main() {

	// creating a slice
	a := [5]int{1,2,3,4,5}
	b := a[1:4] // creates a slice from a[1] to a[3]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Type of a is %T, Value of a is %v\n", a, a)
	fmt.Printf("Type of b is %T, Value of b is %v\n", b, b)

	// another way of creating a slice
	var c = []int{6,7,8} // creates an array and returns a slice reference
	fmt.Println(c)

	/* 
		modifying a slice 
		a slice does not own any of its data
		any modifications done in the slice will affect the underlying array
	*/
	fmt.Println("Modifying a slice")
	darr := [...]int{1,2,3,4,5,6,7,8,9,10}
	dslice := darr[2:4]
	fmt.Println("array before", darr)
	for i := range dslice{
		dslice[i]++
	}
	fmt.Println("array after", darr)

	// when a number of slice share the same underlying array
	// modifying each slice will affect the original array
	numa := [3]int{10,11,12}
	nums1 := numa[:] // creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("original array before change 1", numa)
	nums1[1] = 200
	fmt.Println("array after modification to slice nums1", numa)
	nums2[2] = 300
	fmt.Println("array after modification to slice nums2", numa)

	/*
		length and capacity of the slice

		> The length of the slice is the number of elements in the slice
		> The capacity of the slice is the number of elements in the underlying array 
			starting from the index from which the slice is created.
	*/
	fmt.Println("Length and capacity")
	cars_array := [...]string{"BMW", "Range Rover", "Audi", "Ford", "Ferrari"}
	cars_slice := cars_array[1:3]
	fmt.Println(cars_slice)
	fmt.Printf("Length of the slice is %d, Capacity of the slice is %d\n", len(cars_slice), cap(cars_slice))

	// a slice can be re-sliced upto its capacity
	fmt.Println("Re-Slicing")
	fruitsarray := [...]string{"apple", "orange", "grapes", "watermelon", "banana", "guava"}
	fruitslice := fruitsarray[1:3]
	fmt.Printf("Length of fruitslice is %d, Capacity of fruitslice is %d\n", len(fruitslice), cap(fruitslice))
	fruitslice = fruitslice[:cap(fruitslice)] // re-slicing the fruitslice upto its capacity
	fmt.Printf("Length of fruitslice is %d, Capacity of fruitslice is %d\n", len(fruitslice), cap(fruitslice))

	/*
		creating a slice using make function
		make function takes in 3 arguments(3rd arg. is optional)
		make([]T, len, cap) -> type, length and capacity
	*/
	fmt.Println("Creating a slice using make function")
	makeslice := make([]int, 5, 5)
	fmt.Println(makeslice)

	/*
		appending to a slice.
		arrays are restricted to fixed length and their length cannot be increased.
		slices are dynamic and new elements can be appended to slice using "append"
		function.

		if slices are backed by arrays and arrays are of fixed length then how come
		a slice is of dynamic length.

		what happens under the hood is when new elements are appended to the slice
		a new array is created. the elements of the existing array is copied to the 
		new array and a new slice reference of the new array is returned.

		the capacity of the new slice is now twice that of the old.
	*/
	cars := []string{"BMW", "Range Rover", "Audi", "Ford", "Ferrari"}
	fmt.Println("cars ", cars, "Length is ", len(cars), "Cap. is ", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars ", cars, "Length is ", len(cars), "Cap. is ", cap(cars))
	cars  = append(cars, "TATA")
	fmt.Println("cars ", cars, "Length is ", len(cars), "Cap. is ", cap(cars))
	cars = append(cars, "Benz")
	fmt.Println("cars ", cars, "Length is ", len(cars), "Cap. is ", cap(cars))

	// the zero value of a slice is nil
	// we can append values to a nil function
	var zeroslice []int
	if zeroslice == nil {
		zeroslice = append(zeroslice, 12, 24, 34)
		fmt.Println(zeroslice)
	}

	// it is also possible to append one slice to another slice
	// we can use ... operator
	fmt.Println("Appending to one slice to another slice")
	veggies := []string{"carrot", "brinjal", "potato", "onion"}
	fruits := []string{"banana", "orange"}
	food := append(veggies, fruits...)
	fmt.Println("Food", food)

	/*
		passing a slice to a function
		A slice contains the length, capacity and a pointer to the zeroth element of the array.
		When a slice is passed to a function, even though it's passed by value, 
		the pointer variable will refer to the same underlying array. 
		Hence when a slice is passed to a function as parameter, 
		changes made inside the function are visible outside the function too.
	*/
	fmt.Println("passing a slice to a function")
	numbers := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("Before passing to function", numbers)
	subtractOne(numbers)
	fmt.Println("After passing to function", numbers)

	/*
		Multidimensional Slice
	*/

	prog_languages := [][]string{
		{"C", "C++"},
		{"Java", "Javascript"},
		{"Go", "Rust"},
	}

	for _, v1 := range prog_languages {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	/*
		Memory Optimization

		slices hold references to the underlying array.
		as long as the slice is in memory, the array cannot be garbage collected
		one way to solve this problem is to use copy function
		this way we can create a new slice, and the original array can be garbage collected
	*/
	needed_countries := countriesneeded()
	fmt.Println(needed_countries)
}