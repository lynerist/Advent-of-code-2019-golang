package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	var sum int
	var riga string = "start"
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

func calculate(n int)int{
	n/=3
	n-=2
	if n<=0{
		return 0
	}
	return n + calculate(n)
	
}