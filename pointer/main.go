package main

import "fmt"

func main() {
	// numberA := 5
	// numberB := &numberA // referencing

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB) // dereferencing

	// *numberB = 10

	// fmt.Println(numberB)
	// fmt.Println(*numberB)
	// fmt.Println(numberA)

	var numberA int = 5
	var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// *numberB = 10

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	fmt.Printf("Nilai awal (%v): %v\n", numberA, numberB)

	change(&numberA, 9)

	fmt.Println("Apakah numberA sama dengan *numberB?", numberA == *numberB)
	fmt.Printf("Nilai akhir (%v): %v\n", *numberB, numberB)
}

func change(old *int, new int) {
	// old = new
	*old = new

	fmt.Printf("Nilai di dalam, function (%v): %v\n", *old, old)
}
