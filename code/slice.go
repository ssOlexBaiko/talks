package main

import (
	"fmt"
)

func main() {
	var s1 []int
	s2 := []int{1}
	s3 := make([]int, 2, 10)
	fmt.Printf("s1: len = %d, cap = %d, is nil %t\n", len(s1), cap(s1), s1 == nil)
	fmt.Printf("s2: len = %d, cap = %d, is nil %t\n", len(s2), cap(s2), s2 == nil)
	fmt.Printf("s3: len = %d, cap = %d, is nil %t\n", len(s3), cap(s3), s3 == nil)

	for i := 0; i < 2<<10; i++ {
		s1 = append(s1, i)
	}
	fmt.Println("After append:", "len =", len(s1), "cap =", cap(s1))
}
