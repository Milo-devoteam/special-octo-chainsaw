package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Read the input file
	f, err := os.Open("input")
	check(err)

	totals := make([]int64, 0, 10)

	// print(slice)
	scanner := bufio.NewScanner(f)
	var sum int64 = 0
	// sum all  elf's calories
	for scanner.Scan() {
		text := scanner.Text()

		// split by new line to get each elf's cals.
		if text == "" {
			totals = append(totals, sum)
			// fmt.Printf("new total: %d\n", sum)
			sum = 0
			continue
		}
		num, err := strconv.Atoi(text)
		check(err)
		sum += int64(num)
	}
	var greatest_calories int64 = 0
	// Loop to compare elf totals; keep the greatest number
	for _, total := range totals {
		if total > greatest_calories {
			greatest_calories = total
		}
	}
	fmt.Printf("greatest_calories: %d\n", greatest_calories)

	// Return greatest number

}
