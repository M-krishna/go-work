package main

import "fmt"

/*
	Type Switch.
	A Type Switch is used to compare the concrete type of an interface against multiple types in
	various case statements.
	It is similar to switch case. The only difference being the cases specify types and not the values.

	The syntax for Type Switch is similar to Type Assertion.
	In the syntax i.(T) for Type Assertion, the t should be replaced by the keyword type for Type Switch.
*/

func findType(i interface{}){
	switch i.(type){
		case string:
			fmt.Printf("I am a string with value %s\n", i.(string))
		case int:
			fmt.Printf("I am a int with value %d\n", i.(int))
		default:
			fmt.Println("Unknown Type")
	}
}

/*
	Comapring Type to an interface.
*/
type Describer interface{
	Describe()
}

type Person struct{
	name string
	age int
}

func (p Person) Describe() {
	fmt.Printf("%s is of age %d\n", p.name, p.age)
}

func findType2(i interface{}){
	switch v:= i.(type) {
		case Describer:
			v.Describe()
		default:
			fmt.Println("unknown type")
	}
}

func main(){
	findType("krishna")
	findType(2)
	findType(3.14)

	/*
		It is also possible to compare a type to an interface.
		If we have a type and that type implements an interface, it is possible to compare this type
		with the interface it implements.
	*/
	fmt.Println("Comparing type to an interface")
	p := Person{"krishna", 22}
	findType2(p)
	findType2(1)
}