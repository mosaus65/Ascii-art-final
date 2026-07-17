package main

import (
	"strings"
)

func Generate(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	tar := strings.Split(input, "\\n")

	result := []string{}
	for _, word := range tar {
		line := make([]string, 9)
		for _, char := range word {
			art, ok := banner[char]
			if !ok {
				return ""

			}
			for i := 0; i < 8; i++ {
				line[i] += art[i]

			}
		}
		result = append(result, strings.Join(line, "\n"))
	}
	return strings.Join(result, "\n")
}
