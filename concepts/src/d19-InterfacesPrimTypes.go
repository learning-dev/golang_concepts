package main

import (
    "fmt"
)

type Calculate interface{
    bodyMassIndex()
}

//you notice that you don't have to specify the "implements" keyword
type Metrics float64

func (weight Metrics) bodyMassIndex(height Metrics) Metrics{

    bmi := weight/(height * height)

    return bmi

}

func main() {
    
    var weight Metrics = 58.2
    //height := 167.0 
    var height Metrics = 1.675

    fmt.Printf("BMI is %.2f\n", weight.bodyMassIndex(height))

    fmt.Printf("%v  %T\n", weight, weight)


}