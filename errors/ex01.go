package main

// error handling
// read: https://blog.golang.org/error-handling-and-go

import (
    "errors"
    "fmt"
)

// error from errors package
func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }

    return arg + 3, nil
}

// custom error type
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d: %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {

    for _, v := range []int{7, 42} {
        if r, e := f1(v); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }

    for _, v := range []int{7, 42} {
        if r, e := f2(v); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

	// note the "type assertion" syntax in the follow
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
