package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Candidate struct {
	x int
	y int
}

type Gear struct {
	x       int
	y       int
	numbers []int
}

func main() {
	result1, result2 := day_3("input.txt")
	fmt.Println("Part 1: ", result1)
    fmt.Println("Part 2: ", result2)
}

func day_3(input string) (int, int) {
	file, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	matrix := parse_data(file)

	fmt.Printf("Imported data: %d x %d\n", len(matrix), len(matrix[0]))

	return process_data(matrix)

}

func parse_data(file []byte) [][]string {

	lines := strings.Split(string(file), "\n")
	ly := len(lines)
	lx := len(lines[0])

	matrix := make([][]string, ly)

	for i := range matrix {
		matrix[i] = make([]string, lx)
	}

	for i, line := range lines {
		matrix[i] = strings.Split(line, "")
	}

	return matrix

}

func process_data(matrix [][]string) (int, int) {
	number := ""
	sum := 0
	valid := false
	x, y := -1, -1
	gears := map[string][]int{}
	for i := range matrix {
		for j := range matrix[i] {
			_, err := strconv.Atoi(matrix[i][j])
			if err == nil {
				number += matrix[i][j]
				if !valid {

					x, y = get_gears(matrix, i, j)
					valid = isvalid(matrix, i, j)
				}
			} else {
				if number != "" {
					number_conv, _ := strconv.Atoi(number)
					if valid {
						sum += number_conv
						if x >= 0 && y >= 0 {
							key := fmt.Sprint(x)+ fmt.Sprint(y)

							nums, ok := gears[key]
							if ok {
								gears[key] = append(nums, number_conv)
							} else {
								gears[key] = []int{number_conv}
							}

						}

					}
					number = ""
					valid = false
				}
			}

		}

	}
    gear_ratios_sum := 0
    for _, gear:= range gears{
        if len(gear) == 2{
            gear_ratios_sum += gear[0]*gear[1]
        }

    }
	return sum, gear_ratios_sum

}
func isvalid(matrix [][]string, i int, j int) bool {
	symbols := "*#%$@!^&()+-//<>?=;{}"
	row := len(matrix)
	col := len(matrix[0])

	souroundings, _ := get_surrounding(matrix, i, j, row, col)
	for _, char := range souroundings {
		if strings.ContainsAny(symbols, char) {
			return true
		}
	}

	return false
}

func get_gears(matrix [][]string, i int, j int) (x int, y int) {
	gear := "*"
	row := len(matrix)
	col := len(matrix[0])

	souroundings, coords := get_surrounding(matrix, i, j, row, col)

	for i, char := range souroundings {
		if char == gear {
			return coords[i][0], coords[i][1]
		}

	}
	return -1, -1
}

func get_surrounding(matrix [][]string, i int, j int, x_limit int, y_limit int) ([]string, [][]int) {
	candidates := []Candidate{
		{x: i - 1, y: j - 1},
		{x: i - 1, y: j},
		{x: i - 1, y: j + 1},
		{x: i, y: j - 1},
		{x: i, y: j + 1},
		{x: i + 1, y: j - 1},
		{x: i + 1, y: j},
		{x: i + 1, y: j + 1},
	}

	valid := []string{}
	coords := [][]int{}
	for _, candidate := range candidates {
		if candidate.x >= 0 && candidate.y >= 0 && candidate.x < x_limit && candidate.y < y_limit {
			valid = append(valid, matrix[candidate.x][candidate.y])
			coords = append(coords, []int{candidate.x, candidate.y})
		}
	}

	return valid, coords
}

func print_map[T comparable](m map[T][]int) {
	for k, v := range m {
		fmt.Println(k, " --> ", v)

	}
}
