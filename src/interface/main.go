package main

import (
    "fmt"
    "math"
)

func main() {
    
    var s1 Shape = Rectangle{4.0, 6.0}
    fmt.Printf("Shape Type = %T, Shape Value = %v\n", s1, s1)
    fmt.Printf("Area = %f, Perimeter = %f\n", s1.Area(), s1.Perimeter())
}