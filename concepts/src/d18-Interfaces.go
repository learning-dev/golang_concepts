package main

import (
    "fmt"
)


type Shape interface{
    area() float64
    perimeter() float64
} 

type Square struct {
    side float64
}

type Triangle struct {
    height float64
    base float64
    
}


func (tri Triangle) area() float64{
    return 0.5 * tri.height * tri.base
}
func (sqr Square) area() float64{
    return sqr.side * sqr.side 
}

func (tri Triangle) perimeter() float64{
    return tri.base * 3
}

func (sqr Square) perimeter() float64{
    return sqr.side * 4
}

func main() {
    
    var Tri Shape = Triangle{6.06, 7}
    var sqr Shape = Square{4.5}

    fmt.Printf("The area of the square is %.2f\n", sqr.area())
    fmt.Printf("The area of the Triangle is %.2f\n", Tri.area())
    fmt.Printf("The Perimeter of the Triangle is %.2f\n", Tri.perimeter())
    fmt.Printf("The Perimeter of the square is %.2f\n", sqr.perimeter())

    fmt.Printf("%v  %T\n", Tri, Tri)

}


