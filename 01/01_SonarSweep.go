package main

import (
	"fmt"
	"math"

	"github.com/ktunprasert/AdventOfCode_2021_Go/helper"
)

func countDepthIncreased(depths []int64) (counter int64) {
	var prev int64 = math.MaxInt64
	for _, v := range depths {
		if v > prev {
			counter++
		}
		prev = v
	}
	return
}

func slidingWindowSum(depths []int64) (windowDepth []int64) {
	for i := 0; i < len(depths)-2; i++ {
		windowDepth = append(windowDepth, depths[i]+depths[i+1]+depths[i+2])
	}
	return
}

func main() {
	// Big Test 1
	big_test := helper.ReadFile("01.txt")
	// fmt.Println(big_test)
	out := countDepthIncreased(big_test)

	// Big Test 2
	window_test := slidingWindowSum(big_test)
	out2 := countDepthIncreased(window_test)
	fmt.Println(out, out2)

	// test_input := []int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	// out := countDepthIncreased(test_input)
	// out := slidingWindowSum(test_input)

	// fmt.Println(out)

}
