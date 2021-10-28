package solid

import "fmt"

/*
Dependency Inversion Principle
High level module should not depend on Low Level Module
Both should depend on abstractions
 */

type RelationShip int

const (
	Parent RelationShip = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from *Person
	relationship RelationShip
	to *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (rel *Relationships) FindAllChildrenOf(name string) []*Person{
	res := make([]*Person, 0)
	for _, v := range rel.relations {
		if v.relationship == Parent && v.from.name == name {
			res = append(res, v.to)
		}
	}
	return res
}

func (rel *Relationships) AddParentAndChild(parent, child *Person){
	rel.relations = append(rel.relations, Info{parent, Parent, child})
	rel.relations = append(rel.relations, Info{child, Child, parent})
}

// High-Level Module

type Research struct {
	// get list of relationships of parent to child
	// using relationships Relationships breaks Dependency Inversion Principle

	// use a low level method that allows us to browse for parent-child relationships
	browse RelationshipBrowser
}

func (r *Research) Investigate(name string){
	for _, child := range r.browse.FindAllChildrenOf(name) {
		fmt.Printf("%s has a child named %s\n", name, child)
	}
}

func DependencyInversionPrinciple(){
	fmt.Println("/* SOLID: Dependency Inversion Principle */")

	parent := Person{"Yaadata"}
	child1 := Person{"Demascus"}
	child2 := Person{"Haruun"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate("Yaadata")
}
