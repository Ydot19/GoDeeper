package solid

import "fmt"

/**
Deals with base class and inherited classes
 */

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func(r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int){
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	r Rectangle  // width, height
}

func (s *Square) GetSideLength() int{
	return s.r.GetWidth()
}

func (s *Square) SetSideLength(side int) {
	s.r.SetWidth(side)
	s.r.SetHeight(side)
}

func LiskovSubstitutionPrinciple(){
	fmt.Println("/* SOLID: Liskov Substitution Principle */")
	rc := &Rectangle{2, 3}
	fmt.Printf("rectangle:\t width: %d height: %d\n", rc.GetWidth(), rc.GetHeight())

	sq := &Square{}
	sq.SetSideLength(4)
	fmt.Printf("square:\t side: %d\n", sq.GetSideLength())
}

