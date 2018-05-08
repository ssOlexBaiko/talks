package main

import "errors"

func main() {
	//a := nil compilation error: cannot assign nil without explicit type
	// !!!!! var nil Type // Type must be a pointer, channel, func, interface, map, or slice type
	var nil = errors.New("Never do like this")
	print(nil.Error())
}
