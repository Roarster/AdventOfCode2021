package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type board struct {
	rows    [][]int
	columns [][]int
}

func (board *board) hasWon() bool {
	// check for empty row
	for _, thisRow := range board.rows {
		if len(thisRow) == 0 {
			return true
		}
	}

	// check for empty column
	for _, thisColumn := range board.columns {
		if len(thisColumn) == 0 {
			return true
		}
	}
	return false
}

func (board *board) calculateResult(number int) int {
	sum := 0
	for _, thisRow := range board.rows {
		for _, number := range thisRow {
			sum += number
		}
	}
	return sum * number
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbers, boards := readInput(file)

	part1(numbers, boards)
	part2(numbers, boards)
}

func readInput(file *os.File) (numbers []int, boards []board) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// read numbers
	parts := strings.Split(scanner.Text(), ",")
	for _, part := range parts {
		number, _ := strconv.Atoi(part)
		numbers = append(numbers, number)
	}

	// read boards
	var b board
	firstBoard := true
	rowNumber := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if firstBoard {
				firstBoard = false
			} else {
				boards = append(boards, b)
			}
			b = board{}
			rowNumber = 0
		} else {
			numbers := strings.Fields(text)
			row := []int{}
			var column []int

			for i, part := range numbers {
				number, _ := strconv.Atoi(part)
				row = append(row, number)
				if rowNumber == 0 {
					column = []int{}
					column = append(column, number)
					b.columns = append(b.columns, column)
				} else {
					column = b.columns[i]
					column = append(column, number)
					b.columns[i] = column
				}
			}
			b.rows = append(b.rows, row)
			rowNumber++
		}
	}
	boards = append(boards, b)
	return
}

func part1(numbers []int, boards []board) {

	for _, number := range numbers {
		boards = removeNumberFromBoards(number, boards)
		if board, found := findWinningBoard(boards); found {
			fmt.Println("Part 1 result:", board.calculateResult(number))
			return
		}
	}
}

func part2(numbers []int, boards []board) {

	for _, number := range numbers {
		boards = removeNumberFromBoards(number, boards)
		losingBoards := []board{}
		for _, board := range boards {

			if !board.hasWon() {
				losingBoards = append(losingBoards, board)
			} else if len(boards) == 1 {
				// if this was the last board and has just won
				fmt.Println("Part 2 result:", board.calculateResult(number))
				return
			}
		}
		boards = losingBoards
	}
}

func removeNumberFromBoards(number int, boards []board) (newBoards []board) {

	for _, thisBoard := range boards {
		newBoard := board{}
		// ToDo: abstract this to a single func
		newRows := [][]int{}
		for _, row := range thisBoard.rows {
			filteredRow := []int{}
			for _, entry := range row {
				if entry != number {
					filteredRow = append(filteredRow, entry)
				}
			}
			newRows = append(newRows, filteredRow)
		}
		newBoard.rows = newRows

		newColumns := [][]int{}
		for _, column := range thisBoard.columns {
			filteredColumn := []int{}
			for _, entry := range column {
				if entry != number {
					filteredColumn = append(filteredColumn, entry)
				}
			}
			newColumns = append(newColumns, filteredColumn)
		}
		newBoard.columns = newColumns
		newBoards = append(newBoards, newBoard)
	}
	return
}

func findWinningBoard(boards []board) (board board, found bool) {
	for _, thisBoard := range boards {
		if thisBoard.hasWon() {
			return thisBoard, true
		}
	}

	return
}
