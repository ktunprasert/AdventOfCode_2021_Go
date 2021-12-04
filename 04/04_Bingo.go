package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ktunprasert/AdventOfCode_2021_Go/helper"
)

func makeBingos(lines []string) [][][]int {
	boards := [][][]int{}
	single_board := [][]int{}
	for _, b := range lines {
		if b != "" {
			temp := []int{}
			for _, v := range strings.Split(b, " ") {
				if v != "" {
					v_i, _ := strconv.Atoi(v)
					temp = append(temp, v_i)
				}
			}
			single_board = append(single_board, temp)

		} else {
			boards = append(boards, single_board)
			single_board = [][]int{}
		}
	}
	boards = append(boards, single_board)
	return boards
}

func sum(b []int) (total int) {
	for _, val := range b {
		total += val
	}
	return
}

func boardSum(board [][]int) (total int) {
	for _, row := range board {
		for _, v := range row {
			if v > 0 {
				total += v
			}
		}
	}
	return
}

func checkWin(board [][]int) bool {
	for i, x := range board {
		col, row := []int{}, []int{}
		for j, _ := range x {
			col = append(col, board[i][j])
			row = append(row, board[j][i])
		}
		if sum(col) == -5 || sum(row) == -5 {
			return true
		}
	}
	return false
}

func turnBoard(board [][]int, num int) [][]int {
	for i, row := range board {
		for j, value := range row {
			if value == num {
				board[i][j] = -1
			}
		}
	}
	return board
}

func solveBingo(lines []string) (firstBoard int, lastBoard int) {
	numbers := strings.Split(lines[0], ",")
	boards := makeBingos(lines[2:])
	total_len := len(boards)

	winning_boards := [][][]int{}
	winning_nums := []int{}

	for _, n := range numbers {
		n_int, _ := strconv.Atoi(n)

		for i, board := range boards {
			if i > len(boards) {
				break
			}
			board = turnBoard(board, n_int)
			if checkWin(board) {
				winning_boards = append(winning_boards, board)
				winning_nums = append(winning_nums, n_int)
				fmt.Println(len(boards), i, i+1, len(boards[i+1:]))
				if len(boards[i+1:]) > 0 {
					boards = append(boards[:i], boards[i+1:]...)
				} else {
					fmt.Println("here")
					boards = boards[:i]
				}
				// boards = append(boards[:i], boards[i+1:]...)
			}
		}
	}

	firstBoard = boardSum(winning_boards[0]) * winning_nums[0]
	lastBoard = boardSum(winning_boards[total_len-1]) * winning_nums[total_len-1]
	return
	// fmt.Println(winning_boards, winning_nums)
	// fmt.Println(boardSum(winning_boards[2]))
	// return 0, 0
	// fmt.Println(len(winning_boards), winning_boards, winning_nums)
	// return boardSum(win_board) * win_num
}

func main() {
	// test_case := helper.ReadFileAsStr("04_test.txt")
	// out1, out2 := solveBingo(test_case)
	// fmt.Println(out1, out2)

	big_test := helper.ReadFileAsStr("04.txt")
	ans1, ans2 := solveBingo(big_test)
	fmt.Println(ans1, ans2)
}
