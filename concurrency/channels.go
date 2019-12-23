package main

import (
	"fmt"
)

/*
	Channels.
	Channels can be thought as pipes using which Goroutines communicate. Similar to how water flows
	from one end to another in a pipe, data can be sent from one end and received from the another end
	using channels.
*/

func hello(done chan bool){
	fmt.Println("Hello World")
	// time.Sleep(4 * time.Second) // waits for 4 seconds before writing to the channel.
	done <- true // writing to the channel.
}

// function for calculating squares
func squaresch(number int, sqch chan int){
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	sqch <- sum
}

// function for calculating cubes
func cubesch(number int, cubch chan int){
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubch <- sum
}

// a function that only sends data to a channel
func senddata(sendch chan<- int){
	sendch <- 10
}

// closing a channel.
func Producer(ch chan int){
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // closing a channel
}

/*
	Rewriting the square and cubes function for more code reusability.
*/
func digits(number int, dchcl chan int){
	for number != 0 {
		digit := number % 10
		dchcl <- digit
		number /= 10
	}
	close(dchcl)
}

func calculateSquare(number int, schcl chan int){
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	schcl <- sum
}

func calculateCubes(number int, cubchcl chan int){
	sum := 0
	sch := make(chan int)
	go digits(number, sch)
	for digit := range sch{
		sum += digit * digit * digit
	}
	cubchcl <- sum
}

/* End of rewriting */

func main()  {

	/*
		Declaring Channels.
		Each channel has a type associated with it. This Type is the Type of the data that the 
		channel is allowed to transport.
		No other type is allowed to be transported using the channel.

		chan T is the channel of Type T.

		The zero value of the channel is nil. nil channels are not of any use and hence the channel 
		has to be defined using make similar to maps and slices.
	*/

	var a chan int
	fmt.Println("Declaring Channel")
	fmt.Printf("Type of channel a is %T, value of channel a is %v\n", a, a)
	if a == nil {
		fmt.Println("Channel is nil, going to define it.")
		a = make(chan int)
		fmt.Printf("Type of channel a is %T, value of channel a is %v\n", a, a)
	}

	/*
		Sending and Receiving from a channel.
		The syntax,

		data := <- a (read from a channel)
		a <- data (write to a channel)

		The direction of the arrow with respect to the channel specifies whether the data is sent or
		received.
		
		In line no. 40 the arrow points outwards from a and hence we are reading channel a and storing
		the value to the variable data.

		In line no. 41 the arrow points towards a and hence we are writing to a channel.
	*/

	/*
		Sends and Receives are blocking by default. What the heck is this????
		When data is sent to a channel, the control is blocked in the send statement until someother
		Goroutine reads data from that channel.
		Similarly when data is read from a channel, the read is blocked until some other Goroutine
		writes data to that channel.

		This property of channels is what helps Goroutine communicate effectively without the use of
		explicit locks or conditional variables.
	*/

	// Channel example program.
	// Modifying the same program used in the go-routine.go
	done := make(chan bool) // making a boolean channel
	go hello(done) // passing the channel to hello function.
	<- done // blocking code. (after getting data from the channel this becomes unblocking.)
	fmt.Println("Main Goroutine")


	/*
		Another example for channels.
		Lets write a program that calculate the sum of squares and cubes of a number.
		(Eg.) num = 123
		squares = (1 * 1) + (2 * 2) + (3 * 3)
		cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3)
		sum = square + cubes
	*/
	num := 589
	sqch := make(chan int)
	cubch := make(chan int)
	go squaresch(num, sqch)
	go cubesch(num, cubch)
	squares, cubes := <- sqch, <- cubch
	fmt.Println("The sum of squares and cubes of 123 is ", squares + cubes)
	
	/*
		DeadLock.
		One other important factor to consider while using channels is deadlock.
		If a Goroutine is sending data on a channel, then it is expected that some other Goroutine 
		should be receiving the data. 
		If this does not happen then the program will panic at runtime with Deadlock.
		The below code explains Deadlock
	*/
	// ch := make(chan int)
	// ch <- 5 // cause deadlock

	/*
		Because in the above code, channel ch is created and a value is passed to the channel.
		But no other Goroutines are receiving data from channel ch. Hence this causes deadlock.
	*/

	/*
		Unidirectional Channels.

		So far we have created bidirectional channels that is data can be both send and received on
		them.

		It is also possible to create unidirectional channels, that is channels can only send or receive
		data.
	*/
	sendcha := make(chan<- int)
	go senddata(sendcha)
	// fmt.Println(<-sendcha) // throws an error.

	/*
		We try receive data from a send only channel. This is not allowed when the program is run.
		The compiler will complain stating,
		invalid operation: <-sendcha (receive from send-only type chan<- int)

		But what is the point of creating a send only channel if it cannot be read from.

		This is where channel conversations comes to use. It is possible to convert a bidirectional 
		channel to a send only or receive only channel but not vice versa.
	*/
	sendcha2 := make(chan int) // creating a bidirectional channel
	go senddata(sendcha2) // The senddata function converts this channel into a send only channel in line no. 43.
	// so the channel is sendonly inside the senddata Goroutine and bidirectional in main Goroutine. 
	fmt.Println(<- sendcha2) // receiving value from the send only channel.

	/*
		Closing channels and for range loops on channels.

		Senders have the ability to close the channels to notify the receivers that no more data will
		be sent on the channel.

		Receivers can use additional variable while receiving data from the channel to check whether
		the channel has been closed.

		v, ok := <- ch

		In the above statement ok is true if the value was received by a successful send operation
		to a channel. If ok is false means, that we are reading from a closed channel.
		The value read from a closed channel will be the zero value of the channels type.
		For example, if the channel is an int channel, then the value received from the closed channel
		will be 0.
	*/
	ch1 := make(chan int)
	go Producer(ch1)
	for {
		v, ok := <- ch1
		if !ok {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	// lets rewrite the above program using for range
	ch2 := make(chan int)
	go Producer(ch2)
	fmt.Println("Using for range")
	for v := range ch2{ // automatically breaks after the channel is closed.
		fmt.Println("Received ", v)
	}

	/*
		Lets rewrite the squares and cubes program using for range loop for more code reusability.
	*/
	number := 123
	squarech := make(chan int)
	cubech := make(chan int)
	go calculateSquare(number, squarech)
	go calculateCubes(number, cubech)
	squares1, cubes1 := <- squarech, <- cubech
	fmt.Println("Sum of squares and cubes after rewriting ", squares1 + cubes1)
}