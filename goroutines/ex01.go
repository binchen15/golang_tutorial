package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

	// blocking call
    f("direct")

	// asynchronous call 
    go f("goroutine1")

    go f("goroutine2")

	// an anonymous and asynchronous call
    go func(msg string) {
        fmt.Println(msg)
    }("goroutine 3")

	// Wait for goroutines to finish. (better with WaitGroup)
    time.Sleep(time.Second)
    fmt.Println("done")
}
