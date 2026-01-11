package main

import (
	"errors"
	"fmt"
)

var errInvalidInput = errors.New("invalid input")

func main() {
	strArr := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe\\4\\5",
		"qwe\\45",
	}

	for _, str := range strArr {
		result, err := transform(str)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}

func transform(input string) (string, error) {
	var prev rune
	var tr bool
	var digitCount int
	result := make([]rune, 0, len(input))

	for _, curr := range input {
		if tr || prev == '\x00' {
			tr = false
			prev = curr
			if isDigit(curr) {
				digitCount++
			}
			continue
		}

		if isDigit(curr) {
			digitCount++
			tr = true

			if prev == '\\' {
				result = append(result, curr)
				continue
			}

			count := curr - '0'
			for range count {
				result = append(result, prev)
			}
		} else {
			result = append(result, prev)
		}
		prev = curr
	}
	if !tr && prev != '\x00' {
		if isDigit(prev) {
			count := prev - '0'
			count--
			prev = result[len(result)-1]
			for range count {
				result = append(result, prev)
			}
		} else {
			result = append(result, prev)
		}
	}

	if digitCount == len(input) {
		return "", errInvalidInput
	}

	return string(result), nil
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
