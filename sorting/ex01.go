package main

import (
    "fmt"
    "sort"
)

type byLength []string

//implement sort.Interface - Len, Less, and Swap
func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
	// Sort is inplace.
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}
