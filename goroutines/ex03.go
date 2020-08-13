package main

import "fmt"

func main() {

	// buffered channel, size 1.
	messages := make(chan string, 1)

	// no receiver ready, still succeed
	messages <- "hello"
	go func() {
		messages <- "world"
	}()  // this must be run as a goroutine

	fmt.Println( <-messages ) // buffer read. second msg to buffer
	fmt.Println( <-messages )  // buffer read again
}
