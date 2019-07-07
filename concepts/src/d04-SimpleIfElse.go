package main

import "fmt"

func main() {

    const minAgeToDrive = 16

    var enteredAge int;

    fmt.Println("Enter your age:")
    fmt.Scan(&enteredAge)

    // take the opening braces to next line and compile - show the error 
    /*if (enteredAge >= minAgeToDrive)
        {
     */
    // you can't take the else to "next line "   
    if(enteredAge >= minAgeToDrive){
        fmt.Println("You are allowed to drive.")
        
    } 
    else {
        waitPeriod := minAgeToDrive - enteredAge
        fmt.Println("You are NOT allowed to drive.")
        fmt.Printf("You should wait for another %d years", waitPeriod)
    }


}

