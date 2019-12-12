package main

import (
	"bufio"
	"fmt"
	"os"
)

type moon struct {
	x, y, z    int //position
	vx, vy, vz int //velocity
	gx, gy, gz int //gravity
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var moons []moon
	var initialX, initialY, initialZ string
	// We read the input, you can do go run file.go < input.txt
	for {
		sc.Scan()
		if sc.Text() == "" {
			break
		}
		newMoon := moon{}
		fmt.Sscanf(sc.Text(), "<x=%d, y=%d, z=%d>", &newMoon.x, &newMoon.y, &newMoon.z)
		moons = append(moons, newMoon)
		//We save the first position of each moon in each dimension
		initialX += fmt.Sprint(newMoon.x)
		initialY += fmt.Sprint(newMoon.y)
		initialZ += fmt.Sprint(newMoon.z)
	}

	var frequenceX, frequenceY, frequenceZ int
	//After zero steps is the initial state and after one step there is the first change
	//For this reasons We have to start from the second one
	for step := 2; ; step++ {
		var stateX, stateY, stateZ string
		for i := 0; i < len(moons); i++ {
			for j := i; j < len(moons); j++ {
				// Here we calculate the gravity boost for each moon
				if moons[i].x < moons[j].x {
					moons[i].gx++
					moons[j].gx--
				} else if moons[i].x > moons[j].x {
					moons[i].gx--
					moons[j].gx++
				}
				if moons[i].y < moons[j].y {
					moons[i].gy++
					moons[j].gy--
				} else if moons[i].y > moons[j].y {
					moons[i].gy--
					moons[j].gy++
				}
				if moons[i].z < moons[j].z {
					moons[i].gz++
					moons[j].gz--
				} else if moons[i].z > moons[j].z {
					moons[i].gz--
					moons[j].gz++
				}
			}
			//Here we sum all the accelerations to the velocities
			moons[i].vx += moons[i].gx
			moons[i].vy += moons[i].gy
			moons[i].vz += moons[i].gz
			//We reset the gravities boosts
			moons[i].gx, moons[i].gy, moons[i].gz = 0, 0, 0
			//We sum all the velocities to the positions
			moons[i].x += moons[i].vx
			moons[i].y += moons[i].vy
			moons[i].z += moons[i].vz
			//We keep the state of each moon in each dimension
			stateX += fmt.Sprint(moons[i].x)
			stateY += fmt.Sprint(moons[i].y)
			stateZ += fmt.Sprint(moons[i].z)
		}
		//We want to find the frequence of allineation of the moons for each dimension
		//Found it we'll find the first time when all allineations are present
		if stateX == initialX && frequenceX == 0 {
			frequenceX = step
		}
		if stateY == initialY && frequenceY == 0 {
			frequenceY = step
		}
		if stateZ == initialZ && frequenceZ == 0 {
			frequenceZ = step
		}
		//I found all three frequencies
		if frequenceX != 0 && frequenceY != 0 && frequenceZ != 0 {
			break
		}
	}

	//Now I have to find the mcm of the frequences to see when all three conditions are verified
	//I find the highest frequence and I use it to find with the lowest number of sums
	// It may take a few seconds
	maxFrequence := frequenceX
	if frequenceY > maxFrequence {
		maxFrequence = frequenceY
	}
	if frequenceZ > maxFrequence {
		maxFrequence = frequenceZ
	}
	var mcmFrequences int = maxFrequence

	for mcmFrequences%frequenceX != 0 || mcmFrequences%frequenceY != 0 || mcmFrequences%frequenceZ != 0 {
		mcmFrequences += maxFrequence
	}
	fmt.Println(mcmFrequences)
}
