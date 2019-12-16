package main

import "fmt"

/*
	A structure is a user defined type which represents a collection of fields.
	It can be used in places where it makes sense to group the data into a single unit
	rather than individual types.

	For instance a employee has a firstName, lastName, age. It makes sense to group these 3 properties
	into a single structure employee.
*/


/*
	Declaring a structure.
	the below snippet declares a structure type Employee which has fields firstName, lastName and
	age. 
*/
type Employee struct{
	firstName string
	lastName string
	age int
	salary int
}

// the above structure can also be made more compact by declaring fields that belongs to same type
// in a single line
// the above snippet can be written as

/*
type Employee struct{
	firstName, lastName string
	age, salary int
}
*/

/* 
	the above Employee struct is called "named structure". 
	Because it creates a new type named Employee, which can be used to create structures of type
	Employee.
*/

/*
	It is also possible to declare structures without declaring a new type and these type of
	structures are called "anonymous structures".
*/

// Anonymous structure
var employee struct{
	firstName, lastName string
	age int
}

/* 
	Anonymous fields.
	The snippet below create a struct Person which has 2 anonymous fields string and int. 
*/
type Person struct{
	string
	int
}

func main(){

	// creating a structure using field names
	emp1 := Employee{
		firstName: "Krishna",
		age: 22,
		salary: 20000,
		lastName: "Murugan",
	}

	// creating a structure without using field names
	emp2 := Employee{"Krishna", "Murugan", 22, 20000}

	fmt.Println("Employee 1 with field names", emp1)
	fmt.Println("Employee 2 without field names", emp2)

	// creating anonymous structures
	var emp3 = struct{
		firstName, lastName string
		age, salary int
	}{
		firstName: "John",
		lastName: "Doe",
		age: 22,
		salary: 20000,
	}

	fmt.Println("Anonymous struct", emp3)

	// Zero value of a structure
	var emp4 Employee
	fmt.Println("Zero Value of a structure ", emp4) // {  0 0}
	// firstName & lastName are string, so the zero value is ""
	// age & salary are int, so the zero value is 0

	/*
		It is also possible to specify values for some fields and ignore the rest.
	*/
	emp5 := Employee{
		firstName: "Mark",
		lastName: "Henry",
	}
	fmt.Println("Specifying values to firstName & lastName ", emp5) // {Mark Henry 0 0}

	/*
		Accessing individual fields of a struct.
		The dot (.) operator is used to access the individual fields of a structure.
	*/
	fmt.Println("Accessing individual fields in a structure.")
	emp6 := Employee{
		firstName: "Rama",
		lastName: "Krishna",
		age: 22,
		salary: 20000,
	}
	fmt.Println("First name: ", emp6.firstName)
	fmt.Println("Last name: ", emp6.lastName)
	fmt.Println("Age: ", emp6.age)
	fmt.Println("Salary: ", emp6.salary)

	/*
		It is also possible to create a zero struct.
		And then assign values to its fields later
	*/
	fmt.Println("Creating a zero struct and assigning values later.")
	var emp7 Employee
	emp7.firstName = "John"
	emp7.age = 34
	fmt.Println(emp7)

	/*
		Pointers to a struct. It is also possible to create pointers to a struct.
	*/
	fmt.Println("Pointers to a struct")
	emp8 := &Employee{"Randy", "Orton", 40, 120000}
	fmt.Println("Firstname: ", (*emp8).firstName) // explicit dereferencing
	fmt.Println("Lastname: ", (*emp8).lastName) // explicit dereferencing

	// in go we can use this emp.firstName instead of explicit dereferencing
	fmt.Println("Without using explicit dereferencing ")
	fmt.Println("Firstname: ", emp8.firstName)
	fmt.Println("Age: ", emp8.age)

	/*
		Anonymous fields.
		It is possible create structs with fields which only contain type without the field name.
	*/
	fmt.Println("Using Anonymous fields")
	p := Person{"Krishna", 22}
	fmt.Println(p)

	// Even though an anonymous fields does not have a name, 
	// by default the name of a anonymous field is the name of its type.
	fmt.Println("For Anonymous fields, the field names are the types.")
	var p1 Person
	p1.string = "Krishna"
	p1.int = 22
	fmt.Println(p1)
}