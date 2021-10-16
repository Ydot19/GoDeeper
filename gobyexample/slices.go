package main

import "fmt"

func main(){
	s := make([]string, 3)  // initialize the string with length 3
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("apd:", s)

	// copy arrays
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[2:]
	fmt.Println("sl2:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoDimension := make([][]int, 3)
	for i := 0; i < 3; i ++ {
		innerMaximum := i + 1
		twoDimension[i] = make([]int, innerMaximum)

		for j := 0; j < innerMaximum; j ++ {
			twoDimension[i][j] = i + j
		}
	}

	fmt.Println("2Dimension: ", twoDimension)
}