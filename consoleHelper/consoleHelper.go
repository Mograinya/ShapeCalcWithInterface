package consoleHelper

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetConsoleInput(maxLen int) ([]string, error) {
	readBuffer := make([]byte, maxLen+2) //+2 to take into account CRLF
	overflowFlag := false
	overflowError := errors.New("input buffer overflow")
	inputString := ""
	for {
		n, err := os.Stdin.Read(readBuffer)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if n < 2 {
			if overflowFlag {
				return nil, overflowError
			}

			return nil, errors.New("input error")
		}

		if readBuffer[n-1] == '\n' && readBuffer[n-2] == '\r' {
			if overflowFlag {
				return nil, overflowError
			}
			inputString = string(readBuffer[:n-2])
			break
		}

		overflowFlag = true //not really a happy path, is it?
	}

	return strings.Fields(inputString), nil
}

// Эта хрень нарушает принцип DRY, но я пока не понял, как сделать так, чтоб оно
// само определяло типы. Нужно посмотреть внутрь функций fmt, они это умеют.
// Нашёл ещё это: https://go.dev/play/p/4q9vnXcrhP

func HandleUserInputString(requiredNum int, maxErr string) []string {
	faultValue := []string{""}
	scanned := make([]string, 5)
	counter := 0

	for {
		if num, err := fmt.Scanf("%s", &scanned[counter]); err != nil || num == 0 {
			break
		}
		counter++
	}

	if counter != requiredNum {
		fmt.Println(maxErr)
		return faultValue
	}

	return scanned
}

//func HandleUserInputInt(requiredNum int, maxErr string) []int {
//	var faultValue []int
//	scanned := make([]int, 5)
//	counter := 0
//
//	for {
//		if num, err := fmt.Scanf("%d", &scanned[counter]); err != nil || num == 0 {
//			break
//		}
//		counter++
//	}
//
//	if counter != requiredNum {
//		fmt.Println(maxErr)
//		return faultValue
//	}
//
//	return scanned
//}

func HandleUserInputFloat64(requiredNum int, maxErr string) []float64 {
	var faultValue []float64
	scanned := make([]float64, 5)
	counter := 0

	for {
		if num, err := fmt.Scanf("%f", &scanned[counter]); err != nil || num == 0 {
			break
		}
		counter++
	}

	if counter != requiredNum {
		fmt.Println(maxErr)
		return faultValue
	}

	return scanned
}
