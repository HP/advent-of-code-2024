package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/HP/advent-of-code-2024/utils"
)

func main() {
	scanner, file, err := utils.GetFileScanner("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var code string
	for scanner.Scan() {
		line := scanner.Text()
		code += line
	}

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(code, -1)

	result := 0
	for _, match := range matches {
		if len(match) == 3 {
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				product := num1 * num2
				result += product
			}
		}
	}

	fmt.Println("Sum:", result)
}
