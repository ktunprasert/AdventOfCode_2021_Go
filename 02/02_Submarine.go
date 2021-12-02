package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ktunprasert/AdventOfCode_2021_Go/helper"
)

func splitInstruction(line string) (string, int64) {
	split := strings.Split(line, " ")
	direction := split[0]
	unit, _ := strconv.Atoi(split[1])

	return direction, int64(unit)
}

func calculateFinalDepth(lines []string) (finalDepth int64) {
	var f, d, u int64
	f, d, u = 0, 0, 0

	for _, line := range lines {
		direction, unit := splitInstruction(line)

		if direction == "forward" {
			f += unit
		} else if direction == "up" {
			u += unit
		} else {
			d += unit
		}
	}
	finalDepth = int64(math.Abs(float64(f * (d - u))))

	return
}

func calculateAimDepth(lines []string) (finalDepth int64) {
	var aim, horizontal, depth int64
	aim, horizontal, depth = 0, 0, 0

	for _, line := range lines {
		direction, unit := splitInstruction(line)
		if direction == "forward" {
			horizontal += unit
			depth += unit * aim
		} else if direction == "down" {
			aim += unit
		} else {
			aim -= unit
		}
	}
	return horizontal * depth
}

func main() {
	big_test := helper.ReadFileAsStr("02.txt")
	out1 := calculateFinalDepth(big_test)
	out2 := calculateAimDepth(big_test)
	fmt.Println(out1, out2)
}
