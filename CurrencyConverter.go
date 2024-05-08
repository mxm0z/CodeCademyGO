package main

import (
	"fmt"
)

func main() {
  currencies := map[string]float32{
    "JPY": 130.2,
    "EUR": 0.95,
    "BRL": 32.31,
    "USD": 0.13,
    "CAD": 2.4,
  }
  var dollarAmount float32
  var currency string

  fmt.Println("Insert the amount of dollars:")
  fmt.Scan(&dollarAmount)

  if dollarAmount == 0 {
    fmt.Println("Inserted amount not valid.")
  } else {
    fmt.Println("Insert target currency:")
    fmt.Scan(&currency)
    rate, isValid := currencies[currency]
    if isValid {
      fmt.Println("The final amount is:", rate * dollarAmount)
    } else {
      fmt.Println("Not a valid currency")
    }
  }
}