package main

import (
  "fmt"
)

//this is the bank account`` object
type Question struct {
  theQuestion string
  answers []string
  correctOne int
}
//this is the constructor
func newQuestion(theQuestion string, answers []string, correctOne int) *Question {
  q := Question{theQuestion: theQuestion, answers: answers, correctOne: correctOne}
  return &q
}

func main() {
  choices := [4]string{"a", "b", "c", "d"}
  var userInput string
  var a1 = []string{"Adam", "Eve", "Able", "Lucy"} 
  questionS := []*Question{
    newQuestion("Who was the first man created?", a1, 0),
    newQuestion("Who was the first man killed?", a1, 2),
    }
  fmt.Println("Welcome to the Trivia Game!")
  //this loops through the questionS
  for _, currentQuestion := range questionS {
    fmt.Println(currentQuestion.theQuestion)
    for index, currentChoice := range currentQuestion.answers {
      fmt.Println(choices[index], ") ", currentChoice)
    }
    fmt.Println("Type in the letter of your final answer...")
    fmt.Scanln(&userInput)
    if userInput == choices[currentQuestion.correctOne] {
      fmt.Println("you got it right!")
    }
  }

}
