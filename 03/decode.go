package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

	re := regexp.MustCompile(`(?:do\(\)|don't\(\))?.*?mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(code, -1)

	doCompute := true
	result := 0
	for _, match := range matches {

		if strings.Contains(match[0], "do()") {
			fmt.Println("do", match[0])
			doCompute = true
		}
		if strings.Contains(match[0], "don't()") {
			fmt.Println("don't", match[0])
			doCompute = false
		}
		if !doCompute {
			continue
		}

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
