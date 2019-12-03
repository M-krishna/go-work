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
	
}
