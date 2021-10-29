package factories

import "fmt"

type EmployeePrototype struct {
	Name, Position string
	AnnualIncome int
}

const (
	Developer = iota
	Manager
)

func NewEmployeePrototype(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60_000}
	case Manager:
		return &Employee{"", "Manager", 80_000}
	default:
		panic("unsupported role")
	}
}

func FactoryPrototypes(){
	fmt.Println("/* Factory Design Pattern: Factory Prototype */")
	m := NewEmployeePrototype(Manager)
	m.name = "Yadi Abdalhalim"
	fmt.Println(m)
}