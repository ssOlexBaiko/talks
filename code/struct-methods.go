package main

import (
	"fmt"
)
/START OMIT/
type Animal struct {
    Name string
    mean bool
    Age  int
}
func (a Animal) printName() {
    fmt.Println(a.Name)
}
func (a *Animal) setAge(n int) {
    a.Age = n
}
func main() {
	cat := Animal{Name:"Buddy", Age:1}
	cat.printName()
	cat.setAge(32)
	fmt.Println(cat.Age)
}
/END OMIT/
