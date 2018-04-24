package main

import (
	"fmt"
)
//START OMIT
type xStruct struct{ str string }
func (x xStruct) String() string { return x.str }
func (xStruct) Read(p []byte) (n int, err error) { return 0, nil }

func main() {
	fmt.Println(test(xStruct{"HW!"}))
	fmt.Println(testEmptyInterface("foo"))
	fmt.Println(testEmptyInterface(42))
}

func test(v fmt.Stringer) string {
	//return v.Read() <-- Read() could not be called
	return v.String()
}

func testEmptyInterface(v interface{}) string {
	switch v.(type) {
	case string:
		return "String!"
	case int:
		return "Integer!"
	}
	return "I don't know =("
}
//END OMIT