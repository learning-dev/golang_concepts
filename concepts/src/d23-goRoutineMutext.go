/*
Introduce Mutex in this program and then lock the variable 
before update i.e Readlock and RWrite lock. ( Inside the methods)
We see that locking the variables doesn't give proper output

What we do is we lock the variable before calling the "go routine" 
then, we unlock the variable inside go routine after we are done

Finally, we see a proper output. There is one drawback as we are
trying to syncronize the routines here. We are losing in terms of 
concurrency and speed. 
Next, demo we introduce "Channels" to communicate between go routines

mutex is mutually exclusive that allows Lock and unlock 
to keep variables exclusive when locked, other variables have to wait 
till it is unlocked. RLock is where anybody can read but not write 
RWlock is where you can't both Read and Write till you have unlocked the variable

*/

package main

import (
    "fmt"
    "runtime"
    "sync"
)


var wg = sync.WaitGroup{}
var counter = 1 
var mutex = sync.RWMutex{}

func main() {
    
    runtime.GOMAXPROCS(4)
    fmt.Println("Number of cores on machine", runtime.GOMAXPROCS(-1))
    
    for i := 0; i < 10; i++{
        
        wg.Add(2)

        //first run without uncommenting - show the output 
        // Then uncomment these lines and 
        // comment the locks inside the methods - show the proper output

        //mutex.RLock()
        go fact(counter)
        
        //mutex.Lock()
        go increment()

    }
    wg.Wait()


}

func fact(num int ) {
    
   mutex.RLock()
   fact := 1

    for i := 2; i <= num; i++ {
        fact *= i
    }

    fmt.Printf("Number: %d  fact: %d  \n", num, fact)
   
    mutex.RUnlock()

    wg.Done()
}

func increment() {
       
        mutex.Lock()
        counter ++
        mutex.Unlock()
        wg.Done()
    
}
