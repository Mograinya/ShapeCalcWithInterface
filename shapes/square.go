package shapes

import (
	"ShapeCalcWithInterface/consoleHelper"
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

func (sq Square) DimensionsInputPrompt() {
	fmt.Println("Enter the side of a Square:")
}

func (sq *Square) GetDimensionsInput() {
	for {
		r := consoleHelper.HandleUserInputFloat64(1, "Please enter one number")
		if r != nil {
			sq.side = r[0]
			break
		}
	}
}

func (sq Square) PrintCreationText() {
	fmt.Printf("A Square with %f sides is created", sq.side)
}
