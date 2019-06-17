package main

import "fmt"

func convertCelciusFarenheit(celcius float64) float64 {
  return (celcius * 9/5) + 32
}

func main() {
    fmt.Printf("This program converts Celcius into Farenheit!\n")

  for {
    celcius := 0.0

    fmt.Printf("Insert a value in Celcius:\n-> ")
    fmt.Scanf("%f", &celcius)

    farenheit := convertCelciusFarenheit(celcius)

    fmt.Printf("%.2f Fº\n", farenheit)
  }
}
