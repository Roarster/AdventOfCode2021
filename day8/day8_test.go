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

	ans := part1(readInput(file))
	if ans != 26 {
		t.Errorf("part1(file) = %d; want 26", ans)
	}
}

func TestSampleDataPart2(t *testing.T) {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ans := part2(readInput(file))
	if ans != 61229 {
		t.Errorf("part1(file) = %d; want 61229", ans)
	}
}
