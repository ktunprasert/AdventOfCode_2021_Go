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

func calculateLifeSupport(lines []string) (lifeSupport int) {
	var o2stack []string = lines
	var co2stack []string = lines
	i := 0
	for len(o2stack) > 1 {
		// O2
		o2_bits := linesToBits(o2stack)
		new_o2_stack := []string{}
		o2_take := 1
		if o2_bits[i] < 0 {
			o2_take = 0
		}

		for _, v := range o2stack {
			if convertRuneToInt(rune(v[i])) == o2_take {
				new_o2_stack = append(new_o2_stack, v)
			}
		}

		o2stack = new_o2_stack
		i++
	}

	i = 0
	for len(co2stack) > 1 {
		// CO2
		co2_bits := linesToBits(co2stack)
		new_co2_stack := []string{}
		co2_take := 0
		if co2_bits[i] < 0 {
			co2_take = 1
		}

		for _, v := range co2stack {
			if convertRuneToInt(rune(v[i])) == co2_take {
				new_co2_stack = append(new_co2_stack, v)
			}
		}

		co2stack = new_co2_stack
		i++
	}

	o2, _ := strconv.ParseInt(o2stack[0], 2, 16)
	co2, _ := strconv.ParseInt(co2stack[0], 2, 16)
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
