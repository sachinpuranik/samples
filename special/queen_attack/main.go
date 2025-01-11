package main

// pos := []string{"{1,2}","{4,5}", "{3,4}"}
// taking above array as input , assuming each position in array where queen is placed, check if the other queens are on its attack rout or not.

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getPos(pos string) (x, y int) {
	pos = strings.Trim(pos, " ")
	strippedPos := pos[1 : len(pos)-1]
	split := strings.Split(strippedPos, ",")
	x, _ = strconv.Atoi(strings.Trim(split[0], " "))
	y, _ = strconv.Atoi(strings.Trim(split[1], " "))
	return
}

func checkQueenAttackTrajectory(positions []string) bool {
	for i := 0; i < len(positions); i++ {
		for j := 0; j < len(positions); j++ {
			if i == j {
				continue
			}
			x1, y1 := getPos(positions[i])
			x2, y2 := getPos(positions[j])

			if x1 == x2 || y1 == y2 || math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)) {
				return true
			}
		}
	}
	return false
}

func main() {
	positions := []string{"(0, 0)", "(1, 4)", " (2, 7)", "(3, 5)", " (4, 2)", " (5, 6)", "(6, 1)", "(7, 3)"}
	ret := checkQueenAttackTrajectory(positions)
	fmt.Println("ans :", ret)
	positions = []string{"(0, 0)", "(1, 4)", " (2, 7)", "(3, 5)", " (4, 2)", " (5, 6)", "(6, 2)", "(7, 3)"}
	ret = checkQueenAttackTrajectory(positions)
	fmt.Println("ans :", ret)
}
