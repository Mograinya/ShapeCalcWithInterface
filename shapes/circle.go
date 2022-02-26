package shapes

import (
	"ShapeCalcWithInterface/consoleHelper"
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

func (c Circle) InputPrompt() {
	fmt.Println("Enter the radius of a Circle:")
}

func (c *Circle) GetInput() {
	for {
		r := consoleHelper.HandleUserInputFloat64(1, "Please enter one number")
		if r != nil {
			c.radius = r[0]
			break
		}
	}
}

func (c Circle) PrintCreationText() {
	fmt.Printf("A Circle with %f radius is created", c.radius)
}
