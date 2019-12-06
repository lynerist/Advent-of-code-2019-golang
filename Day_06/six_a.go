package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var input []string
	for {
		sc.Scan()
		if sc.Text() == "" {
			break
		}
		input = append(input, sc.Text())
	}

	orbits := make(map[string]string)
	for _, orbit := range input {
		mass := orbit[:3]
		moon := orbit[4:]
		orbits[moon] = mass
	}

	var countOrbits int

	for k := range orbits {
		var countOrbitsK int
		currentOrbit := k
		for {
			if currentOrbit == "COM" {
				break
			}
			if len(currentOrbit) == 4 { // if the len of the string is 4 I know the steps I need to go to COM
				orbitsAlreadyCounted, _ := strconv.Atoi(currentOrbit)
				countOrbitsK += orbitsAlreadyCounted
				break
			}
			countOrbitsK++
			currentOrbit = orbits[currentOrbit]
		}
		countOrbits += countOrbitsK

		//I change the value orbits[k] with the number of steps needed to go to COM written in a string with len = 4
		if currentOrbit != "COM" {
			orbits[k] = fmt.Sprintf("%.4d", countOrbitsK) 
		}
	}
	fmt.Println(countOrbits)
}
