package three

import (
	"fmt"
	"regexp"
	"strconv"
)

func solution(input string) int {
	regex := regexp.MustCompile("mul\\(([0-9+]{0,}),([0-9+]{0,})\\)")
	matches := regex.FindAllStringSubmatch(input, -2)

	var total int = 0
	for _, match := range matches {
		multiplier, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error parsing multiplier")
			return -1
		}
		value, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println("Error parsing value")
			return -1
		}

		total += multiplier * value
	}
	return total
}

func part2(input string) int {
	regex := regexp.MustCompile("(do|don\\'t|mul)\\((([0-9+]{0,}),([0-9+]{0,}))?\\)")
	matches := regex.FindAllStringSubmatch(input, -2)

	var total int = 0
	active := true
	for _, match := range matches {
		function := match[1]
		if function == "mul" && active {
			multiplier, err := strconv.Atoi(match[3])
			if err != nil {
				fmt.Println("Error parsing multiplier")
				return -1
			}
			value, err := strconv.Atoi(match[4])
			if err != nil {
				fmt.Println("Error parsing value")
				return -1
			}

			total += multiplier * value
		} else {
			active = function == "do"
		}
	}
	return total
}
