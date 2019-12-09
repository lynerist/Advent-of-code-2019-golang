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
	fmt.Println(software(instructions, []int{2}))
}

//This function allocates the space for new elements when and index over the ranges is called
func memoryAllocation(address int, memory []int) []int {
	for address >= len(memory) {
		memory = append(memory, 0)
	}
	return memory
}

//This function gets the operators with the right modes
//it calls memoryAllocation when the number of elements could change
func getOperator(memory []int, relativeBase, index int, mode rune) (int, []int) {
	operator := memory[index]
	if mode == '0' {
		memory = memoryAllocation(operator, memory)
		operator = memory[operator]
	} else if mode == '2' {
		memory = memoryAllocation(operator+relativeBase, memory)
		operator = memory[operator+relativeBase]
	}
	return operator, memory
}

func software(program []int, input []int) int {
	instructions := make([]int, len(program))
	copy(instructions, program)

	var outputString string
	var programCounter, relativeBase int

	for programCounter > -1 {
		opCode := instructions[programCounter] % 100
		parameters := fmt.Sprintf("%.3d", instructions[programCounter]/100)

		//Here we get the operators that we need with the right modes
		var a, b, c int
		switch opCode {
		case 1, 2, 7, 8:
			a, instructions = getOperator(instructions, relativeBase, programCounter+1, rune(parameters[2]))
			b, instructions = getOperator(instructions, relativeBase, programCounter+2, rune(parameters[1]))
			//We can't get c with mode 0 because we want write in position instruction[c]
			//Getting it with mode 0 we could write in position instruction[instruction[c]]
			//due to the fact that getOperator with mode 0 give us the referenced value
			c, instructions = getOperator(instructions, relativeBase, programCounter+3, '1')
			//For the same reasons we must shift the value of c out the function in this way
			if parameters[0] == '2' {
				c += relativeBase
			}
			instructions = memoryAllocation(c, instructions)
		case 3:
			//The same worth for all the operators we need for write in the memory
			a, instructions = getOperator(instructions, relativeBase, programCounter+1, '1')
			if parameters[2] == '2' {
				a += relativeBase
			}
			instructions = memoryAllocation(a, instructions)
		case 5, 6:
			a, instructions = getOperator(instructions, relativeBase, programCounter+1, rune(parameters[2]))
			b, instructions = getOperator(instructions, relativeBase, programCounter+2, rune(parameters[1]))
			// we will jump at the address b so we have to control that it exists.
			instructions = memoryAllocation(b, instructions)
		case 4:
			a, instructions = getOperator(instructions, relativeBase, programCounter+1, rune(parameters[2]))
		case 9:
			a, instructions = getOperator(instructions, relativeBase, programCounter+1, rune(parameters[2]))
		}

		switch opCode {
		case 1: // Add
			instructions[c] = a + b
			programCounter += 4 //Number of instructions
		case 2: // Multiply
			instructions[c] = a * b
			programCounter += 4 //Number of instructions
		case 3: // 	Input
			instructions[a] = input[0]
			programCounter += 2 //Number of instructions
			input = input[1:]
		case 4: //	Output
			outputString += fmt.Sprint(a)
			programCounter += 2 //Number of instructions
		case 5: // Branch if Not Equal (jump)
			programCounter += 3
			if a != 0 {
				programCounter = b
			}
		case 6: // Branch if Equal (jump)
			programCounter += 3
			if a == 0 {
				programCounter = b
			}
		case 7: //Set less then
			if a < b {
				instructions[c] = 1
			} else {
				instructions[c] = 0
			}
			programCounter += 4 //Number of instructions
		case 8: //Set 1 on equal
			if a == b {
				instructions[c] = 1
			} else {
				instructions[c] = 0
			}
			programCounter += 4 //Number of instructions
		case 9: //Change relative base
			relativeBase += a
			programCounter += 2 //Number of instructions
		case 99: //Exit
			programCounter = -1
		default:
			fmt.Println("Error at instruction number ", programCounter)
		}
	}
	output, _ := strconv.Atoi(outputString)
	return output
}