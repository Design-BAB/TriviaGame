package main

import (
    "database/sql"
    "fmt"
    _ "github.com/glebarez/go-sqlite"
)

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

func Insert(db *sql.DB, q *Question) (*Question) {
    sql := `INSERT INTO qanda (question, answer1, answer2, answer3, answer4, rightanswer) 
            VALUES (?, ?, ?, ?, ?, ?);`
    //the Exec command is the one that actually pulls that command in the database
    result := db.Exec(sql, q.theQuestion, q.answers[0], q.answers[1], q.answers[2], q.answers[3], q.correctOne)
    return q()
}

func Delete(db *sql.DB, id int) (int64, error) {
  //Note when it is a single line command you use double qoutes
  sql := "DELETE FROM qanda WHERE id = ?"
  result, err := db.Exec(sql, id)
  if err != nil {
    return 0, err
  }
  return result.RowsAffected()
}
//
//
//this is the function that doesn't work
func findById(db *sql.DB, id int) (*Question, error) {
  sql := "SELECT question, answer1, answer2, answer3, answer4, rightanswer FROM qanda WHERE id = ?"
  row := db.QueryRow(sql, id)
  q := &Question{}
  err := row.Scan(&q.theQuestion, &q.answers[0], &q.answers[1], &q.answers[2], &q.answers[3], &q.correctOne)
  if err != nil {
    return nil, err
  }
  return q, nil
}

//note to myself, did the first tuturial. This basicly looks for my.db. Now i gotta make the other part where i make a database
func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite", "./q.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
    fmt.Println("Connected to the SQLite database successfully.")

    //now we will try to add a new question and see if this works
    // it works
    var a1 = []string{"Adam", "Eve", "Able", "Lucy"}
    firstManQuestion := newQuestion("Who was the first man made?", a1, 0)
    theNewestQuestion := findById(db, 1)
    fmt.Println("Here is one question: ", firstManQuestion)
    fmt.Println("Ok! Here is another one! This is the question: ", theNewestQuestion)
  }
