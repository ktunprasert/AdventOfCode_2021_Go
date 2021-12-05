package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ktunprasert/AdventOfCode_2021_Go/helper"
)

type Grid struct {
	grid map[string]int
}

func (g *Grid) addCoords(xy string) {
	g.grid[xy] += 1
}

func makeRange(rangeStr string, skipDiag bool) (ranges []string) {
	start_X, start_Y := -1, -1
	end_X, end_Y := -1, -1

	for _, v := range strings.Split(rangeStr, " -> ") {
		for _, token := range strings.Split(v, ",") {
			token_int, _ := strconv.Atoi(token)

			if start_X < 0 {
				start_X = token_int
				continue
			}
			if start_Y < 0 {
				start_Y = token_int
				continue
			}
			if end_X < 0 {
				end_X = token_int
				continue
			}
			if end_Y < 0 {
				end_Y = token_int
				continue
			}
		}

	}

	if skipDiag {
		if int(math.Abs(float64(end_X)-float64(start_X))) == int(math.Abs(float64(end_Y)-float64(start_Y))) {
			return
		}
	}

	rangeVar := math.Max((math.Abs(float64(end_X) - float64(start_X))), (math.Abs(float64(end_Y) - float64(start_Y))))
	x_inc, y_inc := false, false
	if (end_X - start_X) != 0 {
		x_inc = true
	}
	if (end_Y - start_Y) != 0 {
		y_inc = true
	}

	for i := 0; i <= int(rangeVar); i++ {
		temp_i := i

		x_str := start_X
		y_str := start_Y

		if end_X < start_X {
			temp_i *= -1
		}

		if x_inc {
			x_str += temp_i
		}

		temp_i = i

		if end_Y < start_Y {
			temp_i *= -1
		}

		if y_inc {
			y_str += temp_i
		}

		ranges = append(ranges, fmt.Sprintf("%d,%d", x_str, y_str))
	}

	return
}

func calculateOverlap(lines []string, skipDiag bool) (total int) {
	g := Grid{make(map[string]int)}

	for _, line := range lines {
		ranges := makeRange(line, skipDiag)
		if len(ranges) > 0 {
			for _, key := range ranges {
				g.addCoords(key)
			}
		}
	}

	for _, val := range g.grid {
		if val > 1 {
			total += 1
		}
	}

	fmt.Println(total)
	return
}

func main() {
	// test_case := helper.ReadFileAsStr("05_test.txt")
	// calculateOverlap(test_case, true)
	// calculateOverlap(test_case, false)
	big_test := helper.ReadFileAsStr("05.txt")
	calculateOverlap(big_test, true)
	calculateOverlap(big_test, false)
}
