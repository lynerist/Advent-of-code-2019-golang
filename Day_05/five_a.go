package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
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
	var programCounter int

	for programCounter > -1 {
		opCode := instructions[programCounter]%100
		parameters := fmt.Sprintf("%.2d", instructions[programCounter]/100)

		a := instructions[programCounter+1]
		if parameters[1:] == "0" && (opCode == 1 || opCode == 2){
			a = instructions[a]
		}
		b := instructions[programCounter+2]
		if parameters[:1] == "0" && (opCode == 1 || opCode == 2){
			b = instructions[b]
		}

		switch opCode{
		case 1:
			instructions[instructions[programCounter+3]] = a + b
			programCounter += 4	//Number of instructions
		case 2:
			instructions[instructions[programCounter+3]] = a * b
			programCounter += 4	//Number of instructions
		case 3:
			sc.Scan()
			instructions[instructions[programCounter+1]], _ = strconv.Atoi(sc.Text())
			programCounter += 2	//Number of instructions
		case 4:
			fmt.Println(instructions[instructions[programCounter+1]])
			programCounter += 2	//Number of instructions
		case 99:
			programCounter = -1
		default:
			fmt.Println("Error at instruction number ", programCounter)
		}
	} 
}