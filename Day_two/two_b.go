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

	var input []int
	for _, s := range slice {
		n, _ := strconv.Atoi(s)
		input = append(input, n)
	}

	for j := 0; j < 99; j++ {
		for k := 0; k < 99; k++ {
			num := make([]int, len(input))
			copy(num, input)
			num[1] = j
			num[2] = k

			for i := 0; ; i += 4 {
				a := num[num[i+1]]
				b := num[num[i+2]]
				if num[i] == 1 {
					num[num[i+3]] = a + b
				} else if num[i] == 2 {
					num[num[i+3]] = a * b
				} else if num[i] == 99 {
					break
				} else {
					fmt.Println("ERROR")
					fmt.Println(i)
					break
				}
				if num[0] > 19690720 {
					continue
				} else if num[0] == 19690720 {
					fmt.Print(j, k)
					return
				}
			}
		}
	}
}
