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

	scanner := bufio.NewScanner(file)
	forward, depth, aim := 0, 0, 0
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		delta, _ := strconv.Atoi(parts[1])
		if parts[0] == "forward" {
			forward += delta
			depth += aim * delta
		}
		if parts[0] == "up" {
			aim -= delta
		}
		if parts[0] == "down" {
			aim += delta
		}
	}
	fmt.Println("Result:", forward*depth)
}
