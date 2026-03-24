package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse_puzzle_input(filepath string) ([][]rune, error) {
	var matrix [][]rune

	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		row := []rune(line)
		matrix = append(matrix, row)
	}

	return matrix, nil
}

func count_neighbors(matrix [][]rune, i int, j int) uint16 {
	neighbors_i := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	neighbors_j := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	var neighbor_count uint16

	for n := 0; n < 8; n++ {
		this_i := i + neighbors_i[n]
		this_j := j + neighbors_j[n]

		if this_i < 0 || this_j < 0 {
			continue
		}
		if this_i >= len(matrix) || this_j >= len(matrix[0]) {
			continue
		}
		if matrix[this_i][this_j] == '@' {
			neighbor_count++
		}
	}

	return neighbor_count
}

func filter_by_neighbors(matrix [][]rune) uint32 {
	var count uint32

	for i, row := range matrix {
		for j, char := range row {
			if char == '.' {
				fmt.Printf("%c", char)
				continue
			}
			if count_neighbors(matrix, i, j) < 4 {
				fmt.Printf("x")
				count++
			} else {
				fmt.Printf("@")
			}
		}
		fmt.Println()
	}

	return count
}

func main() {
	matrix, err := parse_puzzle_input("puzzle_input.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	scrolls := filter_by_neighbors(matrix)
	fmt.Printf("Scrolls with less than 4 neighbors: %d\n", scrolls)
}
