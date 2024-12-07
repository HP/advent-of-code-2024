package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
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

func readIntsFromFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\s+`)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		values := re.Split(line, -1)
		if len(values) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		leftValue, err := strconv.Atoi(strings.TrimSpace(values[0]))
		if err != nil {
			return nil, nil, fmt.Errorf("invalid integer value: %s", values[0])
		}
		rightValue, err := strconv.Atoi(strings.TrimSpace(values[1]))
		if err != nil {
			return nil, nil, fmt.Errorf("invalid integer value: %s", values[1])
		}

		left = append(left, leftValue)
		right = append(right, rightValue)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

func main() {

	left, right, err := readIntsFromFile("day1-input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total distance: %d\n", totalDistance(left, right))
	fmt.Printf("Total simplicity score: %d\n", totalSimplicityScore(left, right))
}
