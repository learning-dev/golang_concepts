/* Run the program show the output 
   Run "Again", you see that there is different ouput
   everytime. Both go routines are competing to complete. 
   There is no syncronization between go routines. 
*/

package main

import (
    "fmt"
    "sync"
)

var wg = sync.WaitGroup{}
var counter = 1

func main() {
    
    for i := 0; i < 10; i++{
        
        wg.Add(2)

        go fact(counter)
        go increment()

    }
    wg.Wait()
}

func fact(n int ) {
    fact := 1

    for i := 2; i <= n; i++ {
        fact *= i
    }

    fmt.Printf("number: %d  fact: %d  \n", n, fact)
    wg.Done()
}

func increment() {
        counter ++
        wg.Done()
    
}