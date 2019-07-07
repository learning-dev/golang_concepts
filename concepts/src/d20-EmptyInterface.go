package main

import (
    "fmt"
)

func main() {
    
    var someInterface interface{} = 23

    switch someInterface.(type){
    case int:
        fmt.Println("Given inteface is of type integer")
    
    case string:
        fmt.Println("Given interface is of type string")
    
    default:
        fmt.Printf("Given interface is of  type %T\n", someInterface)

    }

    tri := Triangle{24.1}
    mass := Metrics(2.3)

    describe(tri)
    describe(mass)
 
}

type Shape interface{
    area() float64
    perimeter() float64
} 


type Metrics float64

type Triangle struct {
    side float64
}

// this can take any type of interface as paramter and print it. 
func describe(blankInterface interface{}){

    fmt.Printf("The type of interface %T and value is %v\n",
     blankInterface, blankInterface)

}