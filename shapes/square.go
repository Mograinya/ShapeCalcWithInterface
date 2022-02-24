package shapes

import (
	"fmt"
	"log"
)

type Square struct {
	side float64
}

func NewSquare() *Square {
	return &Square{side: 0}
}

func (sq *Square) SetDimensions(i interface{}) {
	value, ok := i.(float64)

	if ok {
		sq.side = value
	} else {
		log.Print("Ошибка задания стороны квадрата.")
	}
}

func (sq Square) Area() float64 {
	return sq.side * sq.side
}

func (sq Square) Perimeter() float64 {
	return 4 * sq.side
}

func (sq Square) InputPrompt() {
	fmt.Println("Enter the side of a Square:")
}

func (sq *Square) GetInput() {
	fmt.Scanf("%f", &sq.side)
}

func (sq Square) PrintCreationText() {
	fmt.Printf("A Square with %f sides is created", sq.side)
}
