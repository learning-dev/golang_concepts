package main

import  "fmt"


type Book struct {
    name string
    author string 
    price float64
    numberOfCopies int
    
}

// method
func (b Book ) isInBudget (budget float64) bool{
    return b.price <= budget
} 

func (b Book) isAvailable() (bool, int) {
    
    if (b.numberOfCopies > 0){
        return true, b.numberOfCopies
    }
    return false, 0
}

//update without the pointer and show that 
//value is not update. Then, show the pointer
//func (b Book) updateNumberOfCopies(){
func (b *Book) updateNumberOfCopies(){
    
    b.numberOfCopies --;

}


func main() {
    
    var textBook  = Book { 
        "Programming Pearls", 
        "Jon Bentley",
        5.44, 
        258, 
    }

    checkBudget := textBook.isInBudget(53.48)
    fmt.Println("Is it in budget ", checkBudget)

    available, numCopies := textBook.isAvailable()

    if (available == true){
        
        fmt.Printf("\nThere are %d copies of book available\n", numCopies)
    
    } else {

        fmt.Printf("\nThere are no copies of the book available now.")

    }

    fmt.Println("Lending a book to the Member")

    textBook.updateNumberOfCopies()

    // you realize the number has not been updated. 
    fmt.Println("The updated number of copies of the current book is ",
        textBook.numberOfCopies)

    ps := &textBook

    ps.updateNumberOfCopies()

    fmt.Println("The updated number of copies of the current book is ",
        textBook.numberOfCopies)

}