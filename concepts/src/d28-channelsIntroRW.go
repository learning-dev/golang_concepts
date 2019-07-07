package main

import (
    "fmt"
    "sync"
)

var waitgrp = sync.WaitGroup{}

func main() {
    // you need to use "make" function to create channel
    //passing in "chan" keyword to create a channel 
    // specify the datatype "int" and it is strongly type 
    
    channel := make(chan int)
        
    value := 546

    waitgrp.Add(2)
    
    go func() {
        
        reversed := <- channel
        
        fmt.Println("Reversed:", reversed)
        
        /* this is to show that a single channel can be
            used as both "Sender" and "Reciever", but it is 
            recommended that you have diff channels for both.
        */  

        //Uncomment this line to show 

        // fmt.Println("Sending Ack to Sender")
        
        // channel <- 1
        
        waitgrp.Done()
    }()

    go func(num int) {
        rev := 0 
        
        for(num != 0){
            rem := num % 10
            rev = rev * 10 + rem 
            num = num / 10
        } 

       
       channel <- rev
       /* 
        if check := <- channel; check == 1{
            fmt.Println("Message recieved successful by reciever")
        } */

        waitgrp.Done()

    }(value)

    waitgrp.Wait()

}