package main

import (
  "fmt"
)

func main() {
  var name string
  var age uint
  fmt.Println("What's your name?")
  fmt.Scan(&name)
  fmt.Println("What's your age?")
  fmt.Scan(&age)

  newMessage := fmt.Sprintf("Hello %v, you are %d years old", name, age)

  fmt.Println(newMessage)
}