package main

import "testing"


func TestDay4(t *testing.T){
    part1, part2 := day_4("test.txt")

    if part1 != 13{
        t.Errorf("Day 4 part 1 should be 13, got %d", part1)
    }
    if part2 != 30{
        t.Errorf("Day 4 part 2 should be 30, got %d", part2)
    }
    

}
