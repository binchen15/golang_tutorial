package main

import "fmt"

// channel direction. type safety mechanism

func ping(c chan<- string, msg string){
	c <- msg
}

func pong(r <-chan string, w chan<- string){
	msg :=  <-r
	w <- msg
}

func main(){
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	ping(c1, "hello")
	pong(c1, c2)
	fmt.Println(<-c2)
}
