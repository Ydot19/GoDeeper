package factories

import "fmt"

type Employee struct {
	name, Position string
	AnnualIncome int
}

// functional approach to factory generation

/*
NewEmployeeFactory return a function that is a factory give a position and annualIncome

param: position

param: annualIncome
*/
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

type EmployeeFactory struct {
	Position string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactorNonFunctional(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{Position: position, AnnualIncome: annualIncome}
}

func FactoryGeneration(){
	fmt.Println("/* Factory Design Pattern: Factory Generation - Functional */")
	developerFactory := NewEmployeeFactory("Developer", 60_000)
	managerFactory := NewEmployeeFactory("Manager", 80_000)

	developer := developerFactory("Yadi Abdalhalim")
	manager := managerFactory("James Read")

	fmt.Printf("%s \tTitle: %s \tSalary: %d\n", developer.name, developer.Position, developer.AnnualIncome)
	fmt.Printf("%s \tTitle: %s \tSalary: %d\n", manager.name, manager.Position, manager.AnnualIncome)

	fmt.Println("/* Factory Design Pattern: Factory Generation - Idiomatic */")

	bossFactory := NewEmployeeFactorNonFunctional("CEO", 100_000)
	bossFactory.AnnualIncome = 120_000  // can change, with the functional approach we could not change the factory
	boss := bossFactory.Create("The TechLead")
	fmt.Printf("%s \tTitle: %s \tSalary: %d\n", boss.name, boss.Position, boss.AnnualIncome)
}


