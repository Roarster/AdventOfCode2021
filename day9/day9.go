package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := readInput(file)
	fmt.Println("Part 1", part1(input))
	fmt.Println("Part 2", part2(input))
}

func readInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	input := [][]int{}

	for scanner.Scan() {
		line := []int{}
		text := scanner.Text()
		for _, c := range text {
			number, _ := strconv.Atoi(string(c))
			line = append(line, number)
		}
		input = append(input, line)
	}
	return input
}

func part1(lines [][]int) int {

	count := 0
	for x, line := range lines {
		for y, value := range line {

			if (x == 0 || value < lines[x-1][y]) &&
				(y == 0 || value < lines[x][y-1]) &&
				(x == len(lines)-1 || value < lines[x+1][y]) &&
				(y == len(lines[0])-1 || value < lines[x][y+1]) {
				count += 1 + value
			}
		}
	}
	return count
}

func part2(lines [][]int) int {
	totals := []int{}
	for x, line := range lines {
		for y, _ := range line {
			total := traverse(lines, x, y)
			if total > 0 {
				totals = append(totals, total)
			}
		}
	}
	fmt.Println()
	sort.Ints(totals)
	return totals[len(totals)-1] * totals[len(totals)-2] * totals[len(totals)-3]
}

func traverse(lines [][]int, x int, y int) int {
	if x >= len(lines) {
		return 0
	}
	if y >= len(lines[0]) {
		return 0
	}
	if x < 0 {
		return 0
	}
	if y < 0 {
		return 0
	}
	if lines[x][y] == 9 {
		return 0
	}
	lines[x][y] = 9
	return 1 + traverse(lines, x+1, y) + traverse(lines, x, y+1) + traverse(lines, x-1, y) + traverse(lines, x, y-1)
}
