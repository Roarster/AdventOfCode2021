package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	xPos int
	yPos int
}

type line struct {
	start point
	end   point
}

type oceanFloor struct {
	ventCounts [1000][1000]int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := readInput(file)
	part1(lines)
	part2(lines)
}

func readInput(file *os.File) []line {
	scanner := bufio.NewScanner(file)

	// read numbers
	lines := []line{}
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Fields(text)
		lines = append(lines, line{convertStringToPoint(parts[0]), convertStringToPoint(parts[2])})
	}
	return lines
}

func convertStringToPoint(text string) point {
	parts := strings.Split(text, ",")
	xPos, _ := strconv.Atoi(parts[0])
	yPos, _ := strconv.Atoi(parts[1])
	return point{xPos, yPos}
}

func part1(lines []line) {
	oceanFloor := oceanFloor{}
	for _, line := range lines {
		if line.isHorOrVert() {
			oceanFloor.markVents(line)
		}
	}

	fmt.Println("Part 1:", oceanFloor.getNumberOfOverlappingVents())
}

func part2(lines []line) {
	oceanFloor := oceanFloor{}
	for _, line := range lines {
		oceanFloor.markVents(line)
	}
	fmt.Println("Part 2:", oceanFloor.getNumberOfOverlappingVents())
}

func (line *line) isHorOrVert() bool {
	return line.start.xPos == line.end.xPos || line.start.yPos == line.end.yPos
}

func (oceanFloor *oceanFloor) markVents(line line) {
	if line.start.xPos == line.end.xPos {
		// vertical line
		start := getMin(line.start.yPos, line.end.yPos)
		end := getMax(line.start.yPos, line.end.yPos)
		for y := start; y <= end; y++ {
			oceanFloor.ventCounts[line.start.xPos][y]++
		}
	} else if line.start.yPos == line.end.yPos {
		// horizontal line
		start := getMin(line.start.xPos, line.end.xPos)
		end := getMax(line.start.xPos, line.end.xPos)
		for x := start; x <= end; x++ {
			oceanFloor.ventCounts[x][line.end.yPos]++
		}
	} else {
		// diagonal line
		count := 0
		if line.start.xPos < line.end.xPos { // left to right, x increasing
			for x := line.start.xPos; x <= line.end.xPos; x++ {
				if line.start.yPos > line.end.yPos {
					oceanFloor.ventCounts[x][line.start.yPos-count]++
				} else {
					oceanFloor.ventCounts[x][line.start.yPos+count]++
				}
				count++
			}
		} else {
			for x := line.start.xPos; x >= line.end.xPos; x-- {
				if line.start.yPos > line.end.yPos {
					oceanFloor.ventCounts[x][line.start.yPos-count]++
				} else {
					oceanFloor.ventCounts[x][line.start.yPos+count]++
				}
				count++
			}
		}

	}
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (oceanFloor *oceanFloor) getNumberOfOverlappingVents() int {

	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if oceanFloor.ventCounts[x][y] > 1 {
				count++
			}
		}
	}
	return count
}
