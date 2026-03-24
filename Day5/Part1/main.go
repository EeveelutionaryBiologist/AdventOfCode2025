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

func (rang IDRange) inRange(val uint64) bool {
	if val >= rang.low && val <= rang.high {
		return true
	}
	return false
}

func parseInputRanges() ([]IDRange, error) {
	file, err := os.Open("puzzle_input_1.txt")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var set []IDRange

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		vals := strings.Split(line, "-")

		if len(vals) != 2 {
			fmt.Println("Illegal line format? Skipping...")
			continue
		}
		low, _ := strconv.Atoi(vals[0])
		high, _ := strconv.Atoi(vals[1])

		set = append(set, IDRange{low: uint64(low), high: uint64(high)})
	}

	return set, nil
}

func parseInputIds() ([]uint64, error) {
	file, err := os.Open("puzzle_input_2.txt")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ids []uint64

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		val, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println("Error parsing line", err)
			continue
		}
		ids = append(ids, uint64(val))
	}

	return ids, nil
}

func main() {
	valid_set, err := parseInputRanges()

	if err != nil {
		fmt.Println("Error parsing input file 1: ", err)
	}
	ids, err := parseInputIds()

	if err != nil {
		fmt.Println("Error parsing input file 2: ", err)
	}
	var count uint32

	for _, id := range ids {
		for _, rang := range valid_set {
			ok := rang.inRange(id)
			if ok {
				count++
				break
			}
		}
	}
	fmt.Printf("Number of valid IDs: %d\n", count)
}
