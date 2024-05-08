package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true

  if eludedGuards := rand.Intn(100); eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }

  fmt.Println("Heist status:", isHeistOn)

  if openedVault := rand.Intn(100); isHeistOn && openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn && openedVault < 70 {
    isHeistOn = false
    fmt.Println("The vault can not be opened.")
  }

  if leftSafely := rand.Intn(5); isHeistOn {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("No one left safely. The heist has failed completely!")
      case 5:
        fmt.Println("Start the getaway car!")
      default:
        fmt.Println("Some members didn't make it.")
    }
  }

  if amtStolen := 10000 + rand.Intn(1000000); isHeistOn {
    fmt.Println("The total amount stole was", amtStolen)
  }
}
