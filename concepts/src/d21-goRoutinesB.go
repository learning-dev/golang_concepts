package main

import (
    "fmt"
    "sync"
)

var waitGrp = sync.WaitGroup{}

func main() {
    
    fmt.Println("Start of main func")
    
    // specify that there is one go routine to wait for.
    waitGrp.Add(1)

    go printSum()

    //Wait for the go routine to execute.
    waitGrp.Wait()

    fmt.Println("End of main")
    
}

func printSum() {
    
    fmt.Println("Printing from the function")

    sum := 0

    for i := 1; i <= 10; i++ {
                sum += i
    }
    
    fmt.Printf("Sum of the first 10 numbers is %d\n", sum)

    //Inform the waitgroup that we are done executing 
    waitGrp.Done()

}