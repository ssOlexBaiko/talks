package main

import (
	"fmt"
)

func main() {
	var someText string = "日本語"
	fmt.Println(string(someText[0]))
	fmt.Println(string(someText[1]))
	fmt.Println(string(someText[2]))

	fmt.Println("-----------")
	fmt.Println(string(someText[0:3]))
	fmt.Println(string(someText[3:6]))
	fmt.Println(string(someText[6:]))

	fmt.Println("-----------")
	fmt.Println(string([]rune(someText)[0]))
	fmt.Println(string([]rune(someText)[1]))
	fmt.Println(string([]rune(someText)[2]))
}
