package main

import "fmt"

func main(){
	messages := make(chan string)

	go func(){messages <- "ping"}()

	/**
	The <-channel syntax receives a value from the channel.
	Here we’ll receive the "ping" message we sent above and print it out.
	 */
	msg := <-messages
	fmt.Println(msg)
}