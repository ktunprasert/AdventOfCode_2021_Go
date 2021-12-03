package main

import (
	"fmt"
	"strconv"

	"github.com/ktunprasert/AdventOfCode_2021_Go/helper"
)

func convertRuneToInt(r rune) int {
	return int(r - '0')
}

func linesToBits(lines []string) []int {
	var bits []int = make([]int, len(lines[0]))
	for _, v := range lines {
		for i, bit_rune := range v {
			bit := convertRuneToInt(bit_rune)
			if bit > 0 {
				bits[i] += 1
			} else {
				bits[i] -= 1
			}
		}
	}
	return bits
}

func calculatePowerConsumption(lines []string) (powerConsumption int) {
	numBits := linesToBits(lines)

	var gamma, epsilon int = 0, 0
	for _, bit := range numBits {
		if bit > 0 {
			gamma |= 1
			epsilon |= 0
		} else if bit < 0 {
			gamma |= 0
			epsilon |= 1
		}
		gamma = gamma << 1
		epsilon = epsilon << 1
	}

	powerConsumption = gamma >> 1 * epsilon >> 1
	return
}

func filterStack(lines []string, shouldReverse int) (filtered string) {
	var stack []string = lines
	i := 0
	for len(stack) > 1 {
		bits := linesToBits(stack)
		new_stack := []string{}
		take := 1 ^ shouldReverse
		if bits[i] < 0 {
			take = 0 ^ shouldReverse
		}

		for _, v := range stack {
			if convertRuneToInt(rune(v[i])) == take {
				new_stack = append(new_stack, v)
			}
		}

		stack = new_stack
		i++
	}
	return stack[0]
}

func calculateLifeSupport(lines []string) (lifeSupport int) {
	o2, _ := strconv.ParseInt(filterStack(lines, 0), 2, 16)
	co2, _ := strconv.ParseInt(filterStack(lines, 1), 2, 16)

	lifeSupport = int(o2 * co2)
	return
}

func main() {
	// test_case := helper.ReadFileAsStr("03_test.txt")
	// out := calculatePowerConsumption(test_case)
	// life := calculateLifeSupport(test_case)
	// fmt.Println(out, life)

	big_test := helper.ReadFileAsStr("03.txt")
	test_1 := calculatePowerConsumption(big_test)
	test_2 := calculateLifeSupport(big_test)
	fmt.Println(test_1, test_2)
}
