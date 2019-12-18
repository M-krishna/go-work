package main

import "fmt"

/*
	Implementing Multiple Interfaces.
	A Type can implement more than one interface.
*/

type SalaryCalculator interface{
	calculateSalary()
}

type LeaveCalculator interface{
	calculateLeave()
}

type Employee struct{
	firstName, lastName string
	basicPay, pf int
	totalLeaves, leavesTaken int
}

func (e Employee) calculateSalary() {
	fmt.Printf("%s %s has salary %d\n", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) calculateLeave() {
	fmt.Printf("Leave Balance %d\n", (e.totalLeaves - e.leavesTaken))
}

func main() {

	emp := Employee{
		firstName: "Krishna",
		lastName: "Murugan",
		basicPay: 20000,
		pf: 1800,
		totalLeaves: 12,
		leavesTaken: 6,
	}
	var s SalaryCalculator
	s = emp
	s.calculateSalary()
	var l LeaveCalculator
	l = emp
	l.calculateLeave()

	/*
		The program above has 2 interfaces. SalaryCalculator and LeaveCalculator.
		The Employee struct implements both the interface SalaryCalculator and LeaveCalculator 
		by implementing calculateSalary() method of SalaryCalculator and calculateLeave() method 
		of LeaveCalculator.

		Now Employee implements both SalaryCalculator and LeaveCalculator.
		Because of that line no. 43 and line no. 46 are possible.
	*/

} 