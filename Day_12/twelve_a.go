package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// I used the float64 to use the func Abs in the calculation of the energy without casts
type moon struct {
	x, y, z    float64 //position
	vx, vy, vz float64 //velocity
	gx, gy, gz float64 //gravity
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var moons []moon

	// We read the input, you can do go run file.go < input.txt
	for {
		sc.Scan()
		if sc.Text() == "" {
			break
		}
		newMoon := moon{}
		fmt.Sscanf(sc.Text(), "<x=%f, y=%f, z=%f>", &newMoon.x, &newMoon.y, &newMoon.z)
		moons = append(moons, newMoon)
	}

	for step := 0; step < 1000; step++ {

		//We see all the couples of moons 1000 times
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
		}
	}

	//Finally we calculate the energy
	var totalEnergy float64
	for _, moon := range moons {
		potential := math.Abs(moon.x) + math.Abs(moon.y) + math.Abs(moon.z)
		kinetic := math.Abs(moon.vx) + math.Abs(moon.vy) + math.Abs(moon.vz)
		totalEnergy += potential * kinetic
	}
	fmt.Printf("%.0f", totalEnergy)
}
