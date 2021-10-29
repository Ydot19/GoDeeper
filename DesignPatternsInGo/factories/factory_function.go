package factories

import "fmt"

type Person struct {
	Name string
	Age int
	EyeCount int	// typically people have two eyes
	EligibleToDrive bool
}

// NewPerson Instead of instantiating 2 all the time for eyes, we can use a factory to cover this
func NewPerson(name string, age int) *Person{
	if age < 16 {
		//... do Something
	}
	return &Person{name, age, 2, true}
}

func FactoryFunction(){
	fmt.Println("/* Factory Design Pattern: Basics */")
	person := NewPerson("Yadi Abdalahlim", 26)

	fmt.Printf("name: %s \nage: %d \neligibleToDrive: %t \n\n",person.Name, person.Age, person.EligibleToDrive)
}