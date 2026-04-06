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
	"slices"
	"strings"
)

func Upper(text string) string {
	field := strings.Fields(text)

	for i := range field {

		field[i] = strings.ToUpper(field[i])
	}
	return strings.Join(field[1:], " ")
}

func ToLower(text string) string {
	words := strings.Fields(text)
	word := strings.ToLower(strings.Join(words[1:], " "))
	return word
}

func Caps(text string) string {
	field := strings.Fields(text)

	for i := range field {
		field[i] = strings.ToUpper(string(field[i][0])) + strings.ToLower(field[i][1:])
	}
	return strings.Join(field[1:], " ")
}

func Titlecase(text string) string {
	fields := strings.Fields(text)

	connect := []string{"a", "an", "the", "and", "but", "or", "for", "nor", "on", "at", "to", "by", "in", "of", "up", "as", "is", "it"}

	//iterate fields
	for i, val := range fields {
		//range through connect
		for _, con := range connect {
			//check if connect contains val of fields
			myflag := slices.Contains(connect, val)
			//if my returns false the implement titlecase to the false elements
			if !myflag {
				fields[i] = strings.ToUpper(string(fields[i][0])) + strings.ToLower(string(fields[i][1:]))
			}
			//if first string returns true implement titlecase on it
			if fields[1] == con {
				fields[1] = strings.ToUpper(string(fields[1][0])) + strings.ToLower(string(fields[1][1:]))
			}
		}
	}
	return strings.Join(fields[1:], " ")
}

// func snake(text string) string {
// 	fields := strings.Fields(text)

// 	for i := range fields {
// 		if strings.HasSuffix(fields[i], ".,';!?") {
// 			fields[i] = strings.TrimSuffix(fields[i], "!.,;'?")

//			}
//		}
//		return strings.Join(fields, "_")
//	}
func Snake(text string) string {

	fields := strings.Fields(strings.ToLower(text))

	for i := range fields {
		fields[i] = strings.TrimRight(fields[i], ".,?!")
	}
	return strings.Join(fields[1:], "_")
}

func ReverseString(text string) string {
	words := strings.Fields(text)

	for i, word := range words {
		runes := []rune(word)
		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
			runes[j], runes[k] = runes[k], runes[j]
		}
		words[i] = string(runes)

	}
	return strings.Join(words[1:], " ")
}
