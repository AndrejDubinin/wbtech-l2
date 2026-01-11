package main

import (
	"fmt"
)

func main() {
	strArr := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
	}

	for _, str := range strArr {
		result := transform(str)
		fmt.Println(result)
	}
}

func transform(input string) string {
	var prev rune
	result := make([]rune, len(input))

	for i, char := range input {
		if isDigit(char) {
			if i == 0 {
				return ""
			}

			count := char - '0'
			for range count {
				result = append(result, prev)
			}
		} else if !isDigit(prev) {
			result = append(result, prev)
		}
		prev = char
	}
	if !isDigit(prev) {
		result = append(result, prev)
	}
	return string(result)
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
