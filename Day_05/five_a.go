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
	var numberInstructions int

	for programCounter := 0 ; programCounter > -1; programCounter += numberInstructions{
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
			numberInstructions = 4
		case 2:
			instructions[instructions[programCounter+3]] = a * b
			numberInstructions = 4
		case 3:
			sc.Scan()
			instructions[instructions[programCounter+1]], _ = strconv.Atoi(sc.Text())
			numberInstructions = 2
		case 4:
			fmt.Println(instructions[instructions[programCounter+1]])
			numberInstructions = 2
		case 99:
			programCounter = -1
			numberInstructions = 0
		default:
			fmt.Println("Error at instruction number ", programCounter)
		}
	} 
}