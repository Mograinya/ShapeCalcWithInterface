package shapes

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
	DimensionsInputPrompt()
	GetDimensionsInput()
	PrintCreationText()
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

func PrintShapeCalculations(s Shape) {
	fmt.Printf("\nThe area of this figure is %f", s.Area())
	fmt.Printf("\nThe perimeter of this figure is %f", s.Perimeter())
}
