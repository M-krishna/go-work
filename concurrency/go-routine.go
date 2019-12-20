package main

import (
	"fmt"
	"time"
)

/*
	Go is a concurrent language not a parallel one.

	Concurrency:
	Concurrency is the capability to deal with lots of things at once.

	Parallelism:
	Parallelism is doing lots of things at the same time.

	Concurrency is handled by GoRoutines in Go.
*/

/*
	What are Go Routines?
	 Goroutines are functions or methods that run concurrently with other functions or methods.
	 Goroutines can be thought of as a light weight thread.
	 The cost of creating a Goroutine is tiny when compared to a thread.
	 Hence its common for Go applications to have thousands of Goroutines running concurrently.
*/

/*
	Advantages of Goroutines over threads.
		* Goroutines are extremely cheap when compared to threads. They are only a few KB in stack size
		and the stack can grow and shrink according to the needs of the application. In case of threads
		the stack size has to be specified and fixed.

		* The Goroutines are multiplexed to fewer number of OS threads. There might be only one thread 
		in a program with thousands of Goroutines. If any Goroutine in that thread blocks lets say waiting
		for user input, then another OS thread is created and remaining Goroutines are moved to the new
		OS thread. All these are taken care by the runtime.

		* Goroutines communicate using channels. Channels by design prevent race conditions from happening
		when accessing shared memory using Goroutines. Channels can be thought of as a pipe using which
		Goroutines communicate. 
*/

func hello() {
	fmt.Println("Hello World")
}

func numbers(){
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func letters(){
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}


func main(){
	/*
		How to start a Goroutine?
		prefix the function or method call with keyword go and you will have a new Goroutine 
		running concurrently.
	*/
	/* go hello()
	fmt.Println("Main Function")*/

	/*
		The above snippet outputs only "Main Function".
		what happened to the Goroutine we started??

		* when a new Goroutine is started, the goroutine call returns immediately. Unlike functions,
		the control does not wait for the Goroutine to finish executing. The control returns 
		immediately to the next line of code after the Goroutine call and any returns from the 
		Goroutines are ignored.

		* The main Goroutine should be running for any other Goroutines to run.
		If the main Goroutine terminates then the program will be terminated and no other Goroutine
		will run.
	*/

	// lets rewrite the above code to fix that.
	/*go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("Main Function")*/

	/*
		In the above snippet we put the main Goroutine to sleep for 1s.
		The call to the Goroutine hello has enough time to execute before the main Goroutine terminates.
		This will first print "Hello World" and "Main Function".
	*/

	/*
		This way of using sleep in main Goroutine is a way of hack to show how Goroutines work.
		We'll use channels to block the main Goroutine from executing before other Goroutines.
	*/

	/*
		Starting Multiple Goroutines.
	*/
	go numbers()
	go letters()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Main Function")
}