package main

import (
	"fmt"
	"time"
)

/*
	Select Statement.
	The select statement is used to choose from multiple send/receive statements.
	The select statement blocks until one of the send/receive operation is ready.
	If multiple operations are ready then one of them is chosen at random.
	The syntax is similar to switch except that each of the case statement will be a channel operation.
*/

func server1(ch chan string){
	time.Sleep(6 * time.Second)
	ch <- "from server 1"
}

func server2(ch chan string){
	time.Sleep(3 * time.Second)
	ch <- "from server 2"
}

func process(ch chan string){
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

/* Example functions for random selection */

func server3(ch chan string){
	ch <- "from server 3"
}

func server4(ch chan string){
	ch <- "from server 4"
}

/* End of random selection*/


func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
		case s1 := <- output1:
			fmt.Println(s1)
		case s2 := <- output2:
			fmt.Println(s2)
	}

	/*
		In the above code the server1 function in line 16 sleeps for 6 seconds then writes the text
		"from server 1" to ch.
		The server2 function in line 21 sleeps for 3 seconds then writes the text "from server 1" to
		ch.

		The main function calls the server1 and server2 in line no. 30 & 31.

		In line 32 the control reaches the select statement. The select statement blocks until one of
		its cases is ready. In the above program the server1 writes after 6 seconds and server2 writes
		after 3 seconds.

		So the select statement will block for 3 seconds and will wait for server 2 Goroutine to 
		write to the output2 channel. After 3 seconds the program will print "from server 2"
	*/
	/*ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
			case v := <- ch:
				fmt.Println("value received ",v)
				return
			default:
				fmt.Println("No value received")
		}
	}*/

	/*
		In the above program, the process function sleeps for 10500 milliseconds(10.5 seconds) then 
		writes process successful to the ch channel. This function is concurrently called in line no.
		60.

		after calling the process Goroutine concurrently, an infinite loop is started in the main
		Goroutine. The infinite loop sleeps for 1000 Millisecond(1 second) during the start of the 
		iteration and then performs the select statement.

		During the first 10500 milliseconds, the first case of the select statement namely v := <-ch
		will not be ready since the process Goroutine will write to the channe ch after 10500 
		milliseconds. Hence the default case is executed during this time and the program will
		print No Value received 10 times.

		after 10.5 seconds, the process Goroutine writes process successful to the channel ch.
		Now the first case of the select statement will be executed and the program will print 
		value received process successful and then it will terminate.
	*/


	/*
		Deadlock and default case.
	*/
	/*ch1 := make(chan string)
	select {
		case <- ch1:
	}*/

	/* The above snippet will cause a deadlock because we try to read from the channel. The 
		select statement will block forever since no other Goroutine is writing to this channel and 
		hence will result in deadlock in runtime.
	*/

	/*
		If a default case is present then the deadlock will not happen since the default case will be
		executed when no other case is ready. The above program can be rewritten as 
	*/
	ch1 := make(chan string)
	select {
		case <- ch1:
		default:
			fmt.Println("default case executed")
	}

	/*
		Similarly the default case is executed even if the select has only nil channels.
	*/
	var ch2 chan string
	select {
		case v1 := <- ch2:
			fmt.Println("value received ", v1)
		default:
			fmt.Println("default value received")
	}

	/*
		Random Selection.
		when multiple cases in the select statement are ready, one of them will be executed at random
	*/
	output3 := make(chan string)
	output4 := make(chan string)
	go server3(output3)
	go server4(output4)
	time.Sleep(1 * time.Second)
	select {
		case v1 := <- output3:
			fmt.Println("Value received from server 3", v1)
		case v2 := <- output4:
			fmt.Println("Value received from server 4", v2)
	}

	/*
		Empty Select.
	*/
	// select{} // throws an deadlock error.
	/*
		we know the select statement will block until one of the case gets executed.
		 In this case the select statement doesn't have any cases and hence it will block forever
		 resulting in a deadlock.
	*/ 
}