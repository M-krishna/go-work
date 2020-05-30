package main

import (
	"fmt"
	"sync"
	"time"
)

// Channels are a way to communicate between different goroutines

func main() {
	// Here we are calling the same function twice with different values
	// If you notice the function has an infinite for loop and it sleeps for a second before executing
	// Which means the function Hello("Kathir") wont execute
	// HelloV1("Krishna")
	// HelloV1("Kathir")

	// We can fix this by adding a go keyword
	// go HelloV1("Krishna")
	// go HelloV1("Kathir")
	// Now if you try to run the code, the program exits quickly
	// its because the main function in go is a goroutine
	// it wont wait for other goroutines to finish

	//There are several ways to fix this
	// 1st method we can use a WaitGroup from sync package
	//var wg sync.WaitGroup
	//wg.Add(1) // This is a counter, if this becomes 0, the main goroutine stops executing
	//go HelloV2("Krishna", &wg)
	//go HelloV2("Kathir", &wg)
	//wg.Wait() // Here we are telling the main goroutine to wait for the other goroutines to finish

	// The above method is not so useful, so solve the problem we can use channels
	//c := make(chan string) // creating an unbuffered channel
	//go HelloV3("Krishna", c)
	// msg := <- c
	// fmt.Println(msg)
	// If you run the above code, it stops execution after receiving one value from the channel
	// instead of receiving values 5 times (for loop)
	// what we can do is use a for loop
	//for {
	//	msg, open := <- c // The receiving channel also tell us if the channel is open/close
	//
	//	if !open {
	//		break
	//	}
	//	fmt.Println(msg)
	//}

	// The above for loop is quite messy, we can refactor that
	// we can use range keyword on channel
	//for msg := range c {
	//	fmt.Println(msg)
	//}


	// Lets look up another scenario
	//c := make(chan int) // unbuffered channel
	//c <- 1

	//msg := <-c
	//fmt.Println(msg)
	// If you the above code it'll cause our program to deadlock
	// because after sending a value to our channel, the send will block until something is received from the channel
	// so the code wont progress to the next line, hence cause the deadlock
	// to fix this we can spawn a separate goroutine for receiving from the channel.

	// Instead of doing that we can create a buffered channel
	//c := make(chan int, 2)
	//c <- 1

	//msg := <- c
	//fmt.Println(msg)

	// In the above code we created a buffered channel with a capacity of 2
	// so the sender wont block the code, since its a buffered channel
	// but we can only send two(capacity) values at a time
	//c <- 2
	//msg1 := <- c
	//fmt.Println(msg1)

	// We can use select statements for solving other problems in channels
	c1 := make(chan string)
	c2 := make(chan string)
	//
	go forEveryHalfMilliSeconds("Krishna", c1)
	go forEveryTwoSeconds("Kathir", c2)
	//
	//for {
	//	fmt.Println(<-c1)
	//	fmt.Println(<-c2)
	//}

	// If you run the above program you can notice the first goroutine is blocked by 2nd goroutine because of
	// the time interval
	// To fix that we can use select statement
	for {
		select {
		case msg := <- c1:
			fmt.Println(msg)
		case msg := <- c2:
			fmt.Println(msg)
		}
	}

	// If you run the above code you can notice that the 1st goroutine function value gets printed
	// after 500ms, the next one gets printed after 2 seconds


}

func HelloV1(name string) {
	for i := 0; true; i++ {
		fmt.Printf("Hello %v", name)
		time.Sleep(time.Second)
	}
}

// 2nd version of Hello function which accept a waitgroup
func HelloV2(name string, wg *sync.WaitGroup){
	for i := 0; i < 5; i++ {
		fmt.Printf("Hello %v\n", name)
		time.Sleep(time.Second)
	}
	wg.Done()
}


// 3rd version of Hello function which accepts channel
func HelloV3(name string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- "Hello " + name // Here we are writing to the channel
		time.Sleep(time.Second)
	}
	close(c)
}

func forEveryHalfMilliSeconds(name string, c1 chan string) {
	for {
		c1 <- name
		time.Sleep(time.Millisecond * 500)
	}
}

func forEveryTwoSeconds(name string, c2 chan string) {
	for {
		c2 <- name
		time.Sleep(time.Second * 2)
	}
}
