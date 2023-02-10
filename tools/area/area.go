package area

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func Circle() {
	c1 := circle{10}
	c2 := circle{25}
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}
func Rectangle() {
	r1 := rectangle{12, 2}
	r2 := rectangle{9, 4}
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
}
