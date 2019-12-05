package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func calculate(n int)int{
	n = n/3 -2
	if n<=0{
		return 0
	}
	return n + calculate(n)
}

func main(){
	var sum int
	var riga string
	sc := bufio.NewScanner(os.Stdin)
	for {
		sc.Scan()
		riga = sc.Text()
		if riga == ""{
			break
		}
		n, _ := strconv.Atoi(riga)
		n = calculate(n)
		sum+=n
	}
	fmt.Println(sum)
}