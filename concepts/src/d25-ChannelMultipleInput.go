/*
From the previous program where we were reversing the number. 
Now, what if we want to check if it is a palindrome? 
Let's passin in another number (original number) in the channel again 
Then you see that program will blow up. 
Next, we introduce buffered channels i.e. adding 3rd args (size) 
in make()
We run the program again then we realize that program runs fine but 
we don't see the second number. 
Next, we add another extraction and print statement from the channel. 
Now, it prints out the numbers correctly 

Then, we introduce the loops
*/



package main


import (
    "fmt"
    "sync"
)

var waitGrp = sync.WaitGroup{}

func main() {

    //1.First run, After program blows up add another third argument
    // Re run the program
    //After re-running the program 
    //you don't get the second number
    
    //2.
    ch := make(chan int, 50)
    //ch := make(chan int)

    value := 546
 
    waitGrp.Add(2)

    go func(channel <- chan int) {
        
        // reversed := <- channel
       
        // fmt.Println("Reversed: ", reversed)
        
        // // Add another statement to take the data from channel
        // originalNum := <- channel 

        // fmt.Println("Palindrome : ", reversed == originalNum)
        
        

        waitGrp.Done() 

    }(ch)
    
    //sending the data into channel 
    go func(cha chan <- int, num int) {
        rev := 0 
        
        origNum := num

        for(num != 0){
            rem := num % 10
            rev = rev * 10 + rem 
            num = num / 10
        } 

        cha <- rev

        // Let's say you want to check for palindrome
        // you want to send the original number as well 
        // program will blow up 
        cha <- origNum
    
        //finally
        // then close the channel 
        close(cha)

        waitGrp.Done()  

    }(ch, value)

    waitGrp.Wait()
    
}



