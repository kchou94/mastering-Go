package main

import "fmt"

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1 int
	F2 S1
}

func Print(s interface{}) {
	fmt.Println(s)
}

func main() {
	V1 := S1{1, "a"}
	V2 := S2{-1, V1}
	Print(V1)
	Print(V2)

	Print(123)
	Print("abc")
}
