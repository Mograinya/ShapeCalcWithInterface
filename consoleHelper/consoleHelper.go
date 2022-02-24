package consoleHelper

import "fmt"

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
