package main

import "fmt"

func number() int{
	num := 15 * 5
	return num
}

func main(){
	
	// switch statement
	// the example below prints the finger name
	finger := 4
	switch finger{
		case 1:
			fmt.Println("Thumb")
		case 2:
			fmt.Println("Index")
		case 3:
			fmt.Println("Middle")
		case 4:
			fmt.Println("Ring")
		case 5:
			fmt.Println("Pinky")
	}

	// using default case
	switch finger := 8; finger { // optional statement, which is executed before the expression is evaluated
		case 1:
			fmt.Println("Thumb")
		case 2:
			fmt.Println("Index")
		case 3:
			fmt.Println("Middle")
		case 4:
			fmt.Println("Ring")
		case 5:
			fmt.Println("Pinky")
		default:
			fmt.Println("Invalid Finger number")
	}

	// multiple expressions in a case
	var letter = "i"
	switch letter {
		case "a", "e", "i", "o", "u": // multiple expressions separated by comma
			fmt.Println("Vowel")
		default:
			fmt.Println("Not a vowel")
	}

	// expressionless switch statement
	// expressions in a switch is optional
	num := 75
    switch { // expression is omitted
		case num >= 0 && num <= 50:
			fmt.Println("num is greater than 0 and less than 50")
		case num >= 51 && num <= 100:
			fmt.Println("num is greater than 51 and less than 100")
		case num >= 101:
			fmt.Println("num is greater than 100")
	}
	
	// fallthrough
	// the control comes out of the switch statement after the
	// case is executed
	/*
		A fallthrough statement is used to transfer control to the first 
		statement of the case that is present immediately after the case 
		which has been executed.
	*/

	switch num := number(); {
		case num < 50:
			fmt.Printf("%d is lesser than 50\n", num)
			fallthrough
		case num < 100:
			fmt.Printf("%d is lesser than 100\n", num)
			fallthrough
		case num < 200:
			fmt.Printf("%d is lesser than 200", num)
	}
}