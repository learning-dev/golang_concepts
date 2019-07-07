package main

import (
    "fmt"
    "time"
)

// Demo is to show what concurrency is Simple 
func main() {

    // first run without the "go routine" then with "go"
    go printString("first")
    //go printString("first")
    // When you add go to the first print statement then, 
    // program doesn't wait for first method to complete, it prints 
    // out the the second method concurrently 
       printString("Second")
}
// you might have noticed that we have not used waitgroup here. 
// Because even the main func which is also a go routine is waiting 
// to print the 1000 statements. 
// When you make the second print statement as go routine too then,
// Main go routine has nothing to do it will exit. In that case, you tell 
// tell the main that there are go routines wait for them by adding wait group


func printString(randomString string) {
    for i := 0; i < 1000; i++ {

        fmt.Println(randomString)
        time.Sleep(time.Millisecond * 400 )
    }

}