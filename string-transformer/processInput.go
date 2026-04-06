package main

import (
	"fmt"
	"os"
	"strings"
)

func ProcessText(text string) string {
	if text == "quit" {
		fmt.Println("Goodbye")
		os.Exit(0)
	}

	input := strings.Fields(text)

	if len(input) < 2 {
		fmt.Println("Invalid arg length")
	}

	switch input[0] {
	case "upper":
		result := Upper(text)
		return result

	case "lower":
		result := ToLower(text)
		return result

	case "cap":
		result := Caps(text)
		return result

	case "title":
		result := Titlecase(text)
		return result

	case "snake":
		result := Snake(text)
		return result

	case "reverse":
		result := ReverseString(text)
		return result

	default:
		fmt.Println("Invalid transformation")
		fmt.Println("Transformations:\n 1. upper <text>\n 2. lower <text>\n 3. cap <text>\n 4. title <text>\n 5. snake <text>\n 6. reverse <text>")
	}
	return text
}
