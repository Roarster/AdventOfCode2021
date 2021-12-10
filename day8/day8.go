package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Part 1", part1(readInput(file)))
}

func readInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)

	// read lines
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func part1(lines []string) int {

	count := 0
	for _, line := range lines {
		output := strings.Split(line, "|")[1]
		tokens := strings.Fields(output)
		for _, token := range tokens {
			if len(token) == 2 || len(token) == 3 || len(token) == 4 || len(token) == 7 {
				count++
			}
		}
	}
	return count
}

func part2(lines []string) int {

	total := 0
	//for _, line := range lines {
	//	parts := strings.Split(line,"|")
	//	input := parts[0]
	//	output := parts[1]
	//	tokens := append(strings.Fields(input), strings.Fields(output)...)
	//	unsolvedTokens := tokens
	//	for {
	//		tempTokens := []string{}
	//		for _, token := range unsolvedTokens {
	//			if len(token) == 2 {
	//				// 1
	//			} else if len(token) == 3 {
	//				// 7
	//			} else if len(token) == 4 {
	//				// 4
	//			} else if len(token) == 7 {
	//				// 8
	//			} else {
	//				tempTokens = append(tempTokens, token)
	//			}
	//		}
	//		unsolvedTokens = tempTokens
	//		if len(unsolvedTokens) == 0 {
	//			break
	//		}
	//	}
	//}
	return total
}
