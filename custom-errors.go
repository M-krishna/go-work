package main

import (
	"fmt"
	//"errors"
	"math"
)

// struct to get more info. from the error.
type areaError struct{
	err string
	radius float64
}

// implementing the error interface.
func(e *areaError) Error() string {
	return fmt.Sprintf("radius %.2f : %s", e.radius, e.err)
}

// to calculate the area of the circle
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		// return 0, errors.New("Area calculation failed, radius cannot be less than zero.")
		//return 0, fmt.Errorf("Area calculation failed, radius %.2f is less than zero", radius)
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func main()  {

	
	/*
		Creating custom error using the "New" function. It is the simplest way to create a custom error.

		This is how the "New" function is implemented in the errors package.

		package errors

		// New returns an error that formats as the given text.
		func New(text string) error {
			return &errorString{text}
		}

		// errorString is a trivial implementation of error.
		type errorString struct{
			s string
		}

		func (e *errorString) Error() string {
			return e.s
		}

		We'll create a program that calculates the area of the circle and will return an error if the
		radius is negative.
	*/

	radius := -20.1
	/*area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The area of circle is %.2f\n", area)*/

	/*
		Adding more information to the error using Errorf.
		Lets modify the above circleArea function by adding the Errorf function from the fmt package.
	*/

	/*
		Providing more information about the error using struct type and fields.

		It is also possible to use struct types which implement the error interface as errors.
		
		In our above example if we want to access the radius which caused the error the only way
		is to parse the error description (Area calculation failed, radius -20.00 is less than zero).

		This is not the proper way to do this since if the description changes our code will break.

		we'll use the strategy followed by the standard library and use struct fields to provide 
		access to the radius which caused the error.

		we will create a struct type that implements the error interface and use its fields to 
		provide more information about the error.

		The first step to create a struct type to represent the error. The naming convention for error
		types is that the name should end with the text Error. So lets name our struct as areaError.

		The next step is to implement the error interface.

		Now lets rewrite the above circleArea function.

		Now lets rewrite the snippet in the main function.
	*/
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok{
			fmt.Printf("Radius %.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of the circle %.2f", area) 

	/*
		Providing more information about the error using methods on struct types.
	*/
}