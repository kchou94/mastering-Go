package main

import "fmt"

func ats(a1, a2 []string) []string {
	return append(a1, a2...)
}

func main() {
	a1 := [...]string{"a1", "a2", "a3"}
	a2 := [...]string{"b1", "b2", "b3"}
	ats := ats(a1[:], a2[:])
	fmt.Println(ats)
}
