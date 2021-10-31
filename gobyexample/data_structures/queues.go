package main

import "fmt"

/*
Basic implementation of queues
Also: range over channels
*/

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

}
