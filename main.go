package main
import (
	"fmt"
	"math"
)

//interface Shaper
type Shaper interface {
  Area() float64
  Perimeter() float64
  Info()
}

// class Rectangle
type Rectangle struct {
  width float64
  height float64
}

func(r Rectangle) Area() float64 {
  return r.width * r.height
}

func(r Rectangle) Perimeter() float64 {
  return 2.0 * (r.width + r.height)
}

func (r Rectangle) Info() {
  fmt.Printf("Area=%.2f, Perimeter=%.2f\n", r.Area(), r.Perimeter())
}


//class Circle

type Circle struct {
  radius float64
}

func(c Circle) Area() float64 {
  return math.Pi * c.radius * c.radius
}

func(c Circle) Perimeter() float64 {
  return 2 * math.Pi * c.radius
}

func (c Circle) Info() {
  fmt.Printf("Area=%.2f, Perimeter=%.2f\n", c.Area(), c.Perimeter())
}



func main() {
  fmt.Printf("__RectangleInfo__")
	r := Rectangle{ width: 5, height: 5 }
	r.Info()
	fmt.Println("============================")
	
	fmt.Println("__CircleInfo__")
	c := Circle{ radius: 10 }
	c.Info()
}
