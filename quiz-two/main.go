package main

import "fmt"

func main() {
	// menghitung rata-rata
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	scoresLength := len(scores)

	var sumOfScores int
	var goodScores []int

	for _, score := range scores {
		sumOfScores += score

		if score >= 70 {
			goodScores = append(goodScores, score)
		}
	}

	result := float64(sumOfScores) / float64(scoresLength)

	fmt.Println(result)
	fmt.Println(goodScores)
}
