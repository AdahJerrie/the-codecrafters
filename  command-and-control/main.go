package main

import (
	"bufio"
	baseconverter "command-and-control/base-converter"
	calculatorfunctions "command-and-control/calculator-functions"
	stringstransformer "command-and-control/strings-transformer"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Println("════════════════════════════════════════════════\nSENTINEL — COMMAND & CONTROL CONSOLE\nAll systems nominal. Type 'help' to begin.\n════════════════════════════════════════════════\nC&C>\nFrom this single prompt, the user can access\nall three tools using a module prefix:\ncalc   <command>   → the calculator\nbase   <command>   → the base converter\nstr    <command>   → the string transformer\nhelp               → shows all commands\nhistory            → shows last 10 inputs\nexit               → shuts down the console")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		sli := strings.Fields(input)

		indexofpipe := 0
		compoundCommands := false

		for i, val := range sli {
			if val == "|" {
				compoundCommands = true
				indexofpipe = i
				break
			}
		}

		if indexofpipe == 0 && compoundCommands == false {
			operation := sli[0]
			switch operation {
			case "calc":
				subCommand := sli[1]
				if strings.ToLower(subCommand) == "add" {
					if len(sli) < 4 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						input1, _ := strconv.Atoi(sli[2])
						input2, _ := strconv.Atoi(sli[3])
						result := calculatorfunctions.Add(input1, input2)
						fmt.Printf("%d + %d = %d\n", input1, input2, result)
					}
				} else if strings.ToLower(subCommand) == "sub" {
					if len(sli) < 4 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						input1, _ := strconv.Atoi(sli[2])
						input2, _ := strconv.Atoi(sli[3])
						result := calculatorfunctions.Sub(input1, input2)
						fmt.Printf("%d - %d = %d\n", input1, input2, result)
					}
				} else if strings.ToLower(subCommand) == "mult" {
					if len(sli) < 4 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						input1, _ := strconv.Atoi(sli[2])
						input2, _ := strconv.Atoi(sli[3])
						result := calculatorfunctions.Mult(input1, input2)
						fmt.Printf("%d * %d = %d\n", input1, input2, result)
					}
				} else if strings.ToLower(subCommand) == "div" {
					if len(sli) < 4 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						input1, _ := strconv.Atoi(sli[2])
						input2, _ := strconv.Atoi(sli[3])
						if input2 == 0 {
							fmt.Println("denominator cannot be zero")
							break
						} else {
							result := calculatorfunctions.Sub(input1, input2)
							fmt.Printf("%d - %d = %d\n", input1, input2, result)
						}

					}
				} else if strings.ToLower(subCommand) == "mod" {
					if len(sli) < 4 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						input1, _ := strconv.Atoi(sli[2])
						input2, _ := strconv.Atoi(sli[3])
						result := calculatorfunctions.Mod(input1, input2)
						fmt.Printf("%d %% %d = %d\n", input1, input2, result)
					}
				} else if strings.ToLower(subCommand) == "pow" {
					if len(sli) < 4 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						input1, _ := strconv.Atoi(sli[2])
						input2, _ := strconv.Atoi(sli[3])
						result := calculatorfunctions.Pow(input1, input2)
						fmt.Printf("%d ^ %d = %d\n", input1, input2, result)
					}
				}

			case "base":
				subCommand := sli[1]
				if strings.ToLower(subCommand) == "bin" {
					if len(sli) < 3 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						input := sli[2]
						result := baseconverter.BinTodecimal(input)
						fmt.Printf("binary: %s = Decimal: %d\n", input, result)
					}
				} else if strings.ToLower(subCommand) == "hex" {
					if len(sli) < 3 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						input := sli[2]
						result := baseconverter.HexTodecimal(input)
						fmt.Printf("Hexadecimal: %s = Decimal: %d\n", input, result)
					}
				} else if strings.ToLower(subCommand) == "dec" {
					if len(sli) < 3 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						input, _ := strconv.Atoi(sli[2])
						result := baseconverter.Dec(int64(input))
						fmt.Printf("Decimal: %d =  %s\n", input, result)
					}
				}
			case "str":
				subCommand := sli[1]
				if strings.ToLower(subCommand) == "upper" {
					if len(sli) < 3 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						newsli := sli[2:]
						str := strings.Join(newsli, " ")
						result := stringstransformer.Upper(str)
						fmt.Printf("Transformed string: %s\n", result)
					}
				} else if strings.ToLower(subCommand) == "cap" {
					if len(sli) < 3 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						newsli := sli[2:]
						str := strings.Join(newsli, " ")
						result := stringstransformer.Cap(str)
						fmt.Printf("Transformed string: %s\n", result)
					}
				} else if strings.ToLower(subCommand) == "lower" {
					if len(sli) < 3 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						newsli := sli[2:]
						str := strings.Join(newsli, " ")
						result := stringstransformer.Lower(str)
						fmt.Printf("Transformed string: %s\n", result)
					}
				} else if strings.ToLower(subCommand) == "title" {
					if len(sli) < 3 {
						fmt.Print("insufficient number of arguments")
						break
					} else {
						newsli := sli[2:]
						str := strings.Join(newsli, " ")
						result := stringstransformer.Title(str)
						fmt.Printf("Transformed string: %s\n", result)
					}
				}
			}
		}
	}
}
