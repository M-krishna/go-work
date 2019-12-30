package main

import (
	"fmt"
	"time"
	"runtime/debug"
)

/*
	recover is a builtin function which is used to regain control of a panicking goroutine.

	The signature of recover function is 
	func recover() interface{}

	Recover is useful only when called inside deferred functions.
	Executing a call to recover inside a deferred a function stops the panicking sequence by restoring
	normal execution and retrives the error value passed to the call of panic.

	If recover is called outside the deferred function, it will not stop a panicking sequence.
*/

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstName *string, lastname *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: firstname cannot be nil")
	}

	if lastname == nil {
		panic("runtime error: lastname cannot be nil")
	}

	fmt.Printf("%s %s\n", *firstName, *lastname)
	fmt.Println("returned normally from fullName")
}

/* Panic, Recover and Goroutines*/

func recovery(){
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func a(){
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside b")
	panic("oh! B panicked")
}

/* End of Panic, Recover and Goroutines*/

/* Runtime Panics */

func r(){
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}
func aa(){
	defer r()
	a := []int{1,2,3}
	fmt.Println(a[3])
	fmt.Println("normally returned from aa")
}
/* End of Runtime Panics*/

func main()  {
	/*defer fmt.Println("deferred call in main")
	firstName := "Elon"
	lastName := "Musk"
	fullName(&firstName, &lastName)
	fmt.Println("returned normally from main")*/

	/*
		The recoverName() function calls the recover() which returns the value passed to the call of
		panic.

		when fullName panics, the deferred function recoverName() will be called which uses recover()
		to stop the panicking sequence.
	*/

	/*
		Panic, Recover and Goroutines.

		Recover works only when it is called from the same goroutine. It's not possible to recover 
		from a panic that has happened in a different goroutine.
	*/
	// a()

	/*
		In the above program the function b() panics and the function a() calls a deferred function
		recovery() which is used to recover from the panic.

		The function b() is called as a separate Goroutine and the sleep in the next line is just to
		ensure that the program does not terminate before b() has finished running.

		The panic will not be recovered. This is because the recovery function is present in a 
		different Goroutine and the panic is happening in function b() in different Goroutine.

		Hence the recovery is not possible.

		If the function b() is called in the same Goroutine the panic would have been recovered.
		If we change go b() to b() in line no. 52 the recovery will happen now since the panic is 
		happening in the same Goroutine.
	*/

	/*
		Runtime Panics.
		Panics can also be caused by runtime errors such as array out of bounds access. 
		Is it possible to recover from runtime panics. Yes its possible.
		Lets write a program. 
	*/
	aa()

	/*
		Getting stack trace after recover.

		If we recover a panic, we loose the stack trace from the panic. Is it possible to print the 
		stack trace using the PrintStack function from the debug package.

		Lets print the stack trace of the above program after recover.
	*/
}