package main

import "fmt"

type rect struct {
	// inline more than one
	// variable of the same type
	width, height int
}

/**
Preferred to use point to avoid copying method calls
 */
func (r *rect) area () int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * r.width + 2*r.height
}

func main(){
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
