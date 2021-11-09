package main

import (
	"fmt"
	"reflect"
)

type T struct {
	F1 int
	F2 string
	F3 float64
}

func main() {
	A := T{1, "hello", 3.14}
	fmt.Println("A:", A)

	r := reflect.ValueOf(&A).Elem()
	fmt.Println("String value:", r.String())

	typeOfA := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		tOfA := typeOfA.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, tOfA, f.Type(), f.Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k == reflect.Int {
			r.Field(i).SetInt(42)
		} else if k == reflect.String {
			r.Field(i).SetString("Changed!")
		} else if k == reflect.Float64 {
			r.Field(i).SetFloat(3.14159)
		}
	}
	fmt.Println("A:", A)
}
