package main

import (
  "fmt"
  "time"
  "math/rand"
)

func f(from string) {
  for i := 0; i < 3; i++ {
    fmt.Println(from, ":", i)
    time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
  }
}

func main() {
  go f("goroutine1")
  go f("goroutine2")
  go func(msg string) {
    fmt.Println(msg)
  }("goroutine3")
  f("main")

  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}
