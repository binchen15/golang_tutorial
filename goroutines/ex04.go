package main

import ("fmt"; "time")

func main() {

	c := make(chan bool, 1)

	go func(done chan bool) {
		fmt.Println("start working...")
		time.Sleep(time.Second)
		fmt.Println("done working")
		c <- true
	}(c)

	<-c // this is used for synchronization
}
