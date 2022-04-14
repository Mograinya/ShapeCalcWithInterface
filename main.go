package main

import (
	"ShapeCalcWithInterface/consoleHelper"
	"ShapeCalcWithInterface/shapes"
	"fmt"
	"strconv"
	"strings"
)

var quitMessage = "User requested exit."

func parseInitialInput(input []string) (handled, quit bool, message string) {
	if len(input) != 1 {
		return true, false, "Too many arguments"
	}

	if input[0] == "q" {
		return true, true, quitMessage
	}

	if input[0] == "shapes" {
		allShapes := strings.Join(shapes.GetShapesList(), ", ")
		return true, false, fmt.Sprintf("Acceptable shape names: %s\n", allShapes)
	}

	return false, false, ""
}

func handleDimensionsInput(s shapes.Shape) (bool, []float64) {
	inputPrompt, dimNum := s.DimensionsInputPrompt()
	for {
		fmt.Println(inputPrompt)
		input, err := consoleHelper.GetConsoleInput(50)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(input) == 1 && input[0] == "q" {
			return true, nil
		}

		if len(input) != dimNum {
			fmt.Printf("Incorrect input. Got %d dimensions instead of %d.", len(input), dimNum)
			continue
		}

		result := make([]float64, len(input))
		success := true
		for i, v := range input {
			f, e := strconv.ParseFloat(v, 64)
			if e != nil {
				fmt.Printf("Error: %s\n\n", e)
				success = false
				break
			}
			result[i] = f
		}

		if success {
			return false, result
		}
	}
}

func main() {
	fmt.Println("MANDATORY WELCOME MESSAGE")

	for {
		fmt.Println("\nEnter type of figure to proceed with calculations. " +
			"Enter 'shapes' to get a list of supported shapes. Enter 'q' to quit:")

		input, err := consoleHelper.GetConsoleInput(50)
		if err != nil {
			fmt.Println(err)
			continue
		}

		handled, quit, msg := parseInitialInput(input)
		if handled {
			fmt.Printf("%s\n", msg)
			if quit {
				return
			}
			continue
		}

		s := shapes.GetShape(input[0])
		if s == nil {
			fmt.Println("Unknown shape")
			continue
		}

		quit, dimensions := handleDimensionsInput(s)
		if quit {
			fmt.Println(quitMessage)
			return
		}

		msg = s.CreateShape(dimensions)
		fmt.Println(msg)

		shapes.PrintShapeCalculations(s)
		fmt.Println("")
		fmt.Println("")
	}
}
