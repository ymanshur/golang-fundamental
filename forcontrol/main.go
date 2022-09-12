package main

import "fmt"

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

	for _, letter := range title {
		// fmt.Println("index :", index, " letter :", string(letter))
		fmt.Println("letter :", string(letter))
	}
}
