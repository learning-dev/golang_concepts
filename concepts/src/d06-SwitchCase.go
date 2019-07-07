package main

import "fmt"

func main(){
    
    month := 7
   
    switch month {
    // initialize the month in the switch statement 
    //switch month := 12; month {
    case 1: 
        fmt.Println("New Year's Day")
        fmt.Println("Martin Luther King, Jr. Day")

    case 2:
        fmt.Println("George Washington’s Birthday")
    
    case 5:
        fmt.Println("Memorial Day")
    
    case 7:
        fmt.Println("Independence Day")
    
    case 9:
        fmt.Println("Labor Day")
    
    case 11:
        fmt.Println("Veterans Day")
        fmt.Println("Thanksgiving Day")

    
    case 12:
        fmt.Println("Christmas Day")

    default:
        fmt.Println("There are no Holidays in this Month as per", 
            "Company Calender")

    }

}