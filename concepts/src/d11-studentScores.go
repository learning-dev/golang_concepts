package main

import (
    "fmt"
   "sort"
)

func main() {
    
    var numberOfStudents int;

    var rollNo int;
    var score float64;
    var studentData map[int]float64 = make(map[int]float64)
    var choice int;

    fmt.Println("-----Menu-----")
    fmt.Println("1. Calculate Average Student Scores")
    fmt.Println("2. Find the Highest Scoring Student")
  
    fmt.Printf("Please enter your choice:")
    fmt.Scan(&choice)

   switch choice { 

    case 1:

      fmt.Println("\nEnter the number of students you want to add:")
      fmt.Scan(&numberOfStudents)

      fmt.Println("Start Entering the Roll number and scores of the students:")

      for i := 0; i < numberOfStudents; i++ {
          fmt.Printf("Roll no:")
          fmt.Scan(&rollNo)
          fmt.Printf("Score:")
          fmt.Scanf("%f", &score)
          studentData[rollNo] = score
       }

        
      for i, k := range studentData {
          fmt.Printf("Rollno :%d Score: %.2f\n", i, k)
       }

      avgScore := 0.

      for _, score := range studentData{
            avgScore += score
       }

      avgScore = avgScore / float64(len(studentData))

      fmt.Printf("Average Score of the Entered Student Records is %.2f\n", avgScore)

    case 2:

      fmt.Println("\nEnter the number of students you want to add:")
      fmt.Scan(&numberOfStudents)

      fmt.Println("Start Entering the Roll number and scores of the students:")

      for i := 0; i < numberOfStudents; i++ {
          fmt.Printf("Roll no:")
          fmt.Scan(&rollNo)
          fmt.Printf("Score:")
          fmt.Scanf("%f", &score)
          studentData[rollNo] = score
      }
      
      scores := []float64{}
      
      
      for _, score := range studentData{
          scores = append(scores, score)

      }
      // sort in the descending order.
      sort.Slice(scores, func(i, j int) bool{
        return scores[i] > scores[j]
      })

      highestScore := scores[0]

      for key, value := range studentData{
        if (value == highestScore){
            fmt.Printf("The Highest scoring student is of Roll no: %d and with a score of %.2f.", 
              key, value)
            break
        }
      } 

    default:
      fmt.Println("Invalid Option!")

   }


}

