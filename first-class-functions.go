package main

import (
	"fmt"
)

/*
	First Class Functions.

	A language which supports first class functions allows functions to be assigned to variables,
	passed as arguments to other functions and returned from other functions 
*/

/* Creating our own function types. */
type add func(a int, b int) int
/* End of creating our own function types*/

/* Creating Higher order functions*/
func simple(a func(a int, b int) int){
	fmt.Println(a(60, 7))
}
/* End of Higher order functions*/

/* Returning a function from another function */
func simple2() func(a int, b int) int {
	return func(a, b int) int {
		return a + b
	}
}
/* End of Returning function from another function*/

/* Closure is bound to its own surrounding variable */
func appendStr() func(a string) string {
	t := "Hello"
	c := func (b string) string {
		t = t + " " + b
		return t
	}
	return c
}
/* End of closure is bound to its own surrounding variable */

/* Real use case of anonymous function. */
type student struct {
	firstName, lastName, grade, country string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
/* End of real use case of anonymous function. */

func main(){

	/*
		Anonymous Functions.
		Assigning functions to a variable.
	*/

	a := func () {
		fmt.Println("Hello first class functions.")
	}
	a()
	fmt.Printf("%T\n", a) // func()

	/*
		In the above program we have assigned a function to a variable a. This is the syntax of 
		assigning a function to a variable.

		The function assigned to a does not have a name. These kinds of functions are called 
		anonymous functions since they do not have a name.

		The only way to call this function is using the variable a. 
	*/

	/*
		Its also possible to call the Anonymous function without assigning it to a variable.
	*/

	func () {
		fmt.Println("Calling the anonymous function without assigning it to a variable")
	}()

	/*
		It is also possible to pass arguments to the anonymous function just like any other function.
	*/
	func (n string){
		fmt.Println("Passing arguments to an anonymous function: ", n)
	}("Krishna")

	/*
		User defined function types. 
		Just like we define our own struct types, it is also possible to define our own function
		types.
		
		type add func(a int, b int) int
	*/
	var b add = func(a int, b int) int {
		return a + b
	}
	s := b(1,2)
	fmt.Println("User defined function types:", s)

	/*
		Higher order functions.

		> Takes one or more functions as arguments
		> returns a function as its result
	*/

	// Passing functions as arguments to other functions.
	fmt.Println("Higher Order Functions")
	fmt.Println("Passing Functions as arguments to other functions")
	f := func(a int, b int) int {
		return a + b
	}
	simple(f)

	/*
		In the above example we define a function simple which takes a function as an argument 
		which accepts 2 int arguments and returns a int as a parameter. 

		Next we create an anonymous function and assign it to a variable f whose signature matches 
		the parameter of the function simple.

		We call simple and pass f as an argument.
	*/

	// Returning functions from other functions.
	// Lets rewrite the above simple function as simple2
	fmt.Println("Returning functions from other functions.")
	z := simple2()
	fmt.Println(z(1,2))

	/*
		Closures.
		Closures are a special case of anonymous functions.
		Closures are Anonymous functions which access the variables defined outside the body of the
		function.
	*/
	fmt.Println("Closures")
	aa := 5
	func(){
		fmt.Println("aa = ", aa)
	}()

	/*
		In the above program, the anonymous function access the variable aa which is outside its body.
		Hence this anonymous function is a Closure.

		Every closure is bound to its own surrounding variable.
	*/
	fmt.Println("closure is bound to its own surrounding variable")
	ww := appendStr()
	dd := appendStr()

	fmt.Println(ww("World"))
	fmt.Println(dd("Everyone"))

	fmt.Println(ww("Gopher"))
	fmt.Println(dd("!"))

	/*
		In the above program the function appendStr in line no. 33 is a closure. This closure is 
		bound to the variable t.

		The variables ww and dd declared in line no. 144 & 145 are closures and they are bound to their
		own value of t.

		we call ww with the parameter World. Now the value of ww's version of t becomes Hello World.

		Then we call dd with the parameter Everyone. Now the value of dd's version of t becomes Hello Everyone.
		Since dd is bound to its own variable t. dd's version of t has a initial value of Hello again.
	*/

	/*
		Practical use of first class functions.

		lets create a program which filters a slice of students based on some criteria.
	*/
	s1 := student{"Krishna", "Murugan", "A", "India"}
	s2 := student{"Kathir", "Murugan", "B", "India"}
	students := []student{s1, s2}
	f1 := filter(students, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f1)
}