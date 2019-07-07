package main


import (
    "fmt"
    "unsafe"
)

func main() {
    
    var length float32 = 3.20
    var breadth float64 = 8.45
    
    area := 0.0

    fmt.Println("The length of the rectangle is ", length)
    fmt.Println("The breadth of the rectangle is ", breadth)
    fmt.Printf("The type of variable \"length\" is %T\n", length)
    fmt.Printf("The type of variable \"breadth\" is %T\n", breadth)
    fmt.Printf("The default type of the float variable is %T\n", area)
    
    //show it without the float64 conversion and then with the conversion
    //var area = length * breadth
    area = float64(length) * breadth

    //fmt.Println("The area of the rectangle is ", area)
    fmt.Printf("The area of the rectangle is %.2f", area)


    var minAgeToWork int = 14

    // represted in USD
    var minWage int32 = 1160

    var retirementAge int8 = 62
    
    /* The size of the generic int type is platform dependent.
     It is 32 bits wide on a 32-bit system and 64-bits wide on 
     a 64-bit system
    */

    //if size is 8 bytes then, it's 64 bits system
    fmt.Printf("\nDefault size of int is %d bytes\n\n", unsafe.Sizeof(minAgeToWork))


    fmt.Printf("Minimum age to work in United States is %d", minAgeToWork)
    fmt.Printf("\nMinimum wage sums to %d Dollars per month", minWage)
    fmt.Printf("\nAverage Retirement Age is %d", retirementAge)
    

    // Go doesn't have char type 
    // There are byte and rune 
    // byte is for ASCII characters
    // rune is for unicode characers
    var copyright = '©'
    //byte is an alias of uint8 
    //rune is an alias of int32
    var CharacterNotPresent byte = 'Q'

    fmt.Printf("\nDefault variable type of character in go is  %T \n\n", copyright)
    fmt.Printf("Copyright law is Protected under %c copyright Act of 1976. \n", copyright) 
    fmt.Printf("Fun Fact: Letter '%c' is not present in any of the U.S state names.\n", CharacterNotPresent)
    fmt.Printf("Type of letter 'Q' variable is %T \n", CharacterNotPresent)


    var x = 5 + 7i

    
        var y = complex(17, 2.3)

        fmt.Printf("Type: %T\n", x)
        fmt.Printf("Complex number:%G\n", x)
        fmt.Printf("Another Complex Num:%G\n", y)


        var a = 7 + 3i
        var b = 8 + 2i

        var sum = a + b
        var diff = a - b
        var product = a * b
        var div = a / b

        fmt.Printf("Sum: %G, Diff: %G, Product : %G,", 
                    sum, diff, product)
                    
        fmt.Printf(" Division: %.2f\n", div)


}