package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type visibility struct {
	top    int
	bottom int
	left   int
	right  int
}

var test_input string = `30373
25512
65332
33549
35390`

func main() {
	input, _ := os.ReadFile("input")
	tree_rows := strings.Split(string(input), "\n")
	heights := toInts(tree_rows)

	visible_tree_count := totalTreesVisible(heights)
	fmt.Println("total trees visible:", visible_tree_count)

	score := highestScenicScore(heights)
	fmt.Println("top score:", score)
}

func highestScenicScore(tree_heights [][]int) int {
	// Find # trees visible in a direction (starting from the base tree)
	top_scenic_score := 0
	for i, tree_row := range tree_heights {
		if i == 0 || i == len(tree_heights)-1 {
			continue
		}
		for j, tree_height := range tree_row {
			if j == 0 || j == len(tree_row)-1 {
				continue
			}

			vis := visibility{top: 0, left: 0, right: 0, bottom: 0}
			// Calculate visiblility scores in a direction
			// 2. general function to calculate the number of trees visible

			// calculate top
			for row := i - 1; row >= 0; row-- {
				vis.top++
				if tree_height <= tree_heights[row][j] {
					break
				}
			}

			// bottom
			for row := i + 1; row < len(tree_heights); row++ {
				vis.bottom++
				if tree_height <= tree_heights[row][j] {
					break
				}
			}
			// left
			for col := j - 1; col >= 0; col-- {
				vis.left++
				if tree_height <= tree_row[col] {
					break
				}
			}

			// right
			for col := j + 1; col < len(tree_row); col++ {
				vis.right++
				if tree_height <= tree_row[col] {
					break
				}
			}
			scenic_score := vis.bottom * vis.top * vis.left * vis.right

			// fmt.Println("scenic score for", tree_height)
			// fmt.Println("top", vis.top)
			// fmt.Println("bottom", vis.bottom)
			// fmt.Println("left", vis.left)
			// fmt.Println("right", vis.right)
			if scenic_score > top_scenic_score {
				top_scenic_score = scenic_score
			}
		}
	}
	return top_scenic_score
}

func totalTreesVisible(tree_heights [][]int) int {
	var visible_tree_count int
	for i, tree_row := range tree_heights {
		// Make rules for row
		if i == 0 || i == len(tree_heights)-1 {
			visible_tree_count += len(tree_row)
			continue
		}
		for j, tree_height := range tree_row {
			// If on outside of row/column
			// if 1st/last col or 1st/last row:
			if j == 0 || j == len(tree_row)-1 {
				visible_tree_count += 1
				continue
			}

			// Check the surrounding trees
			if isVisibleTop(tree_height, i, j, tree_heights) || isVisibleBottom(tree_height, i, j, tree_heights) || isVisibleLeft(tree_height, j, tree_row) || isVisibleRight(tree_height, j, tree_row) {
				visible_tree_count += 1
				continue
			}
		}
	}
	return visible_tree_count
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}

func isVisibleTop(tree_height, i, j int, tree_heights [][]int) bool {
	for col := 0; col < i; col++ {
		if tree_height <= tree_heights[col][j] {
			return false
		}
	}
	return true
}
func isVisibleBottom(tree_height, i, j int, tree_heights [][]int) bool {
	for row := i + 1; row < len(tree_heights); row++ {
		if tree_height <= tree_heights[row][j] {
			return false
		}
	}
	return true
}
func isVisibleLeft(tree_height, j int, tree_heights []int) bool {
	for col := 0; col < j; col++ {
		if tree_height <= tree_heights[col] {
			return false
		}
	}
	return true
}
func isVisibleRight(tree_height, j int, tree_heights []int) bool {
	for col := j + 1; col < len(tree_heights); col++ {
		if tree_height <= tree_heights[col] {
			return false
		}
	}
	return true
}

func toInts(str_arr []string) [][]int {
	tree_heights := make([][]int, len(str_arr))

	for i, line := range str_arr {
		tree_heights[i] = make([]int, len(line))
		for j, char := range line {
			var err error
			tree_heights[i][j], err = strconv.Atoi(string(char))
			must(err)
		}
	}
	return tree_heights
}
