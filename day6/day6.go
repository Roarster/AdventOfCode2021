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
	values := readInput(file)

	fmt.Println("Part 1", calculate(readFish(values), 80))
	fmt.Println("Part 2", calculate(readFish(values), 256))
}

func calculate(fishes map[int]int, numberOfDays int) int {

	populateFishCounts(fishes, numberOfDays)
	return calculateFishTotal(fishes)
}

func calculateFishTotal(fishes map[int]int) int {
	count := 0
	for _, v := range fishes {
		count += v
	}
	return count
}

func populateFishCounts(fishes map[int]int, numberOfDays int) {
	for day := 0; day < numberOfDays; day++ {
		amountToDuplicate := fishes[0]
		fishes[0] = fishes[1]
		fishes[1] = fishes[2]
		fishes[2] = fishes[3]
		fishes[3] = fishes[4]
		fishes[4] = fishes[5]
		fishes[5] = fishes[6]
		fishes[6] = fishes[7] + amountToDuplicate
		fishes[7] = fishes[8]
		fishes[8] = amountToDuplicate
	}
}

func readInput(file *os.File) []int {
	values := []int{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	parts := strings.Split(scanner.Text(), ",")
	for _, part := range parts {
		number, _ := strconv.Atoi(part)
		values = append(values, number)
	}

	return values
}

func readFish(values []int) map[int]int {
	fishes := make(map[int]int)
	for _, value := range values {
		fishes[value]++
	}
	return fishes
}
