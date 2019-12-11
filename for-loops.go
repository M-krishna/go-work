package main

import "fmt"

func main(){
	//for loop
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\n", i)
	}

	// break statement
	// breaks the for loop from executing
	// goes to the next line of the for loop
	for i := 0; i <= 10; i++ {
		if i > 5{
			break
		}
		fmt.Printf(" %d", i)
	}
	fmt.Printf("Line after for loop\n")

	//nested for loop
	fmt.Println("Nested For loop")
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++{
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	//labels
	// labels can be used to break the outer loop
	// in the above nested loop if I want to break the loop
	// if i and j are equal
	fmt.Println("Nested For Loop with break")
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++{
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				break
			}
		}
	}

	// in the above code break only breaks the inner for loop
	// but we need to break the outer for loop
	// so we can use labels
	fmt.Println("Using labels")
	outer:
		for i := 0; i < 3; i++ {
			for j := 1; j < 4; j++{
				fmt.Printf("i = %d, j = %d\n", i, j)
				if i == j{
					break outer
				}
			}
		}

	// in for loop initialization and post part can be omitted
	// like while loop
	i := 0
	for i <= 10 {
		fmt.Printf(" %d\t", i)
		i += 2
	}

	// is it possible to declare multiple variables in for loop
	// the below program prints this sequence
	/*
		10 * 1 = 10
		11 * 2 = 22
		.
		.
		.19 * 10 = 190
	*/

	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1{
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	// infinite loop
	for {
		fmt.Printf("Hello World")
	}
	// the above piece of code will run infinitely
}