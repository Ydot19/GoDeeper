package solid

import "fmt"

/**
Do not put too much in the same interface
Basically: Everything doesn't have to go into the kitchen sink
 */

type Document struct {

}

type OldFashionPrinter struct {

}

func (ofp *OldFashionPrinter) print(d Document) {
	// ...
}

/**
Usage of Deprecate
See: https://github.com/golang/go/issues/10909
Parsers will show when a method is deprecated
 */

// Deprecated:
func (ofp *OldFashionPrinter) scan(d Document){
	// ...
}

// ISP: Interface Segregation Principle

type Printer interface {
	Print (d Document)
}

type Scanner interface {
	Scan (d Document)
}

// Initial Implementation

type PhotoCopier struct {}

func (p *PhotoCopier) Scan(d Document){
}

func (p *PhotoCopier) Print(d Document)  {
}

// Multi Functional Device

type MultiFunctionMachine struct {
	Printer
	Scanner
}

// Second Approach - Instance based : Decorator approach

type MultiFunctionalMachine struct {
	p Printer
	s Scanner
}



func InterfaceSegregationPrinciple(){
	fmt.Println("/* SOLID: Interface Segregation Principle */")
}
