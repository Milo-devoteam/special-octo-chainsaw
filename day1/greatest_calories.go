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
	// This is horrible, but I can't think
	var greatest_calories int64 = 0
	var second_greatest int64 = 0
	var third_greatest int64 = 0

	// Loop to compare elf totals; keep the greatest number
	for _, total := range totals {
		if total > greatest_calories {
			greatest_calories = total
		} else if total > second_greatest {
			second_greatest = total
		} else if total > third_greatest {
			third_greatest = total
		}
	}
	fmt.Printf("greatest_calories: %d\n", greatest_calories)
	fmt.Printf("second_greatest: %d\n", second_greatest)
	fmt.Printf("third_greatest: %d\n", third_greatest)

	fmt.Printf("sum of top 3 elves: %d\n", greatest_calories+second_greatest+third_greatest)
	// Return greatest number

}
