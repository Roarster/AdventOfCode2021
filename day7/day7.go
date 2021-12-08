package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	values, min, max := readInput(file)
	fmt.Println("Part 1", part1(values, min, max))
	fmt.Println("Part 1", part2(values, min, max))
}

func part1(values map[int]int, min int, max int) int {

	lowestFuel := -1
	for i := min; i <= max; i++ {
		fuel := calculatePart1Fuel(values, i)

		if lowestFuel == -1 || fuel < lowestFuel {
			lowestFuel = fuel
		}
	}
	return lowestFuel
}

func part2(values map[int]int, min int, max int) int {

	lowestFuel := -1
	for i := min; i <= max; i++ {
		fuel := calculatePart2Fuel(values, i)

		if lowestFuel == -1 || fuel < lowestFuel {
			lowestFuel = fuel
		}
	}
	return lowestFuel
}

func readInput(file *os.File) (values map[int]int, min int, max int) {
	values = make(map[int]int)
	min = 1000
	max = 0
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	parts := strings.Split(scanner.Text(), ",")
	for _, part := range parts {
		number, _ := strconv.Atoi(part)
		values[number]++
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
	}

	return
}

func calculatePart1Fuel(values map[int]int, position int) (fuel int) {
	for k, v := range values {
		fuel += abs((position - k) * v)
	}
	return
}

func calculatePart2Fuel(values map[int]int, position int) (fuel int) {
	for k, v := range values {
		fuelForPosition := sum(abs(position-k)) * v
		fuel += fuelForPosition
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sum(n int) (result int) {
	if n > 0 {
		result = n + sum(n-1)
		return result
	}
	return 0
}
