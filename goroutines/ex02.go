// first example of channel

package main

import "fmt"

func main() {

	// make(chan val-type)
    messages := make(chan string)

	// both sender/receiver block until both sides ready
    go func() { messages <- "ping" }()

	// no need of otheer synchronization mechanism
	msg := <-messages

    fmt.Println(msg)
}
