/*Text transformation tool: The program accepts a command followed by a string and applies the correct transformation.

The six transformation
1. upper <text>
2. lower <text>
3 cap <text>
4. title <text>
5. snake <text>
6. reverse <text>

steps --- write a function for each transformation

*/

package main

import (
	"fmt"
	"strings"
)

func Upper(text string) string {
	field := strings.Fields(text)
	for i := range field {
		if field[0] == "upper" {
			field[i] = strings.ToUpper(field[i])
		}
	}
	field[0] = ""
	return strings.Join(field, " ")
}

func ToLower(text string) string {
	word := strings.ToLower(text)
	return word
}

func Caps(text string) string {
	field := strings.Fields(text)

	for i := range field {
		field[i] = strings.ToUpper(string(field[i][0])) + strings.ToLower(field[i][1:])
	}
	return strings.Join(field, " ")
}

func Titlecase(text string) string {
	fields := strings.Fields(text)

	connect := []string{"a", "an", "the", "and", "but", "or", "for", "nor", "on", "at", "to", "by", "in", "of", "up", "as", "is", "it"}

	for _, val := range fields {
		for _, con := range connect {
			if fields[0] == con {
				fields[0] = strings.ToUpper(string(fields[0][0])) + strings.ToLower(fields[0][1:])
			} else if val != con {
				val = strings.ToUpper(string(val[0])) + strings.ToLower(val[1:])
			}
		}
	}
	return strings.Join(fields, " ")
}

func ReverseString(text string) string {
	runes := []rune(text)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	word := "fall of the western power grid"
	fmt.Println(Titlecase(word))
	// start:
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	fmt.Println("Enter text to be edited or type 'quit' to exit")
	// 	if !scanner.Scan() {
	// 		fmt.Println("Try again")
	// 		goto start
	// 	}

	// 	lines := scanner.Text()
	// 	//input := strings.Fields(lines)
	// 	//fmt.Println(input)

	// result := Caps(lines)
	// fmt.Println(result)
}
