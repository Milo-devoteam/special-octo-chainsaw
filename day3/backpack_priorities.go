package main

import (
	"bufio"
	"fmt"
	"os"
)

// var testText = "vJrwpWtwJgWrhcsFMMfFFhFp"

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)
	total_priorities := 0

	for scanner.Scan() {
		backpack_items := scanner.Text()
		compartment1 := backpack_items[:len(backpack_items)/2]
		compartment2 := backpack_items[len(backpack_items)/2:]

		counts := make(map[rune]int, len(backpack_items))

		for _, item := range compartment1 {
			if counts[item] > 0 {
				continue
			}
			counts[item]++
		}
		for _, item := range compartment2 {
			if counts[item] > 0 {
				priority := toPriority(int(item))
				total_priorities += priority
				break
			}
		}
	}
	fmt.Println("total_priorities:", total_priorities)

}

func toPriority(item int) int {
	if item < 97 {
		return item - 38
	} else {
		return item - 96
	}
}
