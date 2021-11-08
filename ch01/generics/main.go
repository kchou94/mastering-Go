package main

import "fmt"

func Print[T any](s []T) {
	for _, v := range s {
		fmt.println(v, " ")
	}
	fmt.Println()
}

func main() {
	Ints := []int{1, 2, 3, 4, 5}
	Strings := []string{"a", "b", "c", "d", "e"}
	Print(Ints)
	Print(Strings)
}
