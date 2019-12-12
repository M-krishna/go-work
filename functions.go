package main

import "fmt"

// function declaration
/*func nameOfTheFunction(parametername type) returntype{

}*/

// sample function
func calculateBillPrice(price, no int) int{ // consecutive parameters(args) with same type
	result := price * no
	return result
}

// multiple return values
func rectProps(length, width float64) (float64, float64){
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

// named return values
func rectProps2(length, width float64) (area, perimeter float64, name string){
	area = length * width
	perimeter = (length + width) * 2
	name = "Krishna"
	return
}

// main function
func main(){
	fmt.Println(calculateBillPrice(10, 5)) // calculate function

	// to get the multiple return values of the function
	area, perimeter := rectProps(10, 5)
	fmt.Printf("Area = %v, Perimeter = %v\n", area, perimeter)

	// named return values
	area1, perimeter1, name := rectProps2(10, 20)
	fmt.Println("Named return values")
	fmt.Printf("Area = %v, Perimeter = %v, name = %v\n", area1, perimeter1, name)

	/* 
		usage of "Blank Identifier" 
		it can be used in place of any value of any type
		
		the rectProps2 function returns 3 values
		but we only need area and perimeter and don't want the name
		to discard the name parameter
	*/

	area2, perimeter2, _ := rectProps2(10, 20)
	fmt.Println("Using Blank Identifier")
	fmt.Printf("Area = %v, Perimeter = %v\n", area2, perimeter2)

}