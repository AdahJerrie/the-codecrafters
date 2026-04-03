package main

import "strings"

func processText(text string) string {
	if strings.Contains(text, "(cap)") {
		word := Cap(text)
		return word
	}
	if strings.Contains(text, "(low)") {
		word := Tolower(text)
		return word
	}
	if strings.Contains(text, "TODO:") {
		word := Todo(text)
		return word
	}

	if strings.Contains(text, "(Reverse)") {
		word := Reverse(text)
		return word
	}

	if len(text) > 80 {
		result := Linelenght(text)
		return result
	}

	if strings.Contains(text, "(cap)") && strings.Contains(text, "(low)") {
		return "Text cannot contain two transformations!"
	}

	return text
}

func processbylines(text string) string {
	lines := strings.Split(text, "\n")
	//Number = len(lines)

	for i, line := range lines {
		lines[i] = processText(line)
	}
	return strings.Join(lines, "\n")
}
