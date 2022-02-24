package shapes

import (
	"fmt"
	"log"
	"math"
)

type Triangle struct {
	sideA, sideB, sideC float64
}

func NewTriangle() *Triangle {
	return &Triangle{sideA: 0, sideB: 0, sideC: 0}
}

func (c *Triangle) SetDimensions(i interface{}) {
	value, ok := i.([3]float64)

	if ok {
		c.sideA, c.sideB, c.sideC = value[0], value[1], value[2]
	} else {
		log.Print("Ошибка задания радиуса круга.")
	}
}

func (c Triangle) Area() float64 {
	p := (c.sideA + c.sideB + c.sideC) / 2
	return math.Sqrt(p * (p - c.sideA) * (p - c.sideB) * (p - c.sideC))
}

func (c Triangle) Perimeter() float64 {
	return c.sideA + c.sideB + c.sideC
}

func (c Triangle) InputPrompt() {
	fmt.Println("Enter three sides of a triangle (separated with spaces):")
}

func (c *Triangle) GetInput() {
	fmt.Scanf("%f %f %f", &c.sideA, &c.sideB, &c.sideC)
}

func (c Triangle) PrintCreationText() {
	//fmt.Printf("A Triangle with %f, %f, %f sides is created", c.sideA, c.sideB, c.sideC)

	fmt.Printf("A Triangle with %f, %f, %f sides is created", c.sideA, c.sideB, c.sideC)
}
