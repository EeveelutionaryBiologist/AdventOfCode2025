package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse_puzzle_input(filepath string) ([][]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var banks [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		values := make([]int, len(line))

		for i, char := range line {
			values[i] = int(char - '0')
		}
		banks = append(banks, values)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return banks, nil
}

func find_maximum(bank []int) int {
	current_max := 0

	for i := 0; i < len(bank)-1; i++ {
		for j := i + 1; j < len(bank); j++ {
			val := (bank[i] * 10) + bank[j]

			if val > current_max {
				current_max = val
			}
		}
	}

	return current_max
}

func main() {
	banks, err := parse_puzzle_input("puzzle_input.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var sum_voltage int

	for _, values := range banks {
		for _, i := range values {
			fmt.Printf("%d", i)
		}
		max_voltage := find_maximum(values)
		sum_voltage += max_voltage

		fmt.Printf(" -> %d", max_voltage)
		fmt.Println()
	}
	fmt.Printf("Sum of voltages: %d\n", sum_voltage)
}
