package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	y int
	x int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	wireA := strings.Split(sc.Text(), ",")
	sc.Scan()
	wireB := strings.Split(sc.Text(), ",")

	matrix := make([][]rune, 40001)
	for i := range matrix {
		matrix[i] = make([]rune, 40001)
	}

	centralPort := cell{20000, 20000}
	var cross []cell

	current := centralPort
	for _, s := range wireA {
		current, cross = track('a', s, matrix, current, cross)
	}

	current = centralPort
	for _, s := range wireB {
		current, cross = track('b', s, matrix, current, cross)
	}

	distances := make(map[cell]int)
	for i := range cross {
		distances[cross[i]] = 0
	}

	distances = misure(wireA, matrix, distances, centralPort)
	distances = misure(wireB, matrix, distances, centralPort)

	var lowestDistance int
	for i := range cross {
		if distances[cross[i]] < lowestDistance || i == 0 {
			lowestDistance = distances[cross[i]]
		}
	}
	fmt.Println(lowestDistance)
}

func direction(cmd string) (i, j int) {
	switch cmd {
	case "D":
		i = 1
		j = 0
	case "U":
		i = -1
		j = 0
	case "L":
		i = 0
		j = -1
	case "R":
		i = 0
		j = 1
	}
	return
}

func track(w rune, cmd string, m [][]rune, start cell, cross []cell) (cell, []cell) {
	i, j := direction(cmd[0:1])

	y := start.y
	x := start.x
	n, _ := strconv.Atoi(cmd[1:])

	for offset := 1; offset <= n; offset++ {
		nextY := y + offset*i
		nextX := x + offset*j
		if m[nextY][nextX] != 0 && m[nextY][nextX] != w && m[nextY][nextX] != 'x' {
			cross = append(cross, cell{y + offset*i, x + offset*j})
			m[nextY][nextX] = 'x'
		} else {
			m[nextY][nextX] = w
		}
	}
	return cell{y + n*i, x + n*j}, cross
}

func misure(cmds []string, m [][]rune, d map[cell]int, centralPort cell) map[cell]int {

	counted := make(map[cell]bool)
	current := centralPort
	var lengthWire int
	for _, cmd := range cmds {
		i, j := direction(cmd[0:1])
		n, _ := strconv.Atoi(cmd[1:])

		for offset := 1; offset <= n; offset++ {
			lengthWire++
			nextY := current.y + offset*i
			nextX := current.x + offset*j
			crossCell := cell{nextY, nextX}

			if m[nextY][nextX] == 'x' && !counted[crossCell] {
				d[crossCell] += lengthWire
				counted[crossCell] = true
			}
		}
		current = cell{current.y + n*i, current.x + n*j}
	}
	return d
}