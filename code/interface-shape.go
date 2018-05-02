package main

import (
	"fmt"
)
//START OMIT
type Shape interface {
    GetPerimeter() int
    GetArea() int
}
type Square struct {
   side  int
}
func (s *Square) GetPerimeter() int { return s.side * 4 }
func (s *Square) GetArea() int {
    return s.side * s.side
}
type Rectangle struct {
   width  int
   height int
}
func (r *Rectangle) GetPerimeter() int {
    return (r.width + r.height) * 2
}
func (r *Rectangle) GetArea() int {
    return r.width * r.height
}
//END OMIT
//START READ
func main() {
    shapes := []Shape{&Square{side: 2},
                      &Rectangle{width: 3, height: 5}}
    var totalArea int
    for _, shape := range shapes {
        totalArea += shape.GetArea()
    }
    fmt.Println("Total area: ", totalArea)
}
//START READ
