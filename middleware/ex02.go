package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logging() Middleware {

    return func(f http.HandlerFunc) http.HandlerFunc {

        // Define the http.HandlerFunc
        return func(w http.ResponseWriter, r *http.Request) {

            // Do middleware things
            start := time.Now()
            defer func() { log.Println(r.URL.Path, time.Since(start)) }()
            f(w, r)
        }
    }
}

// Method ensures that url can only be requested with a specific method, 
// else returns a 400 Bad Request
func Method(m string) Middleware {

    return func(f http.HandlerFunc) http.HandlerFunc {

        // Define the http.HandlerFunc
        return func(w http.ResponseWriter, r *http.Request) {

            if r.Method != m {
                http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
                return
            }
            f(w, r)
        }
    }
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for _, m := range middlewares {
        f = m(f)
    }
    return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello world")
}

func main() {
    http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
    http.ListenAndServe(":8080", nil)
}

