package main

import "fmt"
import "math/rand"
import "time"

func main() {
  guests := [4]string{"Alfred", "Marcos", "Frida", "Putin"}
  objects := [3]string{"toy chest", "crate", "box"}

  fmt.Println(guests, objects)

  rand.Seed(time.Now().UnixNano())

  sliceGuests := guests[:]
  //sliceObjects := objects[:]
  myElement := getRandomElement(sliceGuests)
  fmt.Println(myElement)
}

func getRandomElement (slice []string) string {
  index := rand.Intn(len(slice))
  return slice[index]
}