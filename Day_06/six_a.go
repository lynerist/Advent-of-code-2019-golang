package main

import(
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main(){
	sc := bufio.NewScanner(os.Stdin)
	
	var input []string

	for {
		sc.Scan()
		if sc.Text() == ""{
			break
		}
		input = append(input, sc.Text())
	}

	orbits := make(map[string]string)

	for _, orbit := range input{
		mass := orbit[:3]
		moon := orbit[4:]
		orbits[moon] = mass
	}
	
	var countOrbits int

	for k := range orbits{
		var countOrbitsK int
		currentOrbit := k
		for {
			if currentOrbit == "COM"{
				break
			}
			if len(currentOrbit) == 4{
				orbitsAlreadyCounted, _ := strconv.Atoi(currentOrbit)
				countOrbitsK += orbitsAlreadyCounted
				break
			}
			countOrbitsK++
			currentOrbit = orbits[currentOrbit]
		}
		countOrbits += countOrbitsK
		if  currentOrbit != "COM"{
			orbits[k] = fmt.Sprintf("%.4d", countOrbitsK)
		}
	}
	fmt.Println(countOrbits)
}
