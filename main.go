package main

import (
	"ShapeCalcWithInterface/consoleHelper"
	"ShapeCalcWithInterface/shapes"
	"fmt"
	"strings"
)

//----------------------------------------------------------------

func getShapeName() string {
	result := consoleHelper.HandleUserInputString(1, "Please enter only one word.")
	return strings.ToLower(result[0])
}

//
//func getShapeOrCommandInput() (s shapes.Shape, retry bool, quit bool) {
//	s = nil
//	retry = false
//	quit = false
//
//	n := getShapeName()
//	switch n {
//	case "": //Некорректный ввод
//		return nil, true, false
//	case "q":
//		fmt.Println("User requested exit.")
//		return nil, false, true
//	default:
//		if s := shapes.GetShape(n); s == nil {
//			fmt.Println("Unknown shape")
//			return nil, true, false
//		} else {
//			return s, false, false
//		}
//	}
//}

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

		//s, retry, quit := getShapeOrCommandInput()
		//if quit {
		//	break
		//} else if retry {
		//	continue
		//}

		s.DimensionsInputPrompt()
		s.GetDimensionsInput()
		s.PrintCreationText()
		shapes.PrintShapeCalculations(s)
		fmt.Println("")
		fmt.Println("")
	}
}
