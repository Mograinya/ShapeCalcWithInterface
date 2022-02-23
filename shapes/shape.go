package shapes

type Shape interface {
	SetDimensions(i interface{})
	Area() float64
	Perimeter() float64
	InputPrompt()
	GetInput()
	CreationText()
}
