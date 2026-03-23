package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IDRange struct {
	low  uint64
	high uint64
}

func parse_puzzle_input(filepath string) ([]IDRange, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var id_ranges []IDRange
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		range_strings := strings.Split(line, ",")

		for _, rstr := range range_strings {
			high_low := strings.Split(rstr, "-")

			if len(high_low) < 2 {
				fmt.Println("ERROR: There are not enough items in the range?")
			}
			num1, _ := strconv.Atoi(high_low[0])
			num2, _ := strconv.Atoi(high_low[1])
			id_ranges = append(id_ranges, IDRange{low: uint64(num1), high: uint64(num2)})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return id_ranges, nil
}

func sum(set map[uint64]struct{}) uint64 {
	var total uint64

	for el := range set {
		total += uint64(el)
	}
	return total
}

func is_invalid(val uint64) bool {
	valstr := strconv.FormatInt(int64(val), 10)
	length := len(valstr)

	// The invalid ids have to be a literal duplication of the same digits, so...
	// No longer true in Part II, but still used to trim the search space.
	if length%2 != 0 {
		if valstr[:length/2] == valstr[length/2:] {
			return true
		}
	}

	// Is the ID constructed from -any- combination of the same digits?
	for i := 1; i <= length/2; i++ {
		// We can skip here if length is not a multiple of current slice length
		if length%i != 0 {
			continue
		}
		base := valstr[:i]
		is_multiple := true

		for j := i; j < length; j = j + i {
			if base != valstr[j:j+i] {
				is_multiple = false
				break
			}
		}
		if is_multiple {
			return true
		}
	}

	return false
}

func find_faulty_ids(low uint64, high uint64) []uint64 {
	var invalid_ids []uint64

	for i := low; i <= high; i++ {
		if is_invalid(i) {
			invalid_ids = append(invalid_ids, i)
		}
	}
	return invalid_ids
}

func main() {
	id_ranges, err := parse_puzzle_input("puzzle_input.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	set := make(map[uint64]struct{})

	for _, idrange := range id_ranges {
		// fmt.Printf("%d-%d\n", idrange.low, idrange.high)
		invalid_ids := find_faulty_ids(idrange.low, idrange.high)

		for _, id := range invalid_ids {
			set[id] = struct{}{}
		}
	}

	fmt.Printf("Sum of invalid IDs: %d\n", sum(set))
}
