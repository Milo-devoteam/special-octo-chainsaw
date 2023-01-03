package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("input")

	packet := string(data)
	var start_letter int
	char_arr := make([]rune, 14)

	for i, char := range packet[:14] {
		char_arr[i] = char
	}

	fmt.Println(len(char_arr))
	for i, char := range packet[14:] {
		if isUnique(char_arr) {
			start_letter = i + 14
			break
		}
		dst := make([]rune, 13)
		copy(dst, char_arr[1:])

		char_arr = append(dst, char)

	}

	fmt.Println("starting letter:", start_letter)
}

func isUnique(char_arr []rune) bool {
	char_map := make(map[rune]bool)
	for _, char := range char_arr {
		_, ok := char_map[char]
		if ok {
			return false
		}
		char_map[char] = true
	}
	return true
}
