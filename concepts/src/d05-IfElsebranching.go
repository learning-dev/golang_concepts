/* 
Run the program once. Then, import random package 
generate the random number 1 to 10. 
Show the initiliazation in the if case statement
*/
package main

import (
    "fmt"
    "math/rand"
    "time"
)


func main() {
    
    var rating int; 

    //rand.Seed(time.Now().UTC().UnixNano())
    fmt.Println("Please rate the service provided on the scale of 10:")
    fmt.Scan(&rating)

    //if rating = rand.Intn(10) + 1; rating >8 {
   if(rating > 8 ) {
        fmt.Println("We are Happy that you enjoyed your time here!",
            "Please come back again")
    } else if(rating > 6 && rating <= 8){
        fmt.Println("We are glad that you liked the service.",
            "Please tell us how we can improve it.")
    } else if(rating >= 4 && rating <= 6){
        fmt.Println("Please tell us how we can improve.")
    } else{
        fmt.Println("We are sorry that You had a terrible experience.",
         "Please tell us what went wrong.")
    }

}