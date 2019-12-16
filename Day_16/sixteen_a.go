package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	sequence := sc.Text()

	pattern := []int{0, 1, 0, -1}
	for i:=1; i<=100; i++{
		var sequenceNew string
		for j := 1 ; j<=len(sequence); j++{
			sum := 0
			for k, n := range sequence{
				sum += int(n-'0') * pattern[((k+1)/j) % 4]
			}
			s := fmt.Sprint(sum)
			sequenceNew += s[len(s)-1:]
		}
		sequence = sequenceNew
	}
	fmt.Println(sequence[0:8])
}