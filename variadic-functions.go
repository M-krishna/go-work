package main

import "fmt"

/*
	Functions generally accept fixed number of arguments.
	A Variadic Function is a function that accepts variable number of arguments.
	if the last parameter of the function definition is prefixed by ellipsis ... ,
	then the function can accept any number of arguments to that parameter.

	Only the last parameter of the function is variadic.

	func sample(a int, b ...int){}


*/

func helloWorld(a int, b ...int){
	fmt.Println(a) // 1
	fmt.Printf("The type of b is %T\n", b) // The type of b is []int(slice)
	fmt.Println(b)
}

// example for only the last parameter must be the variadic
/* func lastVariable(b ...int, a int){
	fmt.Println(b)
	fmt.Println(a)
} */

/*
	Creating a variadic function
*/
func find(num int, numbers ...int){
	fmt.Printf("type of numbers is %T\n", numbers)
	found := false
	for i, v := range numbers {
		if v == num {
			fmt.Printf("Number %d found at index %d\n", v, i)
			found = true
		}
	}
	if !found {
		fmt.Println("Not Found")
	}
}

/*
	creating the same find function as above.
	but accepting slice args as the second parameter
*/
func find2(num int, numbers []int){
	fmt.Printf("type of numbers is %T\n", numbers)
	found := false
	for i, v := range numbers {
		if v == num {
			fmt.Printf("Number %d found at index %d\n", v, i)
			found = true
		}
	}
	if !found {
		fmt.Println("Not Found")
	}
}

/*
	change function to check variadic function
*/
func change(s ...string){
	s[0] = "Go"
	s = append(s, "Playground")
	fmt.Println(s)
}


func main(){
	fmt.Println("Variadic Function")
	helloWorld(1, 2, 3, 4)

	// it is also possible to pass zero arguments to a variadic function
	fmt.Println("Not passing arguments to variadic parameter")
	helloWorld(1)

	// making variadic variable as first parameter
	// lastVariable(1,2,3,4,5) // throws an error

	// variadic function (find)
	fmt.Println("User defined variadic function find")
	find(1, 2,3,4,5,1)
	find(10, 1,2,3,4,5)
	find(200)

	/*
		slice arguments vs variadic arguments

		the advantages of using variadic arguments instead of slices
		> there is no need to create slice during each function call.
		> we are creating an empty slice just to satisfy the signature of the find2 function. 
	*/
	fmt.Println("find2 function, which takes slice as an argument.")
	find2(1, []int{1,2,3,4,5})
	find2(100, []int{1,2,3,4})
	find2(0, []int{})

	/*
		passing a slice to a variadic function.
		what if I pass a slice to a variadic function.
	*/
	// find(34, []int{34,26}) // throws an error(cannot use []int literal (type []int) as type int in argument to find)
	/*
		because the variadic function tries to convert the variadic parameters to slice.
		the compiler will try to convert.

		is there any way to pass slice to a variadic function?
		yes we can. there is a syntactic sugar we can use to pass slice to a variadic function.
		we have to suffix the slice with ellipsis ...
		a := []int{1,2,3,4,5}
		find(12, a...)
		lets check
	*/
	a := []int{1,2,3,4,5}
	find(3, a...)


	/*
		Just be sure what you are doing when modifying a slice inside a variadic function.
	*/
	fmt.Println("Checking variadic functions")
	hw := []string{"Hello", "World"}
	change(hw...)
	fmt.Println(hw) // will not affect the slice in the main function.

}