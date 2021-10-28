package solid

import "fmt"

/*
Open for extension, closed for modification
 */

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Specification interface {
	isSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) isSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) isSatisfied(p *Product) bool {
	return p.size == s.size
}

type OpenClosedFilter struct {

}

func (f *OpenClosedFilter) Filter ( products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.isSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}


//func (f *Filter) FilterByColor(products []Product, color Color) []*Product{
//	result := make([]*Product, 0)
//
//	for i, v := range products {
//		if v.color == color {
//			result = append(result, &products[i])
//		}
//	}
//
//	return result
//}

type Product struct {
	name string
	color Color
	size Size
}

func OpenClosedSolid(){
	fmt.Println("/* SOLID: Open Closed Principle */")
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}
	products := []Product{apple, tree, house}
	greenSpec := ColorSpecification{green}
	fitler := OpenClosedFilter{}
	for _, v := range fitler.Filter(products, greenSpec){
		fmt.Printf(" - %s is %s\n", v.name, "green")
	}
}