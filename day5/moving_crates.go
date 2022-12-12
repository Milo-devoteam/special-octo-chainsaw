package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input")

	file := string(data)

	thing := strings.Split(file, "\n\n")
	r1 := strings.NewReader(thing[0])
	r2 := strings.NewReader(thing[1])
	r3 := strings.NewReader(thing[1])

	crates_scanner := bufio.NewScanner(r1)

	crate_columns := parseCrateColumns(*crates_scanner)

	fmt.Println(crate_columns)

	moves_scanner := bufio.NewScanner(r2)
	moves_scanner2 := bufio.NewScanner(r3)

	dst := make([][]string, len(crate_columns))
	copy(dst, crate_columns)

	calculateCrateMover9001(crate_columns, *moves_scanner2)
	calculateCrateMover9000(dst, *moves_scanner)

}

func parseCrateColumns(crates_scanner bufio.Scanner) [][]string {
	// Construct 2d Array
	// 1st dimension == columns, 2nd dimension = row
	crate_columns := make([][]string, 9)

	// Parse 1st part of input
	for crates_scanner.Scan() {
		line := crates_scanner.Text()

		// Check if 1st char is '[', otherwise it's the id line
		if line[0] == byte('[') {
			column_number := 0
			for i := 1; i < len(line); i += 4 {
				if line[i] == 32 {
					column_number++
					continue
				}
				crate_columns[column_number] = append(crate_columns[column_number], string(line[i]))
				column_number++
			}
		}
	}
	return crate_columns
}

func calculateCrateMover9000(crate_columns [][]string, moves bufio.Scanner) {
	for moves.Scan() {
		command := moves.Text()
		cmd_args := strings.Split(command, " ")

		crate_quantity, _ := strconv.Atoi(cmd_args[1])
		source_column, _ := strconv.Atoi(cmd_args[3])
		destination_column, _ := strconv.Atoi(cmd_args[5])

		// Adjust the columns to be 0 indexed
		source_column -= 1
		destination_column -= 1

		// fmt.Println("crate_quantity", crate_quantity)
		// fmt.Println("source_column", source_column)
		// fmt.Println("destination_column", destination_column)

		// fmt.Println("**before**")
		// fmt.Println("source:", crate_columns[source_column])
		// fmt.Println("destination:", crate_columns[destination_column])

		for _, crate := range crate_columns[source_column][:crate_quantity] {
			crate_columns[destination_column] = append([]string{crate}, crate_columns[destination_column]...)
		}
		crate_columns[source_column] = crate_columns[source_column][crate_quantity:]

		// fmt.Println("**after**")
		// fmt.Println("source:", crate_columns[source_column])
		// fmt.Println("destination:", crate_columns[destination_column])
		// fmt.Println("")
		// fmt.Println("")

	}

	str := ""
	for i := range crate_columns {
		str += string(crate_columns[i][0])
	}
	fmt.Println(str)

}

func calculateCrateMover9001(crate_columns [][]string, moves bufio.Scanner) {
	for moves.Scan() {
		move := moves.Text()
		move_args := strings.Split(move, " ")

		crate_quantity, _ := strconv.Atoi(move_args[1])
		source_column, _ := strconv.Atoi(move_args[3])
		destination_column, _ := strconv.Atoi(move_args[5])

		// Adjust the columns to be 0 indexed
		source_column -= 1
		destination_column -= 1

		fmt.Println("crate_quantity", crate_quantity)
		fmt.Println("source_column", source_column)
		fmt.Println("destination_column", destination_column)

		fmt.Println("**before**")
		fmt.Println("source:", len(crate_columns[source_column]), "-", crate_columns[source_column])
		fmt.Println("destination:", len(crate_columns[destination_column]), "-", crate_columns[destination_column])

		dst := make([]string, len(crate_columns[source_column][:crate_quantity]))

		copy(dst, crate_columns[source_column][:crate_quantity])

		crate_columns[destination_column] = append(dst, crate_columns[destination_column]...)

		fmt.Println("**middles**")
		fmt.Println("source:", len(crate_columns[source_column]), "-", crate_columns[source_column])
		fmt.Println("destination:", len(crate_columns[destination_column]), "-", crate_columns[destination_column])

		crate_columns[source_column] = crate_columns[source_column][crate_quantity:]

		fmt.Println("**after**")
		fmt.Println("source:", len(crate_columns[source_column]), "-", crate_columns[source_column])
		fmt.Println("destination:", len(crate_columns[destination_column]), "-", crate_columns[destination_column])
		fmt.Println("")
		fmt.Println("")

	}

	str := ""
	for i := range crate_columns {
		str += string(crate_columns[i][0])
	}
	fmt.Println(str)

}
