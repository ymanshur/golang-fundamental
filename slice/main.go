package main

import "fmt"

func main() {
	// var gamingConsoles []string
	// gamingConsoles = append(gamingConsoles, "PlayStation4")
	// gamingConsoles = append(gamingConsoles, "Nintendo Switch")
	// gamingConsoles = append(gamingConsoles, "XBox One")

	gamingConsoles := []string{"PlayStation4", "Nintendo Switch"}
	gamingConsoles = append(gamingConsoles, "XBox One")

	fmt.Println(gamingConsoles)

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}
}
