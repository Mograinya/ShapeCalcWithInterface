package shapes

type Shape interface {
	SetDimensions(i interface{})
	Area() float64
	Perimeter() float64
	InputPrompt()
	GetInput()
	CreationText()
}

func GetShape(name string) Shape {
	switch name {
	case "square":
		return NewSquare()
	case "circle":
		return NewCircle()
	case "triangle":
		return nil
	default:
		return nil
	}
}
