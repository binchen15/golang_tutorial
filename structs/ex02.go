package main

// methods. function with struct/*struct as receiver
// pointer dereferened automatically by golang

import "fmt"

type rect struct {
    width, height int
}

// pointer receiver
func (r *rect) area() int {
    return r.width * r.height
}

// value receiver
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

	// . dot notation on struct pointer: OK
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
