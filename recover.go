package main

import "fmt"

/*
	recover is a builtin function which is used to regain control of a panicking goroutine.

	The signature of recover function is 
	func recover() interface{}

	Recover is useful only when called inside deferred functions.
	Executing a call to recover inside a deferred a function stops the panicking sequence by restoring
	normal execution 
*/

func main()  {
	
}