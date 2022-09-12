package main

import (
	"fmt"
)

func main() {
	number := 5

	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("DEFAULT")
	}

	// if number == 1 {
	// 	fmt.Println("Satu")
	// } else if number == 2 {
	// 	fmt.Println("Dua")
	// } else if number == 3 {
	// 	fmt.Println("Tiga")
	// }
}
