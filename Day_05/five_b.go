package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	slice := strings.Split(sc.Text(), ",")

	var instructions []int
	for _, s := range slice {
		n, _ := strconv.Atoi(s)
		instructions = append(instructions, n)
	}

	instructions = append(instructions, 0, 0)
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
			sc.Scan()
			instructions[instructions[programCounter+1]], _ = strconv.Atoi(sc.Text())
			programCounter += 2 //Number of instructions
		case 4:
			fmt.Println(instructions[a])
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
}