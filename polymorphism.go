package main

import "fmt"

/*
	How Polymorphism can be achieved in Go?
	Polymorphism is achieved in Go by using Interfaces.
	Interfaces can be implicitly implemented in Go.
	A type implements an interface if it provides definition for all the methods decalred in the interface.

	A variable of type interface can hold any value which implements the interface.
	This property of interface is used to achieve polymorphism in Go.
*/

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours int
	hourlyRate int
}

// adding new stream of income
type Advertisement struct {
	projectName string
	CPC int
	noOfClicks int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return (tm.noOfHours) * (tm.hourlyRate)
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

// methods for new stream of income
func (ad Advertisement) calculate() int {
	return (ad.CPC * ad.noOfClicks)
}

func (ad Advertisement) source() string {
	return ad.projectName
}

func CalculateNetIncome(i []Income) {
	var netincome int = 0
	for _, income := range i {
		fmt.Printf("Income from %s = %d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Println("Net Income of the organisation = ", netincome)
}

/* Constructing FixedBilling and TimeAndMaterialType */

func ConstructFixedBilling(projectName string, biddedAmount int) FixedBilling {
	return FixedBilling{projectName, biddedAmount}
}

func ConstructTimeAndMaterial(projectName string, noOfHours int, hourlyRate int) TimeAndMaterial {
	return TimeAndMaterial{projectName, noOfHours, hourlyRate}
}

// method for constructing new stream of income (Advertisement)
func ConstructAdvertisement(projectName string, CPC int, noOfClicks int) Advertisement {
	return Advertisement{projectName, CPC, noOfClicks}
}

/* End of Constructing*/

func main() {
	project1 := ConstructFixedBilling("Project 1", 25000)
	project2 := ConstructFixedBilling("Project 2", 12000)
	project3 := ConstructTimeAndMaterial("Project 3", 24, 800)

	// Adding a new stream of income
	project4 := ConstructAdvertisement("Project 4", 25, 100)
	incomestreams := []Income{project1, project2, project3, project4}
	CalculateNetIncome(incomestreams)

	/*
		We have successfully achieved polymorphism using CalculateNetIncome function.

		Now we can also add new stream of income without worrying about the calculation of net income.
	*/
}