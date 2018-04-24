package main

import (
	"fmt"
	"strconv"
)

func main() {
	v1, v2 := 123, 3.14

	// int
	fmt.Printf("Value: %v, Type: %T\n", v1, v1)
	fmt.Printf("Value: %v, Type: %T\n", float32(v1), float32(v1))
	fmt.Printf("Value: %v, Type: %T\n", string(v1), string(v1)) // !!! Wrong !!!
	fmt.Printf("Value: %v, Type: %T\n", strconv.Itoa(v1), strconv.Itoa(v1))

	fmt.Println("--------------")
	// float
	fmt.Printf("Value: %v, Type: %T\n", v2, v2)
	fmt.Printf("Value: %v, Type: %T\n", int(v2), int(v2))
	fmt.Printf("Value: %v, Type: %T\n", strconv.FormatFloat(v2, 0, -1, 32), strconv.FormatFloat(v2, 'E', -1, 32))
}
