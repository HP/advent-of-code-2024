package main

import (
	"fmt"

	"github.com/HP/advent-of-code-2024/utils"
)

func isDecreasing(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		if levels[i] > levels[i-1] {
			return false
		}
	}
	return true
}

func isIncreasing(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		if levels[i] < levels[i-1] {
			return false
		}
	}
	return true
}

func hasValidGaps(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		gap := levels[i] - levels[i-1]
		if gap == 0 || gap > 3 || gap < -3 {
			return false
		}
	}
	return true
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	return (isDecreasing(levels) || isIncreasing(levels)) &&
		hasValidGaps(levels)
}

func countSafeLevels(reports [][]int) int {
	safeLevels := 0
	for _, levels := range reports {
		safe := isSafe(levels)
		if !safe {
			for i := range levels {
				if safe = isSafe(utils.RemoveByIndex(levels, i)); safe {
					break
				}
			}
		}
		if safe {
			safeLevels++
		}
	}
	return safeLevels
}

func main() {
	reports, err := utils.ReadIntsFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Safe levels: %d\n", countSafeLevels(reports))
}
