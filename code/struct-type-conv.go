package main

import (
	"fmt"
	"reflect"
)

func main() {
	type (
		a struct {
			s string
			a []int
		}
		b struct {
			s string
			a []int
		}
	)

	aV := a{"a", []int{123}}
	bV := b(aV) // Will work while structs have same fields

	// fmt.Println(aV == a(bV)) // <-- Would not compile, because of slice
	fmt.Println(reflect.DeepEqual(aV, bV))    // false
	fmt.Println(reflect.DeepEqual(aV, a(bV))) // true
}
