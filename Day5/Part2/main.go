package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	id_ranges, err := parseInputRanges()

	if err != nil {
		fmt.Println("Error while parsing ranges:", err)
	}

	// Sort IDRange structs by lower value
	sort.Slice(id_ranges, func(i int, j int) bool {
		return id_ranges[i].low < id_ranges[j].low
	})

	var num_ids uint64
	merged := []IDRange{id_ranges[0]}

	// Approach: Iterate over sorted Ranges
	// Check Range i for overlap with i+1, adjust i+1 range, then add to merged
	for i := 1; i < len(id_ranges); i++ {
		current := id_ranges[i]
		last := &merged[len(merged)-1]
		fmt.Printf("%d-%d\n", current.low, current.high)

		if current.low <= last.high+1 {
			// Update the end of the last range if the current one extends it
			if current.high > last.high {
				last.high = current.high
			}
		} else {
			// No overlap, add as a new distinct range
			merged = append(merged, current)
		}
	}

	for _, rang := range merged {
		num_ids += (rang.high - rang.low) + 1
	}

	fmt.Printf("Total number of valid ids: %d\n", num_ids)
}
