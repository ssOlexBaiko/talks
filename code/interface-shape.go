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
   side  uint
}
func (s *Square) GetPerimeter() uint { return s.side * 4 }
func (s *Square) GetArea() uint {
    return s.side * s.side
}
type Rectangle struct {
   width  uint
   height uint
}
func (r *Rectangle) GetPerimeter() uint {
    return (r.width + r.height) * 2
}
func (r *Rectangle) GetArea() uint {
    return r.width * r.height
}
//END OMIT
//START READ
func main() {
    shapes := []Shape{&Square{side: 2},
                      &Rectangle{width: 3, height: 5}}
    var totalArea uint
    for _, shape := range shapes {
        totalArea += shape.GetArea()
    }
    fmt.Println("Total area: ", totalArea)
}
//START READ
