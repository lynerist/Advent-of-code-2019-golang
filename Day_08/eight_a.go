package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	// We get the input with a scanner (you can do "go run file.go < file.txt")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	image := sc.Text()

	const numberOfPixel = 25 * 6
	var fewestZero, fewestZeroOutput int
	var countZero, countOne, countTwo int

	//We count every digit and when the layer ends we control if it has the fewest number of zeros
	for index, digit := range image{
		if index % numberOfPixel == 0{
			//if it has, or if it is the first layer that we see, we change the outputs counters
			if countZero < fewestZero || index == numberOfPixel{
				fewestZero = countZero
				fewestZeroOutput = countOne * countTwo
			}
			countZero, countOne, countTwo = 0, 0, 0
		}
		switch digit{
		case '0':
			countZero++
		case '1':
			countOne++
		case '2':
			countTwo++
		}
	}
	fmt.Println(fewestZeroOutput)
}