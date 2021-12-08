package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ktunprasert/AdventOfCode_2021_Go/helper"
)

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func recurSum(x int) int {
	if x < 0 {
		return 0
	}
	return x + recurSum(x-1)
}

func solveCrabs(lines []string) (total int) {
	var crabMaps map[int]int = make(map[int]int)
	var max int
	for _, v := range strings.Split(lines[0], ",") {
		v_int, _ := strconv.Atoi(string(v))
		crabMaps[v_int] += 1
		if v_int > max {
			max = v_int
		}
	}

	var idx int
	var min_fuel int = math.MaxInt
	var min_scaled_fuel int = math.MaxInt
	for idx <= max {
		var temp_fuel int
		var temp_scaled_fuel int
		for key, value := range crabMaps {
			abs := absDiffInt(key, idx)
			temp_fuel += abs * value
			temp_scaled_fuel += recurSum(abs) * value
		}

		if temp_fuel < min_fuel {
			min_fuel = temp_fuel
		}

		if temp_scaled_fuel < min_scaled_fuel {
			min_scaled_fuel = temp_scaled_fuel
		}

		idx++
	}

	fmt.Println(min_fuel, min_scaled_fuel)
	return total
}

func main() {
	// test_case := helper.ReadFileAsStr("07_test.txt")
	// solveCrabs(test_case)

	test1 := helper.ReadFileAsStr("07.txt")
	solveCrabs(test1)
}
