package main

import (
	"fmt"
	"reflect"
)

/*
	Reflection is the ability of the program to inspect its variables and values at run time and find
	their type.

	what is the need to inspect a variable and find its type?
	why do we need to find its type at runtime when each and every variable in our program is defined by us
	and we know its type at compile type itself. well this is true most of the times, but not always.
*/

type order struct {
	ordId int
	customerId int
}

type employee struct {
	name string
	id int
	address string
	salary int
	country string
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func genericCreateQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind() // can be both t & v.
	fmt.Println("Type", t)
	fmt.Println("value", v)
	fmt.Println("Kind", k)
}

// NumField() and Field()
func gCreateQuery(q interface{}){
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields: ", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field: %d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

// creating our full query generator.
func createFullQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
				case reflect.Int:
					if i == 0 {
						query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
					} else {
						query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
					}
				case reflect.String:
					if i == 0 {
						query = fmt.Sprintf("%s \"%s\"", query, v.Field(i).String())
					} else{
						query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
					}
				default:
					fmt.Println("Unsupported Type")
					return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("Unsupported Type")
}

func main()  {
	i := 10
	fmt.Printf("%d %T\n", i, i)

	/*
		In the above program, the type of i is known at compile time.

		Now lets understand the need to know the type of the variable at run time.
		Lets say we want to write a simple function that takes simple struct as an argument and will
		create a SQL insert query using it.
	*/

	o := order{1, 101}
	fmt.Println(o)

	/*
		We need to write a function that takes the above struct o as an argument and return the 
		following SQL query.

		insert into order values(1, 101).

		The function is simple to write.
	*/
	fmt.Println(createQuery(o))

	/*
		Now lets take our query creator to the next level. What if we want to generalize our query
		creator and make it work on any struct.

		lets create another struct(employee).

		Now lets create a new function same as createQuery but this function takes an struct as an 
		argument and returns the respective insert query.

		For Eg. if we pass the below struct

		o := {1, 101}
		It should return "insert into order value(1, 101)"

		if we pass the below struct

		e := {"Krishna", 1, "545, East street", 20000, "India"}
		It should return "insert into employee values("Krishna", 1, "545, East street", 20000, "India")"

		since the genericCreateQuery function should work with any struct, it takes interface{} as 
		an argument.
	*/

	/*
		Reflect package.
		The reflect package implements runtime reflection in Go. The reflect package helps to 
		identify the underlying concrete type and the value of a interface variable. 

		This is exactly what we need. 

		The genericCreateQuery function takes a interface{} as an argument and the query needs to be
		created based on the concrete type and value of the interface{} argument.
	*/

	/*
		reflect.Type and reflect.Value

		The concrete type of interface{} is represented by reflect.Type and the underlying value is 
		represented by reflect.Value.

		There are two functions reflect.TypeOf() and reflect.ValueOf() which return the reflect.Type
		and reflect.Value respectively. These two types are the base to create our query generator.
	*/
	genericCreateQuery(o) // Type main.order, Value {1, 101}

	/*
		The genericCreateQuery() function takes a interface{} as an argument. 

		The function reflect.TypeOf() takes a interface as an argument and returns a reflect.Type which
		contains the concrete type of the interface{} argument passed.

		Similarly the reflect.ValueOf() function takes a interface as an argument and returns a 
		reflect.Value which contains the underlying value of the interface{} argument passed.
	*/

	/*
		reflect.Kind

		The types Kind and Type in the reflection package might seem similar but they have a difference.

		Lets add the Kind to the above genericCreateQuery() function.

		Type represents the actual type of the interface{}, in this case main.order and Kind 
		represents the specific kind of the type, in this case its a struct.
	*/

	/*
		NumField() and Field() Methods.

		The NumField() methods returns the number of fields in a struct and the Field(i int) method
		returns the reflect.Value of the ith field.
	*/
	gCreateQuery(o)

	/*
		Int and string methods.

		The methods Int and String help extract the reflect.Value as an int64 and string respectively.
	*/
	a := 50
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type: %T, value: %v\n", x, x)
	b := "Krishna"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type: %T, value: %v\n", y, y)

	/*
		Now that the reflect concept is quite clear, we'll write the createQuery function completely.
	*/
	createFullQuery(o)
	e := employee{"Krishna", 1, "545, East street", 20000, "India"}
	createFullQuery(e)
	ii := 10
	createFullQuery(ii)
}