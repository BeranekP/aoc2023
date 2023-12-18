package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Counts map[int]int

func main() {

	part1, part2 := day_4("input.txt")

	fmt.Println("Part 1: ", part1)
    fmt.Println("Part 2: ", part2)

}

func day_4(file string) (int, int) {
	bytes, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}
	data := strings.TrimSpace(string(bytes))
	return process_data(data)
}

func process_data(data string) (int, int) {
	cards := strings.Split(data, "\n")
	score := 0
	matching := 0
    counts := Counts{}
	for card_id, card := range cards {
		id_values := strings.Split(card, ": ")
		values := strings.Split(id_values[1], " | ")
		winning := strings.Split(strings.ReplaceAll(values[0], "  ", " "), " ")
		winning_conv := to_int(winning)
		my_numbers := strings.Split(strings.ReplaceAll(values[1], "  ", " "), " ")
		my_numbers_conv := to_int(my_numbers)
        counts[card_id] += 1
		for _, number := range my_numbers_conv {
			if slices.Contains(winning_conv, number) {
				matching += 1
                counts[card_id+matching] += counts[card_id]
			}

		}
		if matching > 0 {
			score += int(math.Pow(float64(2), float64(matching-1)))
		}
		matching = 0
	}
	return score, sum_counts(counts)

}

func to_int(input []string) []int {
	result := []int{}
	for _, value := range input {
		if value != "" {
			number, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}

			result = append(result, number)
		}
	}

	return result

}

func sum_counts(counts Counts) int{
    sum := 0
    for _, value:= range counts{
        sum += value
    }
    return sum
}
