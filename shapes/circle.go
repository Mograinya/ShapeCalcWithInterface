package shapes

import (
	"fmt"
	"log"
	"math"
)

type Circle struct {
	radius float64
}

func NewCircle() *Circle {
	return &Circle{radius: 0}
}

func (c *Circle) SetDimensions(i interface{}) {
	value, ok := i.(float64)

	if ok {
		c.radius = value
	} else {
		log.Print("Ошибка задания радиуса круга.")
	}
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) InputPrompt() {
	fmt.Println("Enter the side of a Square:")
}

func (c *Circle) GetInput() {
	fmt.Scanf("%f", &c.radius)
}

func (c Circle) CreationText() {
	fmt.Printf("A Circle with %f radius is created", c.radius)
}
