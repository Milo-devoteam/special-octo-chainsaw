package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// var testText = "vJrwpWtwJgWrhcsFMMfFFhFp"

func main() {
	// This is truly horrible, but I just want to finish the challenge
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	data, _ := os.ReadFile("input")
	backpacks := strings.Split(string(data), "\n")

	total_priorities_1 := calculate_points_p1(*scanner)
	total_priorities_2 := calculate_points_p2(backpacks)

	fmt.Println("total_priorities part 1:", total_priorities_1)
	fmt.Println("total_priorities part 2:", total_priorities_2)
}

func toPriority(item rune) int {
	if item < 97 {
		return int(item) - 38
	} else {
		return int(item) - 96
	}
}

func calculate_points_p2(backpacks []string) int {
	// Group into sets of 3
	total_priorities := 0

	for i := 0; i < len(backpacks)-1; i += 3 {
		group_counts := make(map[rune]int, 57)

		// Separate backpacks into groups
		for j := i; j < i+3; j++ {

			backpack_items := backpacks[j]
			backpack_counts := make(map[rune]int, 57)

			for _, item := range backpack_items {
				if backpack_counts[item] > 0 {
					continue
				}

				// increment counters
				backpack_counts[item]++
				group_counts[item]++

				if group_counts[item] > 2 {
					total_priorities += toPriority(item)
					break
				}
			}
		}
	}
	return total_priorities
}

func calculate_points_p1(scanner bufio.Scanner) int {
	total_priorities := 0
	for scanner.Scan() {
		backpack_items := scanner.Text()

		counts := make(map[rune]int, len(backpack_items)/2)

		// Improvement: perform both operations in the same loop (considering strings are equal length)
		for _, item := range backpack_items[:len(backpack_items)/2] {
			if counts[item] > 0 {
				continue
			}
			counts[item]++
		}
		for _, item := range backpack_items[len(backpack_items)/2:] {
			if counts[item] > 0 {
				total_priorities += toPriority(item)
				break
			}
		}
	}
	return total_priorities
}
