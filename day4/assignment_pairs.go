package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	contains_count := 0
	overlaps_count := 0
	for scanner.Scan() {
		// Parse input pairs
		inputPair := strings.Split(scanner.Text(), ",")
		inputA := toIntArr(inputPair[0])
		inputB := toIntArr(inputPair[1])

		fmt.Println("inputA", inputA)
		fmt.Println("inputB", inputB)

		if contains(inputA, inputB) || contains(inputB, inputA) {
			fmt.Println("CONTAINS")
			contains_count++
		}
		if overlaps(inputA, inputB) || overlaps(inputB, inputA) {
			fmt.Println("OVERLAPS")
			overlaps_count++
		}
	}
	fmt.Println("contains total:", contains_count)
	fmt.Println("overlaps total:", overlaps_count)
}

// Check whether start of 1st input > start of 2nd input && end of 1st input > end of 2nd input
// Or vice-versa.
func contains(a, b []int) bool {
	return (a[0] <= b[0] && a[1] >= b[1])
}

func overlaps(a, b []int) bool {
	return (a[0] <= b[0] && a[1] >= b[0])
}

func toIntArr(rangeStr string) []int {
	rangeInput := strings.Split(rangeStr, "-")
	intArr := make([]int, len(rangeInput))

	for i, str := range rangeInput {
		// Suppress errors like any good engineer
		intArr[i], _ = strconv.Atoi(str)
	}
	return intArr
}
