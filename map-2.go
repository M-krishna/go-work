package main

import "fmt"

/*
	key value pair.

	we can create a map using make function.

	make(map[type of key]type of value)
	make(map[string]int)

	zero value of a map is nil. If you try to add value to nil map panic will occur in runtime.

*/


func main(){

	var personSalary map[string]int // declaring a map(zero value is nil)
	fmt.Printf("Person Salary Declaration %v, Is it nil %v\n", personSalary, personSalary == nil)
	personSalary2 := make(map[string]int) // initializing a map using make function
	fmt.Printf("Person Salary Initialization %v, Is it nil %v\n", personSalary2, personSalary2 == nil)

	// Adding items to a map
	personSalary2["Krishna"] = 12000
	personSalary2["John"] = 20000
	personSalary2["Doe"] = 30000
	fmt.Println("Contents of Person Salary", personSalary2)

	// it is also possible to initialize a map during declaration itself
	personSalary3 := map[string]int {
		"krishna": 12000,
		"John": 30000,
	}
	fmt.Println("Initialization during Declaration", personSalary3)

	// Accessing items of a map
	fmt.Println("Accessing items of a map")
	employee := "krishna"
	fmt.Println("Salary of ", employee, "is ", personSalary3[employee])

	// what happens if the element is not present in the map.
	// it will return the zero value of the type of that element.
	fmt.Println("Element is not present in the map")
	employee1 := "Noname"
	fmt.Println("Salary of ", employee1, "is ", personSalary3[employee1])

	// we did not get any runtime error in the above program
	// so what if we want to know whether an element is present or not
	fmt.Println("Checking if the element is present in map or not.")
	emp := "Noname"
	value, present := personSalary3[emp]
	if present == true{
		fmt.Println("Salary of ", emp, "is ", value)
	} else{
		fmt.Println("Not Present in the map")
	}

	// we can use range to iterate over the map just like for loop
	fmt.Println("Using range to iterate over map")
	for key, value := range personSalary3 {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	// Deleting items from a map
	// delete(map, key). The delete function does not return any value.
	fmt.Println("using delete function")
	fmt.Println("Person Salary before Deleting", personSalary3)
	delete(personSalary3, "krishna")
	fmt.Println("Person Salary after Deleting", personSalary3)

	// Length of the map.
	// The length of the map can be determined using length function.
	fmt.Println("Using Length function on the map")
	fmt.Println("Length = ",len(personSalary3))

	/* 	Map are reference types like slice
		when a map is assigned to a new variable, they both point to a same data structure.
		hence changes made in one will reflect in the other.

		Similar is the case when maps are passed as parameters to functions. 
		When any change is made to the map inside the function, 
		it will be visible to the caller.
	*/
	fmt.Println("Reference Type")
	personSalary3["mike"] = 23400
	fmt.Println("Original person salary", personSalary3)
	newPersonSalary := personSalary3
	newPersonSalary["Kevin"] = 120000
	fmt.Println("Person salary changed", personSalary3)

}