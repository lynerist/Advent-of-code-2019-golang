package main

import (
	"bufio"
	"fmt"
	"os"
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

	orbitsSanta := make(map[string]int)
	currentOrbit := orbits["SAN"]
	for steps := 0; currentOrbit != "COM"; steps++ {
		orbitsSanta[currentOrbit] = steps
		currentOrbit = orbits[currentOrbit]
	}

	var countOrbits int
	currentOrbit = orbits["YOU"]
	for steps := 0; ; steps++ {
		orbitsFromSanta, cross := orbitsSanta[currentOrbit]
		countOrbits = steps + orbitsFromSanta
		if cross {
			break
		}
		currentOrbit = orbits[currentOrbit]
	}
	fmt.Println(countOrbits)
}
