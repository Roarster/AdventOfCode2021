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

	scanner := bufio.NewScanner(file)
	lastValue := 0
	// start with negative 1 since we don't count the first
	increased := -1
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		if i > lastValue {
			increased++
		}
		lastValue = i
	}
	fmt.Println("Number of increased:", increased)
}
