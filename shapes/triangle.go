package shapes

import (
	"ShapeCalcWithInterface/consoleHelper"
	"fmt"
	"math"
)

type Triangle struct {
	a, b, c float64
}

func NewTriangle() *Triangle {
	return &Triangle{a: 0, b: 0, c: 0}
}

func (t Triangle) Area() float64 {
	p := (t.a + t.b + t.c) / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) InputPrompt() {
	fmt.Println("Enter three sides of a triangle (separated with spaces):")
}

func (t *Triangle) GetInput() {
	for {
		r := consoleHelper.HandleUserInputFloat64(3, "Please enter three numbers")
		if r != nil {
			t.a, t.b, t.c = r[0], r[1], r[2]
			break
		}
	}
}

func (t Triangle) PrintCreationText() {
	fmt.Printf("A Triangle with %f, %f, %f sides is created", t.a, t.b, t.c)
}
