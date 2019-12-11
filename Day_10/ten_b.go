package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//We need coordinates and a map from angular coefficient to number of asteroids seen on this stright line
//(At most two so we can use a boolean array with len 2)
type asteroid struct {
	x, y float64
	m    map[float64][2]bool
}

func main() {
	// A value arbitrary that won't be an avaible coefficient to rappresent value NaN
	// for the stright line where a.x-as.x == 0
	const nan = 101010101010101.0

	//We store the input, you can do go run program.go < input.txt
	sc := bufio.NewScanner(os.Stdin)
	var asteroids []asteroid
	for i := 0; sc.Text() != "" || i == 0; i++ {
		sc.Scan()
		for j, ast := range sc.Text() {
			if ast == '#' {
				asteroids = append(asteroids, asteroid{float64(j), float64(i), make(map[float64][2]bool)})
			}
		}
	}

	for i, a := range asteroids {
		for j, as := range asteroids {
			//If the asteroid is the same we don't have to do anything
			if i == j {
				continue
			}
			//We see if the asteroid is in a lower or upper position then the current one
			var direction int
			if a.y < as.y || (as.y == a.y && a.x < as.x) {
				direction = 1
			}

			StraightLineCoefficent := nan
			if (as.x - a.x) != 0 {
				StraightLineCoefficent = (as.y - a.y) / (as.x - a.x)
			}
			// We have to assign the values to the array and after we can set is as value in the map
			seenOnThisStraightLine := asteroids[i].m[StraightLineCoefficent]
			seenOnThisStraightLine[direction] = true
			asteroids[i].m[StraightLineCoefficent] = seenOnThisStraightLine
		}
	}

	//Here we count the asteroids seen from each asteroid
	var maxMonitored int
	var station asteroid
	for _, a := range asteroids {
		var countMonitored int
		for _, v := range a.m {
			if v[0] {
				countMonitored++
			}
			if v[1] {
				countMonitored++
			}
		}
		if countMonitored > maxMonitored {
			maxMonitored = countMonitored
			station = a
		}
	}

	type toDestroy struct {
		x, y     float64
		ac       float64 //Angular coefficient
		distance float64
	}

	var monitoreds []toDestroy
	// We divide it in two slides because the asteroids on the same straight line have the same angular coefficient
	//But we want to destroy before the ones on the right side.
	var leftSide, rightSide []toDestroy

	for _, a := range asteroids {
		//We don't want to destroy the station...
		if a.x == station.x && a.y == station.y {
			continue
		}

		monitored := toDestroy{x: a.x, y: a.y, ac: nan}

		if (a.x - station.x) != 0 {
			monitored.ac = (station.y - a.y) / (a.x - station.x)
		}
		//We'll use the Manhattan distance to see wich asteroid on the same half line has to been destroied before
		monitored.distance = math.Abs(station.x-a.x) + math.Abs(station.y-a.y)
		//Here we assign the asteroids to the two halfs
		if a.x > station.x || (a.x == station.x && a.y < station.y) {
			rightSide = append(rightSide, monitored)
		} else {
			leftSide = append(leftSide, monitored)
		}

	}
	// We sort both sides considering the angular coefficients. (The highest ones will be destroyed before)
	// When the ac is the same we consider the distance
	for i := range leftSide {
		for j := range leftSide {
			if leftSide[i].ac > leftSide[j].ac {
				leftSide[i], leftSide[j] = leftSide[j], leftSide[i]
			} else if leftSide[i].ac == leftSide[j].ac && leftSide[i].distance < leftSide[j].distance {
				leftSide[i], leftSide[j] = leftSide[j], leftSide[i]
			}
		}
	}

	for i := range rightSide {
		for j := range rightSide {
			if rightSide[i].ac > rightSide[j].ac {
				rightSide[i], rightSide[j] = rightSide[j], rightSide[i]
			} else if rightSide[i].ac == rightSide[j].ac && rightSide[i].distance < rightSide[j].distance {
				rightSide[i], rightSide[j] = rightSide[j], rightSide[i]
			}
		}
	}

	//Finally we merge the sides
	monitoreds = append(rightSide, leftSide...)

	for {
		inOrder := true
		precedent := 0.0
		var moved bool
		var monitoredsInOrder, toPutToTheBottom []toDestroy

		//after this one we don't have to sort anymore
		last := monitoreds[len(monitoreds)-1]

		for i, a := range monitoreds {
			if a == last {
				break
			}
			if a.ac != precedent || i == 0 {
				monitoredsInOrder = append(monitoredsInOrder, a)
				moved = false
			} else if moved {
				monitoredsInOrder = append(monitoredsInOrder, a)
			} else {
				//We pick one asteroid with the ac == precedent for each ac that has it
				//And we put it to the bottom of the slice
				toPutToTheBottom = append(toPutToTheBottom, a)
				moved = true
				inOrder = false
			}
			precedent = a.ac
		}
		monitoreds = append(monitoredsInOrder, toPutToTheBottom...)
		//We have sorted all
		if inOrder {
			break
		}
	}
	//Finally the 200th asteroid
	fmt.Printf("%v%2.0f", monitoreds[199].x, monitoreds[199].y)
}
