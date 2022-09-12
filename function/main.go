package main

import (
	"fmt"
)

func main() {
	// sentence := printMyResult("Saya sedang")

	// result := add(10, 20)

	area1, around1 := calculate(10, 20)
	// area, _ := calculate(10, 20)

	fmt.Println(area1)
	fmt.Println(around1)
}

// func calculate(length, width int) (int, int) {
// 	area := length * width
// 	around := 2 * (length + width)

// 	return area, around
// }

func calculate(length, width int) (area int, around int) { // pre-defin e return value
	// var area int
	area = length * width

	// var around int
	around = 2 * (length + width)

	// return area, around
	return
}

func add(number, numberTwo int) int {
	// result := number + numberTwo
	return number + numberTwo
}

func printMyResult(sentence string) string {
	// fmt.Println(sentence)

	newSentence := sentence + " belajar Golang"

	return newSentence
}
