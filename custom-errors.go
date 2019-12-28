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

/* Providing more information about the error using methods on struct types */

type rectAreaError struct {
	err string
	length float64
	width float64
}

// Now we have the error type, now lets implement the error interface and add a couple of methods on
// the error type to provide more info. about the error.

func (e *rectAreaError) Error() string {
	return e.err
}

func (e *rectAreaError) lengthNegative() bool {
	return e.length < 0
}

func (e *rectAreaError) widthNegative() bool {
	return e.width < 0
}

// The next step is to write the area calculation function.
func calculateRectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "Length is less than zero"
	}
	if width < 0 {
		if err == ""{
			err = "Width is less than zero"
		} else {
			err += ", Width is less than zero"
		}
	}
	if err != "" {
		return 0, &rectAreaError{err, length, width}
	}
	return length * width, nil
}

/* End*/

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

	radius := 20.1
	area1, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The area of circle is %.2f\n", area1)

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
	area2, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok{
			fmt.Printf("Radius %.2f is less than zero\n", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of the circle %.2f\n", area2) 

	/*
		Providing more information about the error using methods on struct types.

		Now we'll calculate the area of the rectangle.

		The first step is to create a struct to represent the error.
	*/
	length, width := 20.4, 19.8
	rectArea, err := calculateRectArea(length, width)
	if err != nil {
		if err, ok := err.(*rectAreaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: %.2f is less than zero\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: %.2f is less than zero\n", err.width)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", rectArea)
}