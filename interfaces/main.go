package main

import "fmt"

type shape interface {
  getArea() float64
}

func printArea(s shape) {
  fmt.Println("The figure has area", s.getArea())
}

type triangle struct {
  height float64
  base float64
}

func (t triangle) getArea() float64 {
  return t.base * t.height * 0.5
}

type square struct {
  sideLength float64
}

func (t square) getArea() float64 {
  return t.sideLength * t.sideLength
}


func main() {
  t := triangle{height: 10, base: 10}
  s := square{sideLength: 10}
  printArea(t)
  printArea(s)
}
