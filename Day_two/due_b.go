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

	sc.Scan()
	lin := sc.Text()
	slice := strings.Split(lin, ",")

	var numc []int
	for _, s:= range slice{
		n, _ := strconv.Atoi(s)
		numc = append(numc, n)
	}

	for j :=0 ; j<99; j++{
		for k := 0; k<99; k++{
			num := make([]int, len(numc))
			copy(num, numc)
			num[1] = j 
			num[2] = k

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
				if num[0] > 19690720{
					continue
				}else if num[0] == 19690720{
					fmt.Print(j,k)
					return
				}
			}
		}
	}
}