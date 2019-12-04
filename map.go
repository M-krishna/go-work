package main

import "fmt"

func main(){
	var m map[string]int // declared a map(empty - nil)
	// syntax in map[keyType]valueType
	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is initialised")
	}
	
	// creating a map using make function
	r := make(map[string]int) // initializing a empty map using make func.
	r["power"] = 20
	r["age"] = 26
	fmt.Println(r)

	// initializing a map normally
	var mm =  map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	fmt.Println(mm)

	// another example
	var tindermatch = make(map[string]string)
	tindermatch["Krishna"] = "Angelina"
	tindermatch["Abhishek"] = "Ankita"
	fmt.Println(tindermatch)
	tindermatch["krishna"] = "Emma Watson"
	fmt.Println(tindermatch)

	// retrieving the value using key
	personMobileNo := make(map[string]string)
	personMobileNo["Krishna"] = "9876543210"
	personMobileNo["Abhishek"] = "1234567890"
	mobileNo := personMobileNo["Krishna"]
	fmt.Println("Krishna's mobile no: ", mobileNo)
	// if the key doesn't exist
	fmt.Println("Ravi's mobile no: ", personMobileNo["ravi"])

	// checking if a key really exists or not
	// because if the value of the key is empty
	// its gonna return the type of the map
	name, number := personMobileNo["Kathir"]
	fmt.Println(name, number)

	// To find the length of the map
	// we can use len function
	fmt.Println("Length of the Map: ", len(personMobileNo))

	// To delete a key in the map
	// we can use delete function
	// its not gonna return anything
	// we can use this on non exists key
	delete(personMobileNo, "Krishna")
	fmt.Println(personMobileNo, len(personMobileNo))

	// Maps grow dynamically
	// however if we want, we can supply 2nd argument to make 
	// to set an initial size
	// defining an initial size can help with performance
	lookup := make(map[string]int, 100)
	fmt.Println(lookup)

	// when you need a map as a field of a structure
	// we can define it as
	type Saiyan struct{
		Name string
		Friends map[string]Saiyan
	}

	goku := Saiyan{
		Name: "Goku",
		Friends: make(map[string]Saiyan),
	}
	piccolo := Saiyan{
		Name: "Piccolo",
		Friends: make(map[string]Saiyan),
	}
	goku.Friends["Piccolo"] = piccolo
	fmt.Println(goku)
	// not so clear about what im doing in the above chunk of code.
}
