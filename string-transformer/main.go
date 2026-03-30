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
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Upper(text string) string {
	word := strings.ToUpper(text)
	return word
}

func ToLower(text string) string {
	word := strings.ToLower(text)
	return word
}

func Caps(text string) string {
	low := strings.ToLower(text)
	fields := strings.Title(low)

	return fields
}

func Titlecase(text string) string {
	fields := strings.Fields(text)

	for i := 0; i < len(fields); i++ {

	}
}

func main() {
start:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text to be edited or type 'quit' to exit")
	if !scanner.Scan() {
		fmt.Println("Try again")
		goto start
	}

	lines := scanner.Text()
	//input := strings.Fields(lines)
	//fmt.Println(input)

	result := Caps(lines)
	fmt.Println(result)
}
