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
	forward := 0
	depth := 0
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		delta, _ := strconv.Atoi(parts[1])
		if parts[0] == "forward" {
			forward += delta
		}
		if parts[0] == "up" {
			depth -= delta
		}
		if parts[0] == "down" {
			depth += delta
		}
	}
	fmt.Println("Result:", forward*depth)
}
