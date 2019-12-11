package main

import (
	"fmt"
	"os"
	"bufio"
	"math"
)

type asteroid struct{
	x, y float64
	m map[float64][2]bool
}

func main(){

	sc := bufio.NewScanner(os.Stdin)
	var asts [][]bool
	
	var asteroids []asteroid

	for{
		sc.Scan()
		var asteroidsRow []bool
		for i, ast := range sc.Text(){
			if ast == '#'{
				asteroidsRow = append(asteroidsRow, true)
				asteroids = append(asteroids, asteroid{float64(i),float64(len(asts)), make(map[float64][2]bool)})
			}else{
				asteroidsRow = append(asteroidsRow, false)
			}
		}

		if len(asteroidsRow) == 0{
			break
		}
		asts = append(asts, asteroidsRow)
	}

	for i, a := range asteroids{
		for j, as := range asteroids{
			if i == j{
				continue
			}

			var h int
			if a.y<as.y || (as.y == a.y && a.x<as.x){
				h = 1
			}

			if (as.x-a.x) == 0{
				temp := asteroids[i].m[float64(int(math.NaN()))]
				temp[h] = true
				asteroids[i].m[float64(int(math.NaN()))] = temp
			}else{
				temp := asteroids[i].m[(as.y-a.y)/(as.x-a.x)]
				temp[h] = true
				asteroids[i].m[(as.y-a.y)/(as.x-a.x)] = temp
			}
		}
	}

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
	fmt.Println("\n\n",maxMonitored)
}
