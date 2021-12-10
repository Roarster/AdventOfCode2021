package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func readInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	input := []string{}

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func part1(input []string) int {

	count := 0
	for _, line := range input {
		stack := []string{}
		last := ""
		for _, char := range line {
			switch string(char) {
			case "(":
				stack = append(stack, ")")
			case "[":
				stack = append(stack, "]")
			case "{":
				stack = append(stack, "}")
			case "<":
				stack = append(stack, ">")
			case ")", "]", "}", ">":

				last, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if last != string(char) {
					switch string(char) {
					case ")":
						count += 3
					case "]":
						count += 57
					case "}":
						count += 1197
					case ">":
						count += 25137
					}
					break
				}
			}
		}
	}
	return count
}

func part2(input []string) int {

	scores := []int{}
	for _, line := range input {
		stack := []string{}
		last := ""
		brokenLine := false
		for _, char := range line {
			switch string(char) {
			case "(":
				stack = append(stack, ")")
			case "[":
				stack = append(stack, "]")
			case "{":
				stack = append(stack, "}")
			case "<":
				stack = append(stack, ">")
			case ")", "]", "}", ">":

				last, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if last != string(char) {
					brokenLine = true
				}
			}
		}
		if !brokenLine {
			score := 0
			// this line is incomplete
			for i := len(stack) - 1; i >= 0; i-- {
				char := stack[i]
				score = score * 5
				switch char {
				case ")":
					score += 1
				case "]":
					score += 2
				case "}":
					score += 3
				case ">":
					score += 4
				}
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
