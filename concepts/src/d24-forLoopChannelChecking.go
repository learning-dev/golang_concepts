/*
From the previous program we have removed the reversing the number 
logic and just passing in the numbers manually. 

We are introducing the for loop to interate through
the channel. All the values are printed but throws up an error at the end.
Then we add a "Close" the channel. to signal that there are no more values 
coming.  


Here, the interesting part is that we are using a infinite loop,
where it contains an "if" statement that check if the channel is closed 
or not. This way it won't have to wait for anymore msgs and no panic is occured. 
There is boolean value which will return "true" when channel is open and 
"false" otherwise.   
*/

package main


import (
    "fmt"
    "sync"
)

var waitGrp = sync.WaitGroup{}

func main() {

    ch := make(chan int)

    waitGrp.Add(2)

    num := 426

    go func(channel <- chan int) {

        // for loop will throw up the error 
        
        for element := range channel {
            fmt.Println(element)
        }
        
        

       /* for {
                //checking if channel is closed or not.
                if num, ok := <- channel; ok {
                    fmt.Println(num)
                    fmt.Println("Channel open:", ok)
                } else{
                    fmt.Println("Channel Open:", ok)
                    break

                }

        }
        */
        waitGrp.Done()
        
    }(ch)
    go func(cha chan <- int, num int) {

        cha <- num
        cha <- 624
       // close(ch)

        waitGrp.Done()
    }(ch, num)
    waitGrp.Wait()
        
}
