package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id    int
	red   int
	blue  int
	green int
	valid bool
}

const RED = 12
const GREEN = 13
const BLUE = 14

func main() {

	sum_part1, sum_part2 := day_2("input.txt")

	fmt.Println("Part 1: ", sum_part1)
	fmt.Println("Part 2: ", sum_part2)

}
func day_2(input string) (int, int) {
	file, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum_part1 := 0
	sum_part2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := parse_game(line)
		if game.valid {
			sum_part1 += game.id
		}
		sum_part2 += game.red * game.blue * game.green
	}

	return sum_part1, sum_part2
}

func parse_game(line string) Game {
	game := strings.Split(line, ": ")         // Game, Samples
	game_id := strings.Split(game[0], " ")[1] // Game 0 -> 0
	id, err := strconv.Atoi(game_id)
	if err != nil {
		log.Fatal(err)
	}

	colors := strings.Split(game[1], "; ")

	red, blue, green := 0, 0, 0

	valid := true

	for _, sample := range colors {
		values := strings.Split(sample, ", ")
		for _, color := range values {
			color := strings.Split(color, " ")
			value, err := strconv.Atoi(color[0])
			if err != nil {
				log.Fatal(err)
				value = 0
			}

			switch color[1] {
			case "red":
				if value > red {
					red = value
				}
			case "blue":
				if value > blue {
					blue = value
				}

			case "green":
				if value > green {
					green = value

				}

			}

		}
	}

	result := Game{id: id, red: red, blue: blue, green: green, valid: valid}

	if result.red > RED || result.green > GREEN || result.blue > BLUE {
		result.valid = false
	}
	return result
}
