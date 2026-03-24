package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// For Part I: Change this to 2
const N_DIGITS = 12

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
	left_index := 0
	var selected_values [12]int

	// Since we string up values left->right, the first maximum digit within the
	// allowed window will necessarily make the resulting number larger or equal to any
	// possible alternative.
	for n := N_DIGITS; n > 0; n-- {
		highest_digit := 0

		for i := left_index; i <= len(bank)-n; i++ {
			if bank[i] > highest_digit {
				highest_digit = bank[i]
				left_index = i + 1
			}
		}
		selected_values[N_DIGITS-n] = highest_digit
	}
	var total int

	// The selected digits are just like ordered decimal numbers
	for n := 0; n < N_DIGITS; n++ {
		power := math.Pow10(12 - (n + 1))
		total += int(float64(selected_values[n]) * power)
	}

	return total
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
