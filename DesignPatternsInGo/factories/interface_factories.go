package factories

import "fmt"

/**
Showcase interface factories
Good way to deliver encapsulation
 */

type PersonEncapsulation interface {
	SayHello()
}

type person struct {
	name string
	age int
}

type tiredPerson struct {
	name string
	age int
}

func (p *person) SayHello(){
	fmt.Printf("Hi, my name is %s, I am %d years old\n\n", p.name, p.age)
}

func (p *tiredPerson) SayHello(){
	fmt.Printf("My name is %s - currently I am too tired to talk\n\n", p.name)
}

func NewEncapsulatedPerson(name string, age int) PersonEncapsulation {
	if age > 50 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func InterfaceFactories() {
	fmt.Println("/* Factory Design Pattern: Factory Functions */")
	p := NewEncapsulatedPerson("Yadi Abdalhalim", 26)
	p.SayHello()

	tiredP := NewEncapsulatedPerson("Michael Jordan", 51)
	tiredP.SayHello()
}
