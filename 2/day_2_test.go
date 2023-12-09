package main

import "testing"

func TestDay2(t *testing.T){
    part1, part2 := day_2("test.txt")
    if part1 != 8 {
        t.Errorf("Day 2 part 1 should be 8; got %d", part1)
    }

    if part2 != 2286 {
        t.Errorf("Day 2 part2 should be 2286; got %d", part2)
    }
}
