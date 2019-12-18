package main

import "fmt"

/*
	Embedding Interfaces.

	Although Go does not offer inheritance. It is possible to create new interface by embedding
	other interfaces.
*/

type SalaryCalculator interface{
	DisplaySalary()
}

type LeaveCalculator interface{
	CalculateLeavesLeft() int
}

type EmployeeOperations interface{
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}

func (e Employee) DisplaySalary() {  
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {  
    return e.totalLeaves - e.leavesTaken
}

func main() {  
    e := Employee {
        firstName: "Krishna",
        lastName: "Murugan",
        basicPay: 20000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var empOp EmployeeOperations = e
    empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
	
	/*
		EmployeeOperations interface is created by embedding SalaryCalculator and LeaveCalculator interface.

		Any type is said to implement EmployeeOperations interface if it provides method definitions
		for the methods present in both SalaryCalculator and LeaveCalculator.

		The Employee struct implements EmployeeOperations interface since it provides definitions for
		both DisplaySalary and CalculateLeavesLeft.
	*/
}