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

	values := readValues(file)

	lastValue := 0
	// start with negative 1 since we don't count the first
	increased := -1
	for i := range values {
		if i > 1 {
			total := values[i] + values[i-1] + values[i-2]
			if total > lastValue {
				increased++
			}
			lastValue = total
		}
	}

	fmt.Println("Number of increased:", increased)
}

func readValues(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	var values []int
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		values = append(values, i)
	}
	return values
}
