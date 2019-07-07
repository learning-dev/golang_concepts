package main

import (
    "fmt"
    "sync"
)

// working with unbuffered channels 
var waitGrp = sync.WaitGroup{}

func main() {

    channel := make(chan int)
    
    go func() {
            i := <- ch
            fmt.Println(i)
            waitGrp.Done()
    }()

    for j:= 0; j<5; j++{
        
        waitGrp.Add(2)

        go func() {
            ch <- 42
            waitGrp.Done()
        }()
    }
    waitGrp.Wait()
}