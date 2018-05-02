package main

import (
	"fmt"
)
//START OMIT
type Animal struct {
    Name string
    mean bool
}
type Cat struct {
    Basics Animal
    MeowStrength int
}
type Dog struct {
    Animal
    BarkStrength int
}
func main() {
	var cat Cat
	var dog Dog
	fmt.Println(cat.Basics.Name)
	fmt.Println(dog.Name)
}
//END OMIT
