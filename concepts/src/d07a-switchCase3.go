package main

import "fmt"


func main() {
    
    month := 7

    switch month {

    case 1, 3, 5, 7, 8, 10, 12 :
        fmt.Println("There are 31 days in this month")
    
    case 2:
        fmt.Println("There are 28 to 29 days depending on the year")
    
    case 4, 6, 11, 9:
        fmt.Println("There are 30 days in this month")

    default:
        fmt.Println("Invalid Month!")

    }


}