package main

import "fmt"

/*
	Implementing Interfaces using Pointer receivers vs Value Receivers.
*/

type Describer interface{
	Describe()
}

type Person struct{
	name string
	age int
}

func (p Person) Describe(){ // implemented using value receiver
	fmt.Printf("%s is of age %d\n", p.name, p.age)
}

type Address struct{
	city, state string
}

func (a *Address) Describe(){ // implemented using pointer receiver
	fmt.Printf("state %s, city %s\n", a.state, a.city)
}

func main(){
	var d1 Describer
	p1 := Person{"Krishna", 22}
	d1 = p1
	d1.Describe()

	p2 := Person{"John", 18}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"chennai", "tamilnadu"}
	// d2 = a throws an error.
	d2 = &a
	d2.Describe()

	/*
		Methods with value receivers accept both pointer and value receivers.
		It is legal to call a value method on anything which is a value or whose value can be dereferenced.
		
		In we uncomment line number 42 we'll get an compilation error saying 
		main.go:42: cannot use a (type Address) as type Describer in assignment: 
		Address does not implement Describer (Describe method has pointer receiver).

		This is because the Describer interface using Address Pointer receiver in line number 26
		and we are trying to assign it to a which is a value type and it has not implemented the 
		Describer interface.

		This will be confusing because that methods with pointer receivers will accept both value 
		and pointer receivers. Then why isn't this working.

		The reason is that it is legal to call a pointer-value method on anything that is a pointer
		or whose address can be taken. The concrete value stored in an interface is not addressable
		and hence it is not possible for the compiler to automatically take the address a. Hence this
		code fails.
	*/
}