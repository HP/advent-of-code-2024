package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/HP/advent-of-code-2024/utils"
)

func totalDistance(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0

	for i := range left {
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}
	return totalDistance
}

func totalSimplicityScore(left, right []int) int {
	rightCounts := countOccurrences(right)
	simplicityScore := 0

	for _, value := range left {
		if count, ok := rightCounts[value]; ok {
			simplicityScore += value * count
		} else {
			simplicityScore += value * 0
		}

	}

	return simplicityScore
}

func countOccurrences(slice []int) map[int]int {
	counts := make(map[int]int)
	for _, value := range slice {
		counts[value]++
	}
	return counts
}

func main() {
	ints, err := utils.ReadIntsFromFile("input.txt")
	if err != nil {
		panic(err)
	}
	left := make([]int, len(ints))
	right := make([]int, len(ints))

	for i, values := range ints {
		left[i] = values[0]
		right[i] = values[1]
	}

	fmt.Printf("Total distance: %d\n", totalDistance(left, right))
	fmt.Printf("Total simplicity score: %d\n", totalSimplicityScore(left, right))
}
