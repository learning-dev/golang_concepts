//switch case with no expression, instead condition is inside the switch block 

package main

import "fmt"

func main(){
    
    month := 7
   
    switch  {

    case month == 1: 
        fmt.Println("New Year's Day")
        fmt.Println("Martin Luther King, Jr. Day")

    case month == 2:
        fmt.Println("George Washington’s Birthday")

    case month == 5:
        fmt.Println("Memorial Day")

    case month == 7:
        fmt.Println("Independence Day")

    case month == 9:
        fmt.Println("Labor Day")

    case month == 11:
        fmt.Println("Veterans Day")

    case month == 12:
        fmt.Println("Thanksgiving Day")
        fmt.Println("Christmas Day")

    default:
        fmt.Println("There are no Holidays in this", 
            "Month as per Company Calender")

    }

}