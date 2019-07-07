package main

import (
    "fmt"
    "sync"
)
var waitGrp = sync.WaitGroup{}

func main() {
    
    var value = 10

    waitGrp.Add(1)
    
    // It will print "20"

    /* Because both function are competing to complete
        Main function (first go routine) completes first
        Then next go routine access the value, finds it as 20
    */
    // Anonymous function which is called and defined at same place

    go func() {

        fmt.Println("Value:", value)
        
        waitGrp.Done()      
    }()


    value = 20 

    waitGrp.Wait()

}

//After you have run the program normally. 

// Then run using :
//go run -race demos/src/goRoutinesRaceTest.go

//Then you move on to fixing this 
// By passing the variables as parameters rather than acessing 
// the variables as function "closures"

/*

go func(num int) {

        fmt.Println("Value:", num)
        
        waitGrp.Done()      
    }(value)

    value = 20 

*/









