package main

import (
	"log"
	"os"
	"testing"
)

func TestSampleDataPart1(t *testing.T) {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	values, min, max := readInput(file)

	ans := part1(values, min, max)
	if ans != 37 {
		t.Errorf("part1(file) = %d; want 37", ans)
	}
}

func TestSampleDataPart2(t *testing.T) {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	values, min, max := readInput(file)
	ans := part2(values, min, max)
	if ans != 168 {
		t.Errorf("part2(file) = %d; want 168", ans)
	}
}
