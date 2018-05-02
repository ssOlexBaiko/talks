package main

import (
	"fmt"
)

func returnStrings() (string, string) {
	return "First", "Second"
}

func returnInts() (i int, j int) {
	i, j = 7, 8
	return
}

func main() {
	str1, str2 := returnStrings()
	int1, int2 := returnInts()
	fmt.Printf("Strings: %s-%s\nInts:%d-%d", str1, str2, int1, int2)
}
