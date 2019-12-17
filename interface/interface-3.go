package main

import "fmt"

/*
	Practical example of using an interface.

	Simple Program to calculate the total expenses of company.
*/

type SalaryCalculator interface{
	calculateSalary() int
}

type Permanent struct{
	empname string
	basicpay int
	pf int
}

type Contract struct{
	empname string
	basicpay int
}

func (p Permanent) calculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) calculateSalary() int {
	return c.basicpay
}

func totalExpenses(t []SalaryCalculator){
	//fmt.Println(t)
	expenses := 0
	for _, v := range t{
		//fmt.Println(v)
		expenses = expenses + v.calculateSalary()
	}
	fmt.Printf("Total expenses per month %d\n", expenses)
}


func main() {
	permanent_emp1 := Permanent{"krishna", 20000, 1800}
	permanent_emp2 := Permanent{"Prasanna", 30000, 1800}
	contract_emp1 := Contract{"John", 20000}

	totalExpenses([]SalaryCalculator{permanent_emp1, permanent_emp2, contract_emp1})
}