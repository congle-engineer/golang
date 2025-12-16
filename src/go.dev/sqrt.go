package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := 1.0
  for i := 0; math.Abs(z * z - x) >= 0.0001; i++ {
    z -= (z * z - x) / (2 * z)
    fmt.Println("Round", i)
    fmt.Println("z:", z) 
  }
  return z
}

func main() {
  x := 2.0
  z := Sqrt(x)
  fmt.Println("final result:", z)
  fmt.Println("diff:", math.Abs(z * z - x))
}
