package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func NewCircle() *Circle {
	return &Circle{radius: 0}
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) DimensionsInputPrompt() (string, int) {
	return "Enter the radius of a Circle:", 1
}

func (c *Circle) CreateShape(f []float64) string {
	c.radius = f[0]
	return fmt.Sprintf("A Circle with %g radius is created", c.radius)
}
