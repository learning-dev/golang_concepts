package main

import (
    "fmt"
    "time"

)

var entryCh = make(chan string, 50)
//signal channel 
var endCh = make(chan struct{})

func main() {
    go statusUpdate()

    entryCh <- "The Process has started."
    
    time.Sleep(2* time.Second)

    entryCh <- "Fetching Data..."
    
    time.Sleep(2* time.Second)
    
    entryCh <- "Data record Found."

    time.Sleep(10* time.Microsecond)

    endCh <- struct{}{}

}


func statusUpdate(){

    for {
            select {

            case entry := <- entryCh:

                fmt.Printf("%s: %s\n", time.Now().UTC(), entry )
            
            case  <- endCh:
                    break
            
            }
    }
}