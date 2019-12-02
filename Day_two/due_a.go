package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"strings"
)

func main(){

	sc := bufio.NewScanner(os.Stdin)

	var num []int


	sc.Scan()
	lin := sc.Text()
	slice := strings.Split(lin, ",")

	for _, s:= range slice{
		n, _ := strconv.Atoi(s)
		num = append(num, n)
	}

	for i := 0; ; i+=4 {
		a := num[num[i+1]]
		b := num[num[i+2]]
		if num[i] == 1{
			num[num[i+3]] = a+b
		}else if num[i] == 2 {
			num[num[i+3]] = a*b
		}else if num[i] == 99{
			break
		}else{
			fmt.Println("ERRORE    ")
			fmt.Println(i)
			break
		}
	}
	fmt.Println(num[0])

}