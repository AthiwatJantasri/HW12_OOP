package main

import (
	"fmt"
	"math"
)

type Shaper interface {
  Area() float64
  Perimeter() float64
  Info()
}

type Rectangle struct {
  width float64
  height float64
}

func(r Rectangle) Area() float64 {
  return r.width * r.height
}

func(r Rectangle) Perimeter() float64 {
  return 2 * (r.width + r.height)
}

func(r Rectangle) Info(){
  fmt.Printf("Area=%.2f, Perimeter=%.2f\n", r.Area(), r.Perimeter())
}

type Circle struct {
  radius float64
}

func (c Circle) Area() float64 {
  return math.Pi * c.radius
}

func (c Circle) Perimeter() float64 {
  return 2 * math.Pi * c.radius
}

func(c Circle) Info(){
  fmt.Printf("Area=%.2f, Perimeter=%.2f\n", c.Area(), c.Perimeter())
}



func main() {
	r := Rectangle{width: 4, height: 5}
	r.Info()

	fmt.Println("------------")

	c := Circle{radius: 3}
	c.Info()
}
