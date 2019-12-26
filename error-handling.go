package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
	Errors in Go are plain old values. Errors are represented using the built in error type.

	Just like other builtin types error values can be stored in variables, returned from functions and
	so on.
*/

func main(){
	if len(os.Args) != 2{
		os.Exit(1)
	}
	
	n, err := strconv.Atoi(os.Args[1])
	if err != nil{
		fmt.Println("not a valid number")
	} else{
		fmt.Println(n)
	}

	fs, err := os.Open("/test.txt")
	/*if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(fs.Name(), "opened successfully")*/

	/*
		If the file has been opened successfully, then the open function will return the file handler
		and error will be nil. If there is an error while opening the file a non nil error will be
		returned.

		If a function or a method returns error, then by convention it has to be the last value returned
		from the function. 

		The idiomatic way of handling error in Go is to compare the returned error to nil.
		A nil value indicates that no error has occurred and a non nil value indicates the presence of an
		error.
	*/

	/*
		Error Type Representation.
		error is an interface type with the following definition.

		type error interface{
			Error() string
		}

		It contains a single method with signature Error() string. Any type which implements this 
		interface can be used as an error. This method provides the description of the error.

		when printing the error the fmt.Println function calls the Error() string method internally
		to get the description of the error.
	*/

	/*
		Different ways to extract more information from the error.

		In the Eg. above we printed the description of the error. What if we wanted the actual path
		of the file which caused the error. One possible way is parsing the error string.

		Asserting the underlying struct type and getting more information from the struct fields.

		Lets rewrite the above program.
	*/

	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return 
	}
	fmt.Println(fs.Name(), "opened successfully")

	// in the above program we use type assertion to get the underlying value of the error interface.
	// Then we print the path using err.Path
}
