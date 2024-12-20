package main

import (
  "fmt"
  "math"
)

//this is the bank account`` object
type SavingsAccount struct {
  amount float64
  interest float64
}
//this is the constructor
func newAccount(amount float64, interest float64) *SavingsAccount {
  a := SavingsAccount{amount: amount, interest: interest}
  return &a
}
func (a *SavingsAccount) withdraw(input float64) float64 {
  return a.amount - input
}
func (a *SavingsAccount) deposit(add float64) float64 {
  return a.amount + add
}
func (a *SavingsAccount) monthlyInterest() float64 {
  //we are dividing by 12 because what is in the account is the annual not the monthly rate
  mI := a.interest / 12
  mI = mI / 100
  mI = a.amount * mI
  mI = a.amount + mI
  mI = mI * 100
  mI = math.Round(float64(mI))
  mI = mI / 100
  return float64(mI)
}

func main()  {
  var input float64
  var months int
  fmt.Println("Welcome to the savings account calculator!")
  fmt.Println("Please input what will your % annual interest")
  fmt.Scanln(&input)
  account := newAccount(0, input)
  fmt.Println("How much money would you like to start off with?...")
  fmt.Scanln(&input)
  account.amount = account.deposit(input)
  fmt.Println(account)
  fmt.Println("How many months have passed since the account was opened?")
  fmt.Scanln(&months)
  for i := 0; i < months; i++ {
    fmt.Println("How much money was deposited in this month?")
    fmt.Scanln(&input)
    account.amount = account.deposit(input)
    fmt.Println("How much was withdrawn this month?")
    fmt.Scanln(&input)
    account.amount = account.withdraw(input)
    account.amount = account.monthlyInterest()
  }
  fmt.Println("Alright! This is you end balance: ", account.amount)
}
