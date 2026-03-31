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

func Dashes(text string) (string, bool) {
	clean := strings.TrimSpace(text)
	if clean == " " || strings.Trim(clean, "-") == "" {
		return "", true
	}
	return text, false
}

// func main() {
// 	word := " "
// 	fmt.Println(Dashes(word))
// }
