package main

import (
	"fmt"
	"time"
)

/*
	Buffered Channel.
	All channels we have seen previously were basically unbuffered. 
	Sends and receives to an unbuffered channel are blocking. 

	It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only
	when the buffer is full.
	Similarly receives from a buffered channel are blocked only when the buffer is empty.

	Buffered channels can be created by passing an additional capacity parameter to the make 
	function which specifies the size of the buffer.

	ch := make(chan type, capacity)

	capacity in the above syntax should be greater than 0 for channel to have a buffer. The capacity 
	for an unbuffered channel is 0 by default hence we omitted the capacity parameter while creating
	channels.
*/

// write function 
func write(ch1 chan int){
	for i := 0; i < 5; i++ {
		ch1 <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch1)
}

func main() {
	ch := make(chan string, 2)
	ch <- "Krishna"
	ch <- "Murugan"
	fmt.Println(<- ch) // Krishna
	fmt.Println(<- ch) // Murugan

	/*
		In the above snippet we have created a buffered channel with a capacity of 2.
		Since the channel has the capacity of 2, it is possible to write 2 strings into the channel
		without being blocked.
	*/

	// Another Example
	ch1 := make(chan int, 2)
	go write(ch1)
	time.Sleep(2 * time.Second)
	for v := range ch1 {
		fmt.Println("Reading value ", v, "from ch1")
		time.Sleep(2 * time.Second)
	}

	/*
		In the above program, a buffered channel ch of capacity 2 is created in the main Goroutine
		and passed to the write Goroutine.

		Then the main Goroutine sleeps for 2 seconds. During this time the write Goroutine is running
		concurrently.

		The write Goroutine has a for loop which writes numbers 0 to 4 to the ch1 channel. The capacity
		of this buffered channel is 2 and hence the write Goroutine will be able to write values from 
		0 and 1 to the ch channel immediately and then it blocks until at least 1 value is read from
		ch channel.

		so this program will print the following lines immediately.
		successfully wrote 0 to ch  
		successfully wrote 1 to ch

		After printing the above two lines, the writes to the ch channel in the write Goroutine are 
		blocked until someone reads from the ch channel.
		
		Since the main Goroutine sleeps for 2 seconds before starting to read from the channel, the 
		program will not print anything for next 2 seconds.

		The main Goroutine wakes up after 2 seconds and starts reading from the ch channel using 
		for range loop, prints the read value and then sleeps again for 2 seconds and this cycle 
		continues until the ch is closed.

		So the program will print the following 2 lines after 2 seconds.
		read value 0 from ch  
		successfully wrote 2 to ch 
	*/

	/*
		Deadlock
	*/
	fmt.Println("Deadlock")
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	//ch2 <- 3
	fmt.Println(<- ch2)
	fmt.Println(<- ch2)
	//fmt.Println(<- ch2)

	/*
		In the above snippet we write 3 integers to a buffered channel of capacity 2. 
		when the control reaches the 3rd line, the write is blocked since the channel has exceeds its 
		capacity. 
		Now some Goroutine must read from the channel in order for the write to proceed.
		But in this case there is no concurrent routine reading from this channel.
		Hence there will be a deadlock and the program will panic at runtime.
	*/

	/*
		Length vs Capacity
		The capacity of a buffered channel is the number of values that the channel can hold.
		This is the value we specify when creating the buffered channel using make function.

		The Length of the buffered channel is the number of elements currently queued in it.
	*/
	fmt.Println("Capacity vs Length")
	ch3 := make(chan string, 3)
	ch3 <- "Krishna"
	ch3 <- "Murugan"
	fmt.Println("Capacity of a channel", cap(ch3))
	fmt.Println("Length of a channel", len(ch3))
	fmt.Println("Read value from a channel", <- ch3)
	fmt.Println("New Length of a channel", len(ch3))
}