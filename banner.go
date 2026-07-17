package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(filename string)(map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("banner file is nil")
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("banner file is empty")

	}
	character := strings.ReplaceAll(string(data), "\r\n", "\n")

	
	line := strings.Split(character, "\n")

	banner := map[rune][]string{}

	currentchar := rune(32)

	for i := 1; i < len(line); i += 9 {
		if i+8 > len(line) {
			break
		}
			charArt := line[i : i+8]
			banner[currentchar] = charArt
			currentchar++
		}
		if len(banner) != 95 {
			return nil, fmt.Errorf("banner file is empty")
		}
		return	banner, nil
	}
	
