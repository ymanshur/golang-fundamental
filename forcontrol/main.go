package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Saya sedang belajar Go")

	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Saya sedang belajar Go :", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Saya sedang belajar Go :", i)
	// 	i++
	// }

	title := "Golang the best language"

	for index, letter := range title {
		// fmt.Println("index :", index, " letter :", string(letter))

		// if index%2 == 0 {
		// 	fmt.Println("index :", index, " letter :", string(letter))
		// }

		letterString := strings.ToLower(string(letter))

		// if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
		// 	fmt.Println("index :", index, " letter :", string(letter))
		// }

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index :", index, " letter :", string(letter))
		}
	}
}
