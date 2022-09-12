package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// scores := []int{10, 5, 8, 9, 7}
	// total := sum(scores)

	// fmt.Println(total)

	// result, err := calculate(10, 2, "+")
	// result, err := calculate(10, 2, "-")
	// result, err := calculate(10, 2, "*")
	// result, err := calculate(10, 2, "/")
	result, err := calculate(10, 2, "=")

	if err != nil {
		// fmt.Println(err.Error())
		log.Fatal(err)
	}

	fmt.Println(result)
}

func sum(scores []int) int {
	var total int

	for _, score := range scores {
		total += score
	}

	return total
}

func calculate(number, numberTwo int, operator string) (result int, errorResult error) {
	switch operator {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("unknown operation")
	}

	return
}
