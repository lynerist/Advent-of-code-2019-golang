package main

import (
	"bufio"
	"fmt"
	"math"
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

	var manhattan int
	for i, c := range cross {
		distance := int(math.Abs(float64(c.x-centralPort.x))) + int(math.Abs(float64(c.y-centralPort.y)))
		if distance < manhattan || i == 0 {
			manhattan = distance
		}
	}
	fmt.Println(manhattan)
}

func track(w rune, cmd string, m [][]rune, start cell, cross []cell) (cell, []cell) {
	var i, j int
	switch cmd[0:1] {
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