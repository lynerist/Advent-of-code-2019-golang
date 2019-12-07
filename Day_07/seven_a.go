package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	slice := strings.Split(sc.Text(), ",")
	var instructions []int
	for _, s := range slice {
		n, _ := strconv.Atoi(s)
		instructions = append(instructions, n)
	}
	instructions = append(instructions, 0, 0)

	const factorialNumberPhase = 5*4*3*2
	var combinations []string
	for i:=0; len(combinations)<factorialNumberPhase; i++{
		combination := strconv.FormatInt(int64(i), 5)
		if len(combination) < 5{
			combination = "0" + combination
		}

		var countDigit [5]bool
		for _, digit := range combination{	
			countDigit[digit-'0'] = true
		}
		var IsCombination bool = true
		for _, compare := range countDigit{
			if !compare{
				IsCombination = false
			}
		}
		if IsCombination{
			combinations = append(combinations, combination)
		}
	}

	var maxSignal int
	for _, combination := range combinations{

		var signal int
		for _, phase := range combination{
			signal = software(instructions, []int{int(phase-'0'), signal})
		}

		if signal>maxSignal{
			maxSignal = signal
		}
	}
	fmt.Println(maxSignal)
}

func software(program []int, input []int)int{
	instructions := make ([]int, len(program))
	copy(instructions, program)

	var outputString string
	var programCounter int

	for programCounter > -1 {
		opCode := instructions[programCounter] % 100
		parameters := fmt.Sprintf("%.2d", instructions[programCounter]/100)

		a := instructions[programCounter+1]
		b := instructions[programCounter+2]
		if opCode != 3 && opCode != 4 && opCode != 99 {
			if parameters[1:] == "0" {
				a = instructions[a]
			}
			if parameters[:1] == "0" {
				b = instructions[b]
			}
		}

		switch opCode {
		case 1:
			instructions[instructions[programCounter+3]] = a + b
			programCounter += 4 //Number of instructions
		case 2:
			instructions[instructions[programCounter+3]] = a * b
			programCounter += 4 //Number of instructions
		case 3:
			instructions[instructions[programCounter+1]] = input[0]
			programCounter += 2 //Number of instructions
			input = input[1:]
		case 4:
			outputString += fmt.Sprint(instructions[a])
			programCounter += 2 //Number of instructions
		case 5:
			programCounter += 3
			if a != 0 {
				programCounter = b //Jump
			}
		case 6:
			programCounter += 3
			if a == 0 {
				programCounter = b //Jump
			}
		case 7:
			if a < b {
				instructions[instructions[programCounter+3]] = 1
			} else {
				instructions[instructions[programCounter+3]] = 0
			}
			programCounter += 4 //Number of instructions
		case 8:
			if a == b {
				instructions[instructions[programCounter+3]] = 1
			} else {
				instructions[instructions[programCounter+3]] = 0
			}
			programCounter += 4 //Number of instructions
		case 99:
			programCounter = -1 //Exit
		default:
			fmt.Println("Error at instruction number ", programCounter)
		}
	}
	output, _ := strconv.Atoi(outputString)
	return output
}