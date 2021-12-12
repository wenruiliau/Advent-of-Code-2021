package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Answer1(input string) int {
	content, err := ioutil.ReadFile(input)

	var boards [][][]int
	var board [][]int
	if err != nil {
		log.Fatal(err)
	}

	contentArray := strings.Split(string(content), "\n")
	numbers := contentArray[0]

	numbersInput := make([]int, 0)
	for _, num := range strings.Split(numbers, ",") {
		number, _ := strconv.Atoi(num)
		numbersInput = append(numbersInput, number)
	}

	for _, line := range contentArray[2:] {
		// new board
		if line == "" {
			boards = append(boards, board)
			board = make([][]int, 0)
			continue
		}
		rowNums := strings.Fields(line)
		numsToAppend := make([]int, 0)
		for _, num := range rowNums {
			i, _ := strconv.Atoi(num)

			numsToAppend = append(numsToAppend, i)
		}
		board = append(board, numsToAppend)
	}

	if board != nil {
		boards = append(boards, board)
	}

	for _, num := range numbersInput {
		markBoards(boards, num)
		solvedBoard := isSolved(boards)
		if solvedBoard != nil {
			answer := calculateAnswer(solvedBoard, num)
			return int(answer)
		}

	}
	return 0
}

func Answer2(input string) int {
	content, err := ioutil.ReadFile(input)

	var boards [][][]int
	var board [][]int
	if err != nil {
		log.Fatal(err)
	}

	contentArray := strings.Split(string(content), "\n")
	numbers := contentArray[0]

	numbersInput := make([]int, 0)
	for _, num := range strings.Split(numbers, ",") {
		number, _ := strconv.Atoi(num)
		numbersInput = append(numbersInput, number)
	}

	for _, line := range contentArray[2:] {
		// new board
		if line == "" {
			boards = append(boards, board)
			board = make([][]int, 0)
			continue
		}
		rowNums := strings.Fields(line)
		numsToAppend := make([]int, 0)
		for _, num := range rowNums {
			i, _ := strconv.Atoi(num)

			numsToAppend = append(numsToAppend, i)
		}
		board = append(board, numsToAppend)
	}
	boards = append(boards, board)

	remainingBoards := rangeArray(len(boards))

	var last int
	for _, num := range numbersInput {
		markBoards(boards, num)
		for index, board := range boards {
			lastSolvedBoard := isBoardSolved(board)
			if lastSolvedBoard {
				remainingBoards = remove(remainingBoards, index)
				if len(remainingBoards) == 1 {
					last = remainingBoards[0]
				} else if len(remainingBoards) == 0 {
					return calculateAnswer(boards[last], num)
				}
			}
		}
	}
	return 0
}

func rangeArray(length int) []int {
	var array []int
	for i := 0; i < length; i++ {
		array = append(array, i)
	}
	return array
}

func remove(slice []int, s int) []int {
	for i, other := range slice {
		if s == other {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func markBoards(boards [][][]int, num int) bool {
	for _, board := range boards {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == num {
					board[i][j] = -1
				}
			}
		}
	}
	return true
}

func isSolved(boards [][][]int) (board [][]int) {
	for _, board := range boards {
		if len(board) == 0 {
			continue
		}

		// check row
		for i := 0; i < len(board); i++ {
			rowWinner := true
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] != -1 {
					rowWinner = false
					break
				}
			}
			if rowWinner {
				return board
			}
		}

		// check column
		for k := 0; k < len(board[0]); k++ {
			colWinner := true
			for f := 0; f < len(board); f++ {
				if board[f][k] != -1 {
					colWinner = false
					break
				}
			}
			if colWinner {
				return board
			}
		}
	}

	return nil
}

func isBoardSolved(board [][]int) bool {
	if len(board) == 0 {
		return true
	}

	// check row
	for i := 0; i < len(board); i++ {
		rowWinner := true
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != -1 {
				rowWinner = false
				break
			}
		}
		if rowWinner {
			return true
		}
	}

	// check column
	for k := 0; k < len(board[0]); k++ {
		colWinner := true
		for f := 0; f < len(board); f++ {
			if board[f][k] != -1 {
				colWinner = false
				break
			}
		}
		if colWinner {
			return true
		}
	}
	return false
}

func calculateAnswer(board [][]int, lastCalled int) int {
	sum := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == -1 {
				continue
			}
			sum += board[i][j]
		}
	}
	return sum * lastCalled
}

func main() {
	fmt.Println("Answer for 1,", Answer1("input.txt"))
	fmt.Println("Answer for 2,", Answer2("input.txt"))
}
