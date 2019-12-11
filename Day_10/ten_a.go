package main

import (
	"fmt"
	"os"
	"bufio"
)

//We need coordinates and a map from angular coefficient to number of asteroids seen on this stright line
//(At most two so we can use a boolean array with len 2)
type asteroid struct{
	x, y float64
	m map[float64][2]bool
}

func main(){

	//We store the input, you can do go run program.go < input.txt
	sc := bufio.NewScanner(os.Stdin)
	var asteroids []asteroid
	for i:=0; sc.Text() != "" || i == 0; i++{
		sc.Scan()
		for j, ast := range sc.Text(){
			if ast == '#'{
				asteroids = append(asteroids, asteroid{float64(j),float64(i), make(map[float64][2]bool)})
			}
		}
	}

	for i, a := range asteroids{
		for j, as := range asteroids{
			//If the asteroid is the same we don't have to do anything
			if i == j{
				continue
			}
			//We see if the asteroid is in a lower or upper position then the current one
			var direction int
			if a.y<as.y || (as.y == a.y && a.x<as.x){
				direction = 1
			}

			// A value arbitrary to rappresent the NaN of the stright line without coefficent (where a.x-as.x == 0)
			StraightLineCoefficent := 0.0101010101010101 

			if (as.x-a.x) != 0{
				StraightLineCoefficent = (as.y-a.y)/(as.x-a.x)
			}
			seenOnThisStraightLine := asteroids[i].m[StraightLineCoefficent]
			seenOnThisStraightLine[direction] = true
			asteroids[i].m[StraightLineCoefficent] = seenOnThisStraightLine
		}
	}

	//Here we count the asteroids seen from each asteroid
	var maxMonitored int
	for _, a := range asteroids{
		var countMonitored int
		for _, v:= range a.m{
			if v[0]{
				countMonitored++
			}
			if v[1]{
				countMonitored++
			}
		}
		if countMonitored>maxMonitored{
			maxMonitored = countMonitored
		}
	}
	fmt.Println(maxMonitored)
}
