package main

import (
    "fmt"
   "sort"
)

func main() {
  
  var studentData map[int]float64 = make(map[int]float64)
  
  choice := choiceMenu()

  switch choice {

    case 1:
      getInput(studentData)
      calculateAverageScore(studentData)
      
    case 2:

      getInput(studentData)
      getHighestScorer(studentData)
     
    default:
      fmt.Println("Invalid Option!")

   }

}


func choiceMenu() (choice int){

  fmt.Println("-----Menu-----")
  fmt.Println("1. Calculate Average Student Scores")
  fmt.Println("2. Find the Highest Scoring Student")

  fmt.Printf("Please enter your choice:")
  fmt.Scan(&choice)

  return 
}

func calculateAverageScore(studentData map[int]float64 ) (){

   avgScore := 0.

   for _, score := range studentData{
        avgScore += score
   }
   
   avgScore = avgScore / float64(len(studentData))

   fmt.Printf("Average Score of the Entered Student Records is %.2f\n", avgScore)

}

func getInput(studentData map[int]float64) {

  var rollNo int;
  var score float64;
  var numberOfStudents int;

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

    
  for student, score := range studentData {
       fmt.Printf("Rollno :%d Score: %.2f\n", student, score)
   }

  return
}

func getHighestScorer(studentData map[int]float64) {

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
        fmt.Printf("The Highest scoring student is of Roll no: %d and with a score of %.2f.\n", 
          key, value)

    }
  } 
  
}



