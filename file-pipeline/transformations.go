/*project file-pipeline
transformations - i. caps(title case)
				  ii. All lowercase(to lower all uppercase)
				  iii. lines starting with todo: replace the todo with ACTION
				  iv.Dashes/blanks(remove such lines)
				  v. flag any line longer than 80 words with truncated
*/

package main

import (
	"slices"
	"strings"
)

func Cap(text string) string {

	connectors := []string{"an", "to", "the", "and", "but", "for", "or", "at", "to", "by", "on", "in", "of", "as", "is", "it"}
	//field := strings.ToLower(text)
	fields := strings.Fields(text)

	for i, val := range fields {
		last := len(fields) - 1
		if fields[last] == "(cap)" {
			for _, con := range connectors {
				if val == con {

				}

			}
			myFlag := slices.Contains(connectors, val)
			if myFlag == false {
				fields[i] = strings.ToUpper(string(fields[i][0])) + strings.ToLower(fields[i][1:])
			}

		}

	}
	fields[len(fields)-1] = ""
	return strings.Join(fields, " ")
}

func Tolower(text string) string {
	fields := strings.Fields(text)
	last := len(fields) - 1

	for i := range fields {

		if fields[last] == "(low)" {
			fields[i] = strings.ToLower(fields[i])
		}

	}
	fields[last] = ""
	return strings.Join(fields, " ")
}

func Todo(text string) string {
	fields := strings.Fields(text)

	for i := range fields {
		if fields[i] == "TODO:" {
			fields[i] = "ACTION:"
		}
	}
	return strings.Join(fields, " ")
}

func Linelenght(text string) string {
	for i := 0; i < len(text); i++ {
		if len(text) > 80 {
			//text = "TRUNCATED"
		}
	}
	texts := strings.Fields(text)
	return texts[0] + " " + "TRUNCATED"
}

func Reverse(text string) string {
	runes := strings.Fields(text)
	for i := 0; i < len(runes); i++ {
		if runes[len(runes)-1] == "(Reverse)" {
			runes[len(runes)-1] = ""
			for i, j := 1, len(runes)-1; i < j; i, j = i+1, j-1 {
				//if runes[len(runes)-1] == "(Reverse)" {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return strings.Join(runes, " ")
}
