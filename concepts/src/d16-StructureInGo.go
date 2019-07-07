package main

import (
    "fmt"
)

type Book struct{
    
    // name, author string 
    name string 
    author string 
    price float32 
    noOfpages int
} 


func main() {

    var textBook  = Book { 
        name: "Programming Pearls", 
        author:"Jon Bentley",
        price: 5.44, 
        noOfpages: 256, 
    }

    fmt.Println(textBook)

    //Then 
    var novel = Book{}
    novel.name = "The Adventures of Tom Sawyer"
    novel.author = "Mark Twain"
    novel.price = 7.99
    novel.noOfpages = 248

    // 
    fmt.Printf("\nName of the book: %s\n", novel.name)
    fmt.Printf("Author of the book: %s\n", novel.author)
    fmt.Printf("Price of the book: %.2f\n", novel.price)
    fmt.Printf("Number of pages in the book: %d\n", novel.noOfpages)


    var randomBook = novel

    randomBook.name = "Adventures of Huckleberry Finn"

    // Structs are "value types" and not references to same struct 
    // Diff copies are created when you assign a struct variable to another 
    
    fmt.Println("Name of the book: ", randomBook.name)
    fmt.Println("Name of the book:", novel.name)

    //Struct Pointers that 
    bookPointer := &textBook


    // creating a pointer using new() method.
    var anotherBook = new(Book)

    anotherBook = &novel

    fmt.Printf("The type variable \"anotherBook\" is %T\n", anotherBook)

    fmt.Printf("The type variable \"bookPointer\" is %T\n", bookPointer)
    

    fmt.Println("\n",anotherBook)

    fmt.Println("\nName:", (*bookPointer).price)
    fmt.Println("Author:", anotherBook.author)

    isEqual := *anotherBook == novel
    


    fmt.Println("Are both structs novel and *anotherBook equal?  ", isEqual)



}