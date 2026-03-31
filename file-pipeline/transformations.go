/*project file-pipeline
transformations - i. caps(title case)
				  ii. All lowercase(to lower all uppercase)
				  iii. lines starting with todo: replace the todo with ACTION
				  iv.Dashes/blanks(reove such lines)
				  v. flag any line longer than 80 words with truncated
*/

package main

import (
	"fmt"
	"slices"
	"strings"
)

func Cap(text string) string {
	connectors := []string{"an", "to", "the", "and", "but", "for", "or", "at", "to", "by", "on", "in", "of", "as", "is", "it"}
	//field := strings.ToLower(text)
	fields := strings.Fields(text)

	for i, val := range fields {
		if fields[len(fields)-1] == "(cap)" {
			myFlag := slices.Contains(connectors, val)
			if myFlag == false {
				fields[i] = strings.ToUpper(string(fields[i][0])) + strings.ToLower(fields[i][1:])
				fields = append(fields[0:len(fields)-1], fields[len(fields)-1+1:]...)
				i--
			}
		}

	}
	return strings.Join(fields, " ")
}

// func Upper(text string) string {
// 	word
// }

func main() {
	word := "everybody is running to win (cap)"
	fmt.Println(Cap(word))
}
