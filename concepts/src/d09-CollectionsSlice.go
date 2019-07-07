package main

import (
    "fmt"
)

func main() {
    
    var studentIDs = [3]int{1, 2, 3}

    var StudentScores  = [3]int{} 

    StudentMarks[0] = 78
    StudentMarks[1] = 92
    StudentMarks[2] = 56

    fmt.Println("Numbers of students:", len(studentIDs))

    fmt.Println("Roll numbers of the Students:", studentIDs)
    fmt.Println("Marks of the Students:", StudentMarks)

    avg := 0 

    for i := 0; i < len(StudentMarks); i++ {
        avg += StudentMarks[i] 
     } 
    
    avg  = avg/len(StudentMarks)

    fmt.Println("Average marks of given students is ", avg)


    var numberOfStudents int;
    


    fmt.Println("\nEnter the number of students you want to add:")
    fmt.Scan(&numberOfStudents)

    var moreStudents = make([]int, numberOfStudents )
    
    var marksStudents = []int{}

    fmt.Println(moreStudents)    
    fmt.Printf("Capacity of the Rollno Slice: %d \n", cap(moreStudents))
    fmt.Printf("Length of the Rollno Slice: %d \n", len(moreStudents))
    fmt.Printf("Capacity of the marks Slice: %d\n", cap(marksStudents))
    fmt.Printf("Length of the marks Slice: %d .\n", len(marksStudents))


    fmt.Println("Start Entering the Roll number:")

    var input int;

    for i := 0; i < numberOfStudents; i++ {
         fmt.Scan(&input)
         moreStudents[i] = input
    }
 

    fmt.Println("Enter the marks of the students:")

    
    // You can't add the elements same way you did with moreStudents 
    // as this is an empty slice i.e. underlying array is of length 0 (zero)
    // You need to add elements using append()

    for i := 0; i < numberOfStudents; i++ {
        fmt.Scan(&input)
        marksStudents = append(marksStudents, input)           
    }

    for i,_ := range moreStudents {
        fmt.Printf("Rollno: %d  Marks: %d\n", moreStudents[i], marksStudents[i])

    }

    // add an extra element 

    moreStudents = append(moreStudents, 56)
    marksStudents = append(marksStudents, 89)

    fmt.Println("After adding more elements.")

    for i,_ := range moreStudents {
        fmt.Printf("Rollno: %d  Marks: %d\n", moreStudents[i], marksStudents[i])
    }

    fmt.Printf("New Capacity of the Rollno Slice: %d \n", cap(moreStudents))
    fmt.Printf("New Length of the Rollno Slice: %d \n", len(moreStudents))
    fmt.Printf("New Capacity of the marks Slice: %d\n", cap(marksStudents))
    fmt.Printf("New Length of the marks Slice: %d .\n", len(marksStudents))

}