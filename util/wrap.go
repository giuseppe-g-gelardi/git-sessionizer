package util

import (
	"strings"
)

func WrapText(text string, width int) string {
	words := strings.Fields(text)
	var lines []string
	var currentLine string

	for _, word := range words {
		if len(currentLine)+len(word) <= width {
			currentLine += " " + word
		} else {
			lines = append(lines, strings.TrimSpace(currentLine))
			currentLine = word
		}
	}

	lines = append(lines, strings.TrimSpace(currentLine))
	return strings.Join(lines, "\n")
}
