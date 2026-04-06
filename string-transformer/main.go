package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
start:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text to transform e.g 'upper <text>' or 'quit' to exit")
	scanner.Scan()
	input := scanner.Text()

	result := ProcessText(input)
	fmt.Println(result)
	goto start
}
