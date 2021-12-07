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
	values := readInput(file)
	ans := calculate(readFish(values), 80)
	if ans != 5934 {
		t.Errorf("part1(file) = %d; want 5934", ans)
	}
}
