package main

import "testing"


func TestDay3(t *testing.T)  {
    result1, result2:= day_3("test.txt")
    if result1 != 4361{
        t.Errorf("Day 3 part 1 should be 4361, got %d", result1)
    }
    if result2 != 467835{
        t.Errorf("Day 3 part 2 should be 467835, got %d", result2)
    }

}
