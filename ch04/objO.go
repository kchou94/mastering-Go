package main

import "fmt"

type IntA interface {
	foo()
}

type IntB interface {
	bar()
}

type IntC interface {
	IntA
	IntB
}

func processA(s IntA) {
	fmt.Printf("%T\n", s)
}

type a struct {
	XX int
	YY int
}

// Satisfying IntA
func (varC c) foo() {
	fmt.Println("Foo Processing", varC)
}

// Satisfying IntB
func (varC c) bar() {
	fmt.Println("Bar Processing", varC)
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

type compose struct {
	field1 int
	a
}

func (A a) A() {
	fmt.Println("Function A() for A")
}

func (B b) A() {
	fmt.Println("Function A() for B")
}

func main() {
	var iC = c{a{120, 12}, b{"-12", -12}}
	iC.A.A()
	iC.B.A()

	iComp := compose{123, a{456, 789}}
	fmt.Println(iComp.XX, iComp.YY, iComp.field1)

	iC.bar()
	processA(iC)
}
