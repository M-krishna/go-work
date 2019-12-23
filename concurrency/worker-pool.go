package main

import (
	"fmt"
	"time"
	"sync"
)

/*
	To understand worker pools, we need to first know about WaitGroup as it will be used for the
	implementation of worker pool.

	A WaitGroup is used to wait for a collection of Goroutines to finish executing.
	The control is blocked until all Goroutines finish executing.
	Lets say we have 3 concurrently executing Goroutines spawned from the main Goroutine.
	The main Goroutine needs to wait for the other 3 Goroutines to finish before terminating.
	This can be accomplished using WaitGroup.
*/

func Process(i int, wg *sync.WaitGroup){
	fmt.Println("started Goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main(){
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go Process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All Goroutines finished executing.")

	/*
		WaitGroup is a struct type and we are creating a zero value variable of type WaitGroup.

		The way WaitGroup works is by using a counter. When we call Add on the WaitGroup and pass it 
		an int, the WaitGroup's counter is incremented by the value passed to Add.

		The way to decrement the counter is by calling Done() method on WaitGroup.

		The Wait() method blocks the Goroutine in which its called until the counter becomes zero.
	*/

	/*
		In the above program, we call wg.Add(1) inside the for loop which iterates 3 times. 
		So the counter now becomes 3. 
		The for loop also spawns 3 process Goroutines and then wg.Wait() called makes the main
		Goroutine to wait until the counter becomes zero.
		The counter is decremented by the call to wg.Done() in the Process Goroutine.
		Once all the 3 spawned Goroutines finishes their execution, that is once wg.Done() has been 
		called 3 times, the counter will become zero and the main Goroutine will be unblocked.
	*/

	/*
		It is important to pass the address of wg in line no. 32. If the address is not passed then
		each Goroutine will have its own copy of the WaitGroup and main Goroutine will not be 
		notified when they finish executing.
	*/
}