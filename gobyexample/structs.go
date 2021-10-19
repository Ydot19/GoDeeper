package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main(){
	fmt.Println(person{"Bob", 20})

	// indicate by the properties
	fmt.Println(person{name: "Alice", age: 30})

	// leaving a property out self assigns to its zero
	// for that type
	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("John"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s  // pointer
	fmt.Println(sp.age)

	sp.age = 51
	// point back to the same struct
	fmt.Println(sp.age)
	fmt.Println(s.age)
}
