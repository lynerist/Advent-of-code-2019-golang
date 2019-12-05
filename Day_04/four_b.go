package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	lowestBorder, _ := strconv.Atoi(sc.Text()[:6])
	highestBorder, _ := strconv.Atoi(sc.Text()[7:])

	var passwordCounter int
	for i := lowestBorder; i <= highestBorder; i++ {
		number := fmt.Sprint(i)

		var j int
		var adjacent bool
		for j = 0; j < 5; j++ {
			var matching bool
			if number[j] > number[j+1] {
				break
			}
			if number[j] == number[j+1] {
				matching = true
				if (j > 0 && number[j-1] == number[j]) || (j < 4 && number[j] == number[j+2]) {
					matching = false
				}
			}
			if matching {
				adjacent = matching
			}
		}

		if j < 5 || !adjacent {
			continue
		}

		passwordCounter++
	}
	fmt.Println(passwordCounter)
}