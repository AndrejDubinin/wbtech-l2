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
	result := make([]rune, 0, len(input))

	for _, char := range input {
		if prev == '\x00' {
			prev = char
			continue
		}

		if isDigit(char) {
			if isDigit(prev) {
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
	if !isDigit(prev) && prev != '\x00' {
		result = append(result, prev)
	}
	return string(result)
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
