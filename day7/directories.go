package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type file struct {
	name string
	size string
}

type paths map[string]int

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	var prevNotCommand = false

	var path_sizes = make(paths, 0)
	var total_path_sizes = make(paths, 1000)
	var dir_size int
	var total_size int
	var current_dir string = "/"

	var char = "$"
	fmt.Println(char[0])
	for scanner.Scan() {
		input := scanner.Text()
		args := strings.Split(input, " ")
		if isCommand(input) {
			if !prevNotCommand {
				path_sizes[current_dir] = dir_size
				dir_size = 0
				prevNotCommand = true
			}
			if args[1] == "cd" {
				fmt.Println(input)
				current_dir = path.Join(current_dir, args[2])
			}
			continue
		}
		prevNotCommand = false
		if args[0] == "dir" {
			// fmt.Println(input)
			continue // ignore
		}
		size, _ := strconv.Atoi(args[0])
		// fmt.Println(size)
		dir_size += size
	}
	// fmt.Println(path_sizes)

	for path, size := range path_sizes {
		total_path_sizes[path] += size
		// Iteratively check if each path is subpath of another path
		for sub, size2 := range path_sizes {
			if path == sub {
				fmt.Println(sub, "not subpath of", path)
				continue
			}
			rel, err := filepath.Rel(path, sub)
			if err != nil {
				fmt.Println(sub, "not subpath of", path)
				continue
			}
			if strings.HasPrefix(rel, "..") || rel == ".." {
				fmt.Println(sub, "not subpath of", path)
				continue
			}
			// fmt.Println(sub, "is subpath of", path)
			total_path_sizes[path] += size2
		}
		if total_path_sizes[path] <= 100000 {
			total_size += total_path_sizes[path]
		}
	}
	fmt.Println("Root path size:", total_path_sizes["/"])
	free_space := 70000000 - total_path_sizes["/"]
	space_needed := 30000000 - free_space
	fmt.Println("free_space:", free_space)
	fmt.Println("space_needed:", space_needed)

	min_size_path := "/"
	for path, size := range total_path_sizes {
		if size >= space_needed {
			fmt.Println("potential directory:", path, "size:", size)
			if size < total_path_sizes[min_size_path] {
				fmt.Println("New Min path:", path)
				min_size_path = path
			}
		}
	}
	// fmt.Println(total_path_sizes["/"])
	// fmt.Println("total paths size:", total_size)
	fmt.Println("Min path:", min_size_path)
	fmt.Println("Min size:", total_path_sizes[min_size_path])

}

func isCommand(input string) bool {
	return input[0] == '$'
}

// func updateDir(current_dir, newDir string) string {
// 	if path.IsAbs(newDir) {
// 		return newDir
// 	}
// 	return path.Join(current_dir, newDir)
// }
