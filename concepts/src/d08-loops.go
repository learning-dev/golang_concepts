package main

import "fmt"

func main() {


    // run the program once before adding anything 

    fmt.Println("Your current Waiting period is 5 Minutes")
    fmt.Println("Your current Waiting period is 4 Minutes")
    fmt.Println("Your current Waiting period is 3 Minutes")
    fmt.Println("Your current Waiting period is 2 Minutes")
    fmt.Println("Your current Waiting period is 1 Minutes")

    waitingPeriod := 5


    fmt.Println("\nPrinting using loop.")

    for i:= waitingPeriod; i > 0; i--{
        fmt.Printf("Your Current Waiting Period is %d Minutes\n", i)        
    }


    
    fmt.Println("\nUsing while(for in golang) loop")

    for waitingPeriod >0 {
        fmt.Printf("Your Current Waiting Period is %d Minutes\n", waitingPeriod)
        waitingPeriod = waitingPeriod -1 ;
    }

}