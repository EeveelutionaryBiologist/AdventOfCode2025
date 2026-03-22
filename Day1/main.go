package main

import (
	"bufio"
	"fmt"
	"os"
)

type Command struct {
	Action rune
	Value  int16
}

func parse_puzzle_input(filePath string) ([]Command, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var commands []Command
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var action rune
		var val int16

		// Sscanf looks for a character followed by an integer
		_, err := fmt.Sscanf(line, "%c%d", &action, &val)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse line %q: %w", line, err)
		}

		commands = append(commands, Command{Action: action, Value: val})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}

func rotate_lock(position int16, action rune, value int16) (int16, uint16) {
	var crossed_zero uint16

	// Probably would be slightly more efficient with modulo shenanigans
	switch action {
	case 'L':
		for i := int16(0); i < value; i++ {
			position -= 1
			if position == 0 {
				crossed_zero += 1
			} else if position < 0 {
				position = 99
			}
		}
	case 'R':
		for i := int16(0); i < value; i++ {
			position += 1
			if position == 0 {
				crossed_zero += 1
			} else if position > 99 {
				position = 0
				crossed_zero += 1
			}
		}
	default:
		fmt.Println("Error", fmt.Errorf("Invalid commands: %c%d", action, value))
	}

	return position, crossed_zero
}

func main() {
	results, err := parse_puzzle_input("puzzle_input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var position int16 = 50
	var crossed_zero uint16 = 0
	var zeros uint16 = 0

	for _, cmd := range results {
		fmt.Printf("Position: %d Action: %c%d\n", position, cmd.Action, cmd.Value)
		position, crossed_zero = rotate_lock(position, cmd.Action, cmd.Value)
		zeros += crossed_zero
	}
	fmt.Printf("The zero position was crossed a total of %d times.\n", zeros)
}
