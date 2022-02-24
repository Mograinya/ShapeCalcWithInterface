package shapes

type Shape interface {
	Area() float64
	Perimeter() float64
	InputPrompt()
	GetInput()
	PrintCreationText()
}

//var shapeConstructorMap = map[string]func() Square{
//	"square": NewSquare(),
//	"circle": NewCircle(),
//  "triangle": NewTriangle(),
//}

func GetShape(name string) Shape {
	switch name {
	case "square":
		return NewSquare()
	case "circle":
		return NewCircle()
	case "triangle":
		return NewTriangle()
	default:
		return nil
	}
}
