package main

import (
	"fmt"
)

func getValue(n int) {
	n +=1
}

func getPointer(n *int) {
	*n +=1
}

func main() {
	val := 5
	getValue(val)
	fmt.Println(val)
	// -------------
	p := &val
	getPointer(p)
	fmt.Println(val)
}
