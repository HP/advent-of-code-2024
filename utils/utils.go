package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetFileScanner(filename string) (*bufio.Scanner, *os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	return bufio.NewScanner(file), file, nil
}

func ReadLinesFromFile(filename string) ([]string, error) {
	scanner, file, err := GetFileScanner(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func ReadIntsFromFile(filename string) ([][]int, error) {
	scanner, file, err := GetFileScanner(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]int

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

func RemoveByIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}
