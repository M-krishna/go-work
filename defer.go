package main

import (
	"fmt"
	"sync"
)

/*
	Defer.
	Defer statement is used to execute the function call just before the function where the defer 
	statement is present returns.
*/

type Person struct{
	firstname, lastname string
}

// method to get the fullname of the person
func (p Person) getFullName() {
	fmt.Printf("%s %s\n", p.firstname, p.lastname)
}

func Finished(){
	fmt.Println("Finished finding largest number")
}

func FindLargestNumber(num []int) {
	defer Finished()
	fmt.Println("started finding largest number")
	max := num[0]
	for _, v := range num {
		if v > max{
			max = v
		}
	}
	fmt.Println("The largest number is ",max)
}

// Test function for evaluating arguments passed to a deferred function
func testArguments(num int){
	fmt.Println(num)
}

/* Practical use of defer */

type rect struct{
	length, width int
}

func (r rect) area(wg *sync.WaitGroup){
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	wg.Done()
} 

/*
	In the above function we can clearly see that these calls happen just before the area method 
	returns. wg.Done() should be called before the method returns hence these calls can be 
	effectively replaced by a single defer call.
*/

func (r rect) Area(wg *sync.WaitGroup){
	defer wg.Done()
	if r.length < 0 {
        fmt.Printf("rect %v's length should be greater than zero\n", r)
        return
    }
    if r.width < 0 {
        fmt.Printf("rect %v's width should be greater than zero\n", r)
        return
    }
    area := r.length * r.width
    fmt.Printf("rect %v's area %d\n", r, area)
}

/* End of practical use case*/

func main(){
	num := []int{1,2,3,4,5,6,7}
	FindLargestNumber(num)

	/*
		Deferred methods.
		
		Defer is not restricted only to functions. It is perfectly legal to defer a method call too.
	*/
	p := Person{"Krishna", "Murugan"}
	defer p.getFullName()
	fmt.Printf("Welcome ")

	/*
		Arguments Evaluation.
		The arguments of a deferred function are evaluated when the defer statement is executed and
		not when the actual function call is done.
	*/
	a := 5
	defer testArguments(a)
	a = 10
	testArguments(a)
	fmt.Println("Final value of a ", a)

	/*
		Stack of defers.
		when a function has multiple defer calls, they are added on to a stack and executed in
		Last In First Out order(LIFO).

		Lets reverse a string using defer.
	*/
	str := "Krishna"
	fmt.Println("Original string ", str)
	fmt.Println("Reversed string ")
	for _, v := range []rune(str){
		defer fmt.Printf("%c", v)
	}

	/*
		Practical use of defer.
	*/
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
    r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1,r2,r3}
	for _, v := range rects{
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")

	fmt.Println("Using Defer")
	r4 := rect{-67, 89}
    r5 := rect{5, -67}
	r6 := rect{8, 9}
	rects1 := []rect{r4,r5,r6}
	for _, v := range rects1{
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("Using defer in wg.Done()")

}