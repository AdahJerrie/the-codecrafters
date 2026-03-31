package main

import (
	"fmt"
	"os"
	"strings"
)

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
	if text == " " || text == "-" {
		word, err := Dashes(text)
		if err != true {
			panic(err)
		}
		return word
	}

	return text
}

func processbylines(text string) string {
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		lines[i] = processText(line)
	}
	return strings.Join(lines, "\n")
}

func main() { //1. determine no. of aarguments
	//2. access our file through args
	if len(os.Args) != 3 {
		fmt.Println("Incomplete arguments")
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	data, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	result := processbylines(string(data))

	err = os.WriteFile(outputfile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file", err)
		return
	}
	fmt.Println("Printed successfully")

}
