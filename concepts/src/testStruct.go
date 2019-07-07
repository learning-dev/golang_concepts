package main

import "fmt"

type Employee struct {
    firstName, lastName string
    salary              int
    fullTime            bool
}

func main() {
    ross := &Employee{
        firstName: "ross",
        lastName:  "Bing",
        salary:    1200,
        fullTime:  true,
    }
    
    fmt.Println("firstName", (*ross).firstName)
}
