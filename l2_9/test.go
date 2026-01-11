package main

import (
	"fmt"
	"testing"
)

func TestTransform(t *testing.T) {
	strArr := []string{
		"a4bc2d5e",
	}

	for _, str := range strArr {
		result := transform(str)
		fmt.Println(result)
	}
}
