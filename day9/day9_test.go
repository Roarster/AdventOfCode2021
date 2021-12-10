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
	if ans != 15 {
		t.Errorf("part1(file) = %d; want 15", ans)
	}
}
