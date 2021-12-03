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

	lines := readLines(file)
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	gammaString, epsilonString := "", ""
	for i := range lines[0] {
		character := getMostCommonCharacter(lines, i)
		gammaString += character
		if character == "0" {
			epsilonString += "1"
		} else {
			epsilonString += "0"
		}
	}

	gamma, _ := strconv.ParseInt(gammaString, 2, 32)
	epsilon, _ := strconv.ParseInt(epsilonString, 2, 32)
	fmt.Println("Part 1:", gamma*epsilon)
}

func part2(lines []string) {
	var oxygenCount, co2Count int64 = 0, 0
	filteredOxygenLines, filteredCO2Lines := lines, lines

	for i := range lines[0] {
		if oxygenCount == 0 {
			character := getMostCommonCharacter(filteredOxygenLines, i)
			var tempOxygenLines []string
			for _, line := range filteredOxygenLines {
				if string(line[i]) == character {
					tempOxygenLines = append(tempOxygenLines, line)
				}
			}
			if len(tempOxygenLines) == 1 {
				oxygenCount, _ = strconv.ParseInt(tempOxygenLines[0], 2, 64)
			} else {
				filteredOxygenLines = tempOxygenLines
			}
		}

		if co2Count == 0 {
			character := getMostCommonCharacter(filteredCO2Lines, i)
			tempCO2Lines := []string{}
			for _, line := range filteredCO2Lines {
				if string(line[i]) != character {
					tempCO2Lines = append(tempCO2Lines, line)
				}
			}
			if len(tempCO2Lines) == 1 {
				co2Count, _ = strconv.ParseInt(tempCO2Lines[0], 2, 64)
			} else {
				filteredCO2Lines = tempCO2Lines
			}
		}
	}
	fmt.Println("Part 2:", oxygenCount*co2Count)
}

func readLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	return values
}

func getMostCommonCharacter(lines []string, position int) string {
	zeroCount, oneCount := 0, 0
	for _, line := range lines {
		if string(line[position]) == "0" {
			zeroCount++
		} else {
			oneCount++
		}
	}
	if zeroCount > oneCount {
		return "0"
	}
	return "1"
}
