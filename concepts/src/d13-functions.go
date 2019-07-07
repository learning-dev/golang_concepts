package main

import (
    "fmt"
    "math"
)

func main() {
    
    firstNum := 9.0

    secondNum := 15.0

    sumofSquares := squareSum(firstNum, secondNum)

    fmt.Printf("The sum of squares of numbers %.0f and %.0f  is %.0f\n", 
                firstNum, secondNum, sumofSquares) 

    avg := average(3.0, 56.0, 9.0)

    fmt.Printf("The average of numbers %.0f, %.0f and %.0f is %.2f\n", 
            3.0, 56.0, 9.0, avg)

    a := []int {25, 41, 53, 69, 81}


    varAvg := VariadicAverage(a...)

    fmt.Printf("The average from variadic function is %d\n", varAvg)

    mileage, changePercentage := milesPerGallon(126.8, 7.28, 23)

    fmt.Printf("The current mileage of your car is %.2f ", mileage)

    if(changePercentage < 0){
        fmt.Printf("There is %.2f%% decrease in from your previous mileage\n", changePercentage)
    } else{
        fmt.Printf("There is %.2f%% increase in from your previous mileage\n", changePercentage)
    }

}

//typical function 
func squareSum(num1, num2 float64) float64{
    
    sum := math.Pow(num1, 2) + math.Pow(num2, 2)

    return sum

}

//named return 
func average(x float64, y int, z float64)  (avg float64) {

    avg = (x + float64(y) + z) / 3 

    return 

}


// variadic Parameters with named return. 
func VariadicAverage(numbers ...int) (average int) {

    for _, num := range numbers {
        average += num
    }
    
    average = average / len(numbers)

    return 
    
}

// return mutitple variables 
func milesPerGallon(miles, fuelSpent, previousMileage float64) ( float64,  float64){
    
    mileage := miles / fuelSpent 

    changePercentage := ( mileage - previousMileage )/ previousMileage

    changePercentage *= 100

    return mileage, changePercentage

}











