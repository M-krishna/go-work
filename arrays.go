package main

import "fmt"

func main(){
	var students [3]int // one way of declaring array
	var grades = [3]int{97, 98, 100} // integer array with values
	// we dont need to directly specify the size of the array, instead we can do this
	var marks = [...]int{1, 2, 3, 4} // [...] means the literal syntax. just hold whatever data im passing to the array.
	fmt.Printf("%v \n", students)
	fmt.Printf("%v \n", names)
	fmt.Printf("%v \n", grades)
	fmt.Printf("%v \n", marks)

	var students_two [3]string
	students_two[0] = "Krishna"
	students_two[1] = "Khirshanth"
	students_two[2] = "Kathir"
	fmt.Printf("%v \n", students_two)

	// instead of initialising values like above. we can do this
	var students_three = [3]string{"Krishna", "Khirshanth", "kathir"}
	fmt.Printf("%v \n", students_three)
	
	// function calls
	arrayLength()
	identityMatrix()
	arrDiff()
}


// length of an array

func arrayLength(){
	var a = [3]int{1,2,3}
	fmt.Printf("%v \n", len(a))
}

// another representation of an array(like matrix)

func identityMatrix(){
	var imatrix [3][3]int = [3][3]int{ [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1}}
	fmt.Printf("%v \n", imatrix)

	//same as above, but convinient to read
	var imatrix_2 [3][3]int
	imatrix_2[0] = [3]int{1,0,0}
	imatrix_2[1] = [3]int{0,1,0}
	imatrix_2[2] = [3]int{0,0,1}
	fmt.Printf("%v \n", imatrix_2)
}


//difference between array in Go and other Programming languages
func arrDiff(){
	a := [...]int{2,4,6}
	b := a
	b[1] = 10
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)

	// In Go if we assign an newly initialized array to another variable
	// in this case assigning a to b
	// changing the value of the array in b doesn't affect the values of a
	// If we need to change the values in the original array
	// we have to use pointers
}
