package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	sum := day_1("input.txt")

	fmt.Println(sum)

}
func day_1(input string) int {
	file, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		value := parse_line(line)
		sum += value

	}
	return sum
}

func parse_line(line string) int {
	var num string

	for i, v := range line {

		if unicode.IsDigit(v) {
			num = num + string(v)
		}
		num = num + replace_spelled(line[i:])

	}
	num = string(num[0]) + string(num[len(num)-1])

	result, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}
	return result

}

func replace_spelled(line string) string {
	var numbers = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for spelled, number := range numbers {
		if strings.HasPrefix(line, spelled) {
			return number
		}
	}

	return ""

}
