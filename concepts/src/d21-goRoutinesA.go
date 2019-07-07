package main

import (
    "fmt"
    "time"
)


func main() {
    
    //show that go routine didn't run 
    fmt.Println("start of main func")
    
    go printSum()

    fmt.Println("End of main")
    // After adding sleep statement
    /* you realize that main func executes faster 
       than the go routine. 
    */
    time.Sleep(100 * time.Millisecond)

}

func printSum() {
    
    fmt.Println("Printing function")

    sum := 0

    for i := 1; i <= 10; i++ {
                sum += i
    }
    
    fmt.Printf("Sum of the first 10 numbers is %d\n", sum)

}