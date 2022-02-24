package main

import (
	"ShapeCalcWithInterface/shapes"
	"fmt"
	"strings"
)

//----------------------------------------------------------------

func printShapeAreaAndPerimeterInfo(s shapes.Shape) {
	fmt.Printf("\nThe area of this figure is %f", s.Area())
	fmt.Printf("\nThe perimeter of this figure is %f", s.Perimeter())
}

func getShapeName() string {
	var n string
	fmt.Scanf("%s", &n)
	fmt.Scanln() // Избавиться от несчитанного \n
	return strings.ToLower(n)
}

func main() {
	for {
		fmt.Println("Enter type of figure. Enter 'q' to quit:")

		n := getShapeName()
		if n == "q" || n == "Q" { // Проверка на запрос выхода
			fmt.Println("User requested exit.")
			break
		}

		s := shapes.GetShape(n)
		if s == nil {
			fmt.Println("Unknown shape")
			continue
		}

		s.InputPrompt()
		s.GetInput()
		fmt.Scanln() // Избавиться от несчитанного \n
		s.PrintCreationText()
		printShapeAreaAndPerimeterInfo(s)
		fmt.Println("")
		fmt.Println("")
	}
}
