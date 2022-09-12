package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "Agung", "score": "A"},
		{"name": "Mario", "score": "B"},
		{"name": "Budi", "score": "E"},
	}

	// fmt.Println(students)

	for _, student := range students {
		// fmt.Println(student)
		fmt.Println(student["name"], student["score"])
	}
}
