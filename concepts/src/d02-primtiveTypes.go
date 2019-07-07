package main 

import "fmt"


func main() {
    
    
    var firstName string = "James"
    lastName := "Cook"
    var age int = 45

    const height int32 = 164

    var (
        weight float64 = 68.42
        occupation string = "Hotelier"
        married bool = true
    )

    
    fmt.Printf("Name: %s %s, Age: %d,\nHeight: %d, Weight: %f\n ",
                firstName, lastName, age, height, weight)
    

    fmt.Printf("He works as %s \n Is married?: %t\n", occupation, married)

}