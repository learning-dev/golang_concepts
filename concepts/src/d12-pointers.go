package main


import (
    "fmt"
)

func main() {
    
    var num int = 10;

    var p *int = &num 

    //another way of declaring pointer 
    newPtr := new(int)     
    
     // default value of *pointer is zero.
    fmt.Println("Default Value of the \"newPtr\":", *newPtr)

    fmt.Println("Address of variable num is", p)
    fmt.Println("Value of variable num is ", *p)

    newPtr = &num

    // Change the value of variable using pointer 
    *newPtr = 24

    fmt.Println("The new value of \"num\" is ", *p)

    // now that you know the two pointer are equal. let' check using ==
    //there is only one operation allowed with pointers i.e equality check
    isEqual := p == newPtr

    fmt.Println("\nAfter assigning P to newPtr.")
    fmt.Println("Are both pointers \"p\" and \"newPtr\" equal?", isEqual)

    //Pointer to Pointer that points to pointers rather than variables
    var pp **int = &p
    fmt.Println("\nAddress of pointer P is ", pp)
    fmt.Println("Value of the address that P is holding - ", **pp)
    fmt.Println("\nAddress of variable num is ", *pp)


}