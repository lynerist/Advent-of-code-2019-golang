package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	var sum int
	var mass string
	sc := bufio.NewScanner(os.Stdin)
	for {
		sc.Scan()
		mass = sc.Text()
		if mass == ""{
			break
		}
		n, _ := strconv.Atoi(mass)
		sum += (n/3-2)
	}
	fmt.Println(sum)
}