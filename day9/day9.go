package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Part 1", part1(readInput(file)))
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
