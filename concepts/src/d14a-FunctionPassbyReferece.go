package main

import "fmt"

func main() {

    firstNumber := 44
    secondNumber := 38

    fmt.Println("The value of two numbers before swapping are:")
    fmt.Println("Value of firstNum:", firstNumber)
    fmt.Println("Value of secondNum:", secondNumber)

    //call the swap function
    swapNumbers(&firstNumber, &secondNumber)

    fmt.Println("\nInside the main, after swapping the values")
    fmt.Println("Value of firstNum:", firstNumber)
    fmt.Println("Value of secondNum:", secondNumber)

}
func swapNumbers(firstNum, secondNum *int ) {
     
    temp := *firstNum
    *firstNum = *secondNum
    *secondNum = temp

    //*firstNum, *secondNum = *secondNum, *firstNum 

    fmt.Println("\nInside the function, after swapping the values")
    fmt.Println("Value of firstNum:", *firstNum)
    fmt.Printf("Value of secondNum:%d\n", *secondNum)

}