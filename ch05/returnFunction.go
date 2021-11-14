package main

import "fmt"

func funReturn(i int) func(int) int {
	if i < 0 {
		return func(x int) int {
			x = -x
			return x + x
		}
	}

	return func(x int) int {
		return x * x
	}
}

func main() {
	x := 10
	i := funReturn(x)
	j := funReturn(-4)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T %v\n", j, j)
	fmt.Println("j", j, j(-5))

	// Same input parameter but DIFFERENT
	// anonymous functions assigned to i and j
	fmt.Println(i(10))
	fmt.Println(j(10))
}
