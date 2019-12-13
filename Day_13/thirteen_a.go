package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"strings"
)

func main(){

	// We get the input with a scanner (you can do "file.go < input.txt")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	slice := strings.Split(sc.Text(), ",")
	var instructions []int
	for _, s := range slice {
		n, _ := strconv.Atoi(s)
		instructions = append(instructions, n)
	}
	//We run the software
	fmt.Println(software(instructions,nil))

}

//************* INTCODE UTILITIES	**********************	

func memoryAllocation(address int, memory []int) []int {
	for address >= len(memory) {
		memory = append(memory, 0)
	}
	return memory
}

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

//************* END INTCODE UTILITIES	**********************	

func software(instructions []int, input []int) int {

	var output []int
	var countBlocks int

	var programCounter, relativeBase int // INTCODE UTILITIES

	for programCounter > -1 {
	//************* INTCODE UTILITIES	**********************	
		opCode := instructions[programCounter] % 100
		parameters := fmt.Sprintf("%.3d", instructions[programCounter]/100)
		var a, b, c int
		switch opCode {
		case 1, 2, 7, 8:
			a, instructions = getOperator(instructions, relativeBase, programCounter+1, rune(parameters[2]))
			b, instructions = getOperator(instructions, relativeBase, programCounter+2, rune(parameters[1]))
			c, instructions = getOperator(instructions, relativeBase, programCounter+3, '1')
			if parameters[0] == '2' {
				c += relativeBase
			}
			instructions = memoryAllocation(c, instructions)
		case 3:
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
			programCounter += 4 
		case 2: // Multiply
			instructions[c] = a * b
			programCounter += 4 
		case 3: // 	Input
			instructions[a] = input[0]
			programCounter += 2
			input = input[1:]
//************* END INTCODE UTILITIES	**********************	

		case 4: //	HERE WE COUNT THE OUTPUTS == BLOCK
			output = append(output, a)
			if len(output)>2{
				if output[2] == 2{
					countBlocks++
				}
				output = nil
			}
			programCounter += 2 //Number of instructions

//************* INTCODE UTILITIES	**********************	
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
	//************* END INTCODE UTILITIES	**********************	

	return countBlocks
}