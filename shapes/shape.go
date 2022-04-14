package shapes

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
	DimensionsInputPrompt() (string, int)
	CreateShape(f []float64) string
}

var shapeConstructorMap = map[string]Shape{
	"square":   NewSquare(),
	"circle":   NewCircle(),
	"triangle": NewTriangle(),
}

func GetShape(name string) Shape {
	if s, ok := shapeConstructorMap[name]; ok {
		return s
	}

	return nil
}

func GetShapesList() []string {
	result := make([]string, len(shapeConstructorMap))
	index := 0
	for k := range shapeConstructorMap {
		result[index] = k
		index++
	}
	return result
}

func PrintShapeCalculations(s Shape) {
	fmt.Printf("The area of this figure is %g\n", s.Area())
	fmt.Printf("The perimeter of this figure is %g\n\n\n", s.Perimeter())
}
