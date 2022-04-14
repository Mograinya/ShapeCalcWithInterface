package shapes

import (
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

func (t Triangle) DimensionsInputPrompt() (string, int) {
	return "Enter three sides of a triangle (separated with spaces):", 3
}

func (t *Triangle) CreateShape(f []float64) string {
	t.a, t.b, t.c = f[0], f[1], f[2]
	return fmt.Sprintf("A Triangle with %g, %g, %g sides is created", t.a, t.b, t.c)
}
