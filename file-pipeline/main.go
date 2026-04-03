package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Number(len string) int {
	data, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(data)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func main() {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		log.Fatal("Invalid numnber of Arguments")
		//return
	}

	if os.Args[1] != "input.txt" {
		log.Fatal("Invalid argument name")
	}

	if os.Args[1] == os.Args[2] {
		log.Fatal("Input and output cannot be the same file.")
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	//os.stat is used to check if a file exists, if it doesnt os.!IsNotExist handles err if file doesnt exist

	if _, err := os.Stat(inputfile); os.IsNotExist(err) {
		log.Fatalf("%v does not exist", inputfile)
	}

	data, err := os.ReadFile(inputfile)
	if err != nil {
		log.Fatal(err)
		return
	}

	result := processbylines(string(data))

	err = os.WriteFile(outputfile, []byte(result), 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Summary block of printed file:")
	fmt.Printf("Number of lines read: %v\n", Number("input.txt"))
	fmt.Printf("Number of lines written: %v\n", Number("input.txt"))
	fmt.Println("THE RULES APPLIED ARE:")
	fmt.Println("1. Conversion to CAPS")
	fmt.Println("2. Conversion to lowercase")
	fmt.Println("3. Replace TODO with ACTION")
	fmt.Println("4. The string is reversed, for strings ending with (Reverse)")
	fmt.Println("5. Lines longer than 80 words return TRUNCATED")
}
