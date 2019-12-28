package main

import (
	"fmt"
)

/*
	What is panic?

	The idiomatic way of handle abnormal conditions in a program in GO is using errors.
	Errors are sufficient for most of the abnormal conditions arising in the program.

	But there are some situations where the program cannot simply continue executing after an 
	abnormal situation. In this case we use panic to terminate the program.

	when a function encounters a panic its execution is stopped, any deferred functions are executed
	and then the control returns to its caller.

	This process continues until all the functions of the current Goroutine have returned at which 
	point the program prints the panic message, followed by the stack traces and then terminates.

	when panic should be used?
	One important factor is that we should avoid panic and recover and use errors where ever possible.
	Only in cases where the program just cannot continue execution should a panic and recover mechanism
	be used.

	There are 2 valid use cases for panic.

	* An unrecoverable error where the program cannot simply continues its execution.
	One example would be a web server which fails to bind to the required port. In this case its
	reasonable to panic as there is nothing else to do if the port binding itself fails.

	* A programmer error.
	Lets say we have a method which accepts a pointer as a parameter and someone calls this method
	using nil as argument. In this case we can panic as its a programmer error to call a method with 
	nil argument which is expecting a pointer.

	The signature of the build-in panic function
	func panic({}interface)
*/

// example of panic
func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s \n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main()  {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	// lastName := "Musk"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	/*
		The above is a simple program to print the fullname of the person. 
		This function(fullName) checks whether the firstName and lastName pointers are nil.
		If it is nil the function calls panic with corresponding error message. 
		The error message will be printed the program terminates.

		when the panic is encountered the program execution terminates, prints the argument passed to
		panic followed by the stack trace 
	*/

	/*
		Defer while panicking.

		lets modify the above function(fullName) by including defer function calls.

		In this case we have added defer functions in line no 44 and 56. The defer statement in line
		no 44 is executed first. This prints

		"deferred call in fullName"

		And then the control returns to the main function whose deferred calls are executed and
		hence this prints

		"deferred call in main"

		Now the control has reached to the top level of the function and hence the program prints the
		panic message followed by the stack trace and then terminates.

	*/
}