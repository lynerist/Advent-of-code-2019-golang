package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// We get the input with a scanner (you can do "file.go < input.txt")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	slice := strings.Split(sc.Text(), ",")
	var instructions []int
	for _, s := range slice {
		n, _ := strconv.Atoi(s)
		instructions = append(instructions, n)
	}
	instructions = append(instructions, 0, 0)

	// The number of combinations of 5 elements is the factorial 5! so we need 5! elements
	const factorialNumberPhase = 5 * 4 * 3 * 2
	var combinations []string
	for i := 0; len(combinations) < factorialNumberPhase; i++ {
		combination := strconv.FormatInt(int64(i), 5)
		if len(combination) < 5 {
			combination = "0" + combination
		}

		// We check that every digit compares at least one time (so just one time)
		var countDigit [5]bool
		for _, digit := range combination {
			countDigit[digit-'0'] = true
		}
		var IsCombination bool = true
		for _, compare := range countDigit {
			if !compare {
				IsCombination = false
			}
		}
		if IsCombination {
			combinations = append(combinations, combination)
		}
	}

	// We try each combination
	var maxSignal int
	for _, combination := range combinations {
		// We need a software and a pc for each amplifier because it has to keep going without changes until the halt,
		// When we switch amplifier we just pause it's software (like a sleeping process in the OS)
		var amplifierSoftware [][]int
		for i := 0; i < 5; i++ {
			program := make([]int, len(instructions))
			copy(program, instructions)
			amplifierSoftware = append(amplifierSoftware, program)
		}
		programCounters := make([]int, 5)

		var signal int
		for loop := 0; programCounters[4] != -1; loop++ {
			//Here starts the feedback loop, at the first iteration we have to pass to the software also the phase,
			//otherwise we just have to pass him the current signal
			for amplifier, phase := range combination {
				if loop == 0 {
					// We need new phases so we add 5 and the range changes from [0-4] to [5-9]
					signal, programCounters[amplifier] = software(amplifierSoftware[amplifier], []int{int(phase - '0' + 5), signal}, programCounters[amplifier])
				} else {
					signal, programCounters[amplifier] = software(amplifierSoftware[amplifier], []int{signal}, programCounters[amplifier])
				}
			}
		}

		if signal > maxSignal {
			maxSignal = signal
		}
	}
	fmt.Println(maxSignal)
}

func software(instructions []int, input []int, programCounter int) (int, int) {

	var outputString string

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
			//If input is empty we have to stop the software and start the next amplifier
			if len(input) == 0 {
				output, _ := strconv.Atoi(outputString)
				return output, programCounter
			}
			
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
	return output, programCounter
}