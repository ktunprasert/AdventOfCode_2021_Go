package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ktunprasert/AdventOfCode_2021_Go/helper"
)

func calculateNumberOfFishes(lines []string, day int) int {
	fishes := []int{}
	for _, f := range strings.Split(lines[0], ",") {
		f_int, _ := strconv.Atoi(f)
		fishes = append(fishes, f_int)
	}

	for i := 1; i <= day; i++ {
		for f_i, fishDay := range fishes {
			new_fishDay := fishDay - 1
			if new_fishDay < 0 {
				fishes[f_i] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[f_i] = new_fishDay
			}
		}
	}

	return len(fishes)
}

func optimalFish(lines []string, day int) (total int) {
	fishes := []int{}
	for _, f := range strings.Split(lines[0], ",") {
		f_int, _ := strconv.Atoi(f)
		fishes = append(fishes, f_int)
	}

	fish_cur := fishes

	i := 0
	for len(fish_cur) > 0 {
		// fmt.Println(fish_cur)
		new_fish := []int{}
		for _, v := range fish_cur {
			total++
			// fmt.Println(d, v, v-day)
			cur := v - day
			for cur < 0 {
				dcr := day + 1 + cur
				new_fish = append(new_fish, 8+dcr)
				cur += 7
			}
		}
		fish_cur = new_fish
		i++
	}

	return
}

func recurFish(d int) int {
	if d > 0 {
		return recurFish(d-9) + recurFish(d-7)
	}
	return 1
}

func recur(lines []string, day int) (total int) {
	// fishes := []int{}
	for _, f := range strings.Split(lines[0], ",") {
		f_int, _ := strconv.Atoi(f)
		total += recurFish(day - f_int)
		// fishes = append(fishes, f_int)
	}
	fmt.Println(total)
	return

	// for _,

}

func main() {
	// test_case := helper.ReadFileAsStr("06_test.txt")
	// recur(test_case, 256)
	// out := calculateNumberOfFishes(test_case, 80)
	// fmt.Println(out)
	// test_out := optimalFish(test_case, 80)
	// fmt.Println(test_out)
	big_test := helper.ReadFileAsStr("06.txt")
	recur(big_test, 256)
	// out := calculateNumberOfFishes(big_test, 256)
	// fmt.Println(out)
}
