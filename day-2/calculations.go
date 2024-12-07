package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reports, err := readLevelsFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Safe levels: %d\n", countSafeLevels(reports))
}

func countSafeLevels(reports [][]int) int {
	safeLevels := 0
	for _, levels := range reports {
		safe := isSafe(levels)
		if !safe {
			for i := range levels {
				if safe = isSafe(removeByIndex(levels, i)); safe {
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

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	return (isDecreasing(levels) || isIncreasing(levels)) &&
		hasValidGaps(levels)
}

func removeByIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}

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

func readLevelsFromFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]int

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\s+`)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		values := re.Split(line, -1)

		if len(values) == 0 {
			continue
		}

		var levels []int
		for _, value := range values {
			level, err := strconv.Atoi(strings.TrimSpace(value))
			if err != nil {
				panic(fmt.Errorf("invalid integer value: %s", value))
			}

			levels = append(levels, level)
		}
		if len(levels) > 0 {
			reports = append(reports, levels)
		}
	}
	return reports, nil
}
