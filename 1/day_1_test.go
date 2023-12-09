package main

import "testing"

func TestDay1Part1(t *testing.T) {
	result := day_1("test.txt")
	if result != 142 {
		t.Errorf("Day 1 part 1 should be 142, got %d", result)
	}

}

func TestDay1Part2(t *testing.T) {
	result := day_1("test2.txt")
	if result != 281 {
		t.Errorf("Day 1 part 2 should be 281, got %d", result)
	}
}
