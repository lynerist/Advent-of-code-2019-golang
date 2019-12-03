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
	
	var num []int
	for _, s := range slice {
		n, _ := strconv.Atoi(s)
		num = append(num, n)
	}
	num[1] = 12
	num[2] = 2
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
	}
	fmt.Println(num[0])
}
