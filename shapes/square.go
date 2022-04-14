package shapes

import (
	"fmt"
)

type Square struct {
	side float64
}

func NewSquare() *Square {
	return &Square{side: 0}
}

func (sq Square) Area() float64 {
	return sq.side * sq.side
}

func (sq Square) Perimeter() float64 {
	return 4 * sq.side
}

func (sq Square) DimensionsInputPrompt() (string, int) {
	return "Enter the side of a Square:", 1
}

func (sq *Square) CreateShape(f []float64) string {
	sq.side = f[0]
	return fmt.Sprintf("A Square with %g sides is created", sq.side)
}
