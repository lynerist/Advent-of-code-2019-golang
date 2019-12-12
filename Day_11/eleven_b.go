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
	software(instructions)
}

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

// ALL LIKE IN DAY NINE, I JUST CHANGED THE COMMENTED PARTS
func software(instructions []int) {
	//We need a grid where the robot will move
	grid := make([][]rune, 21)
	for i := range grid {
		grid[i] = make([]rune, 81)
	}
	position := [2]int{30, 10}
	grid[position[1]][position[0]] = '1'

	// We'll use 0 = N 1 = E 2 = S 3 = O
	const N, E, S, O = 0, 1, 2, 3
	var direction int = N

	var outputString string
	var programCounter, relativeBase int

	for programCounter > -1 {
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
			// We read the color (if we haven't painted it yet its value will be 0, and -'0' will be negative)
			color := int(grid[position[1]][position[0]] - '0')
			if color < 0 {
				color = 0
			}
			instructions[a] = color
			programCounter += 2 //Number of instructions
		case 4: //	Output
			//We cumulate both outputs before to paint and move
			outputString += fmt.Sprint(a)

			if len(outputString) > 1 {
				//Here we PAINT
				grid[position[1]][position[0]] = rune(outputString[0])
				//Here we read the new direction
				if rune(outputString[1]) == '1' {
					direction++
				} else {
					direction--
				}
				//(Direction must be in range [0-3])
				direction += 4
				direction %= 4
				//We move (y is inverted because we are moving in a matrix)
				switch direction {
				case N:
					position[1]--
				case E:
					position[0]++
				case S:
					position[1]++
				case O:
					position[0]--
				}
				//We reset the output
				outputString = ""
			}

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

	//And finally here we print the code
	for i := range grid {
		for _, r := range grid[i] {
			if r == 49 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
