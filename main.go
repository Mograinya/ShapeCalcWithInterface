package main

import (
	"ShapeCalcWithInterface/shapes"
	"fmt"
	"strings"
)

//----------------------------------------------------------------

func printShapeSummary(s shapes.Shape) {
	s.CreationText()
	fmt.Printf("\nThe area of this figure is %f", s.Area())
	fmt.Printf("\nThe perimeter of this figure is %f", s.Perimeter())
}

func getShapeName() string {
	var n string
	fmt.Println("Enter type of figure:")
	fmt.Scanf("%s", &n)
	fmt.Scanln() // Избавиться от несчитанного \n
	return strings.ToLower(n)
}

func main() {
	n := getShapeName()
	s := shapes.GetShape(n)

	if s == nil {
		fmt.Println("Unknown shape")
		return
	}
	s.InputPrompt()
	s.GetInput()
	printShapeSummary(s)
}
