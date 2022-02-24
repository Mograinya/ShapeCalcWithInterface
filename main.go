package main

import (
	"ShapeCalcWithInterface/consoleHelper"
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
	result := consoleHelper.HandleUserInputString(1, "Please enter only one word.")
	return strings.ToLower(result[0])
}

func main() {
	for {
		fmt.Println("Enter type of figure. Enter 'q' to quit:")

		n := getShapeName()
		if n == "" {
			continue
		} else if n == "q" {
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
		s.PrintCreationText()
		printShapeAreaAndPerimeterInfo(s)
		fmt.Println("")
		fmt.Println("")
	}
}
