package main

import (
    "fmt"
)

func Sum(a int, b int) int {
    return a + b
}

func mapFunc[A any, B any](a []A, f func(A) B) []B {
    n := make([]B, len(a), cap(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func main(){
    names := []string{"one", "two", "three"}
    mappedNames := mapFunc(names, func(n string) string {
        return n + "!" 
    }) 
    fmt.Println(mappedNames)
}
