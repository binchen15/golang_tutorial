package main

// range over channels

import "fmt"

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)  // close first, still can be read 

    for elem := range queue {
        fmt.Println(elem)
    }
}
