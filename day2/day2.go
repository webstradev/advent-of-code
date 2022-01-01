package day2

import (
	"log"
	"strconv"
	"strings"
)

var input = []string{
	"forward 2",
	"down 8",
	"down 1",
	"up 7",
	"forward 4",
	"down 4",
	"down 3",
	"forward 2",
	"forward 5",
	"forward 2",
	"forward 7",
	"forward 8",
	"down 7",
	"forward 6",
	"forward 1",
	"down 8",
	"down 8",
	"up 9",
	"down 9",
	"forward 1",
	"up 1",
	"forward 6",
	"forward 7",
	"down 7",
	"forward 4",
	"forward 2",
	"forward 1",
	"forward 3",
	"forward 8",
	"forward 4",
	"up 6",
	"down 2",
	"forward 4",
	"down 3",
	"up 3",
	"up 2",
	"up 3",
	"forward 8",
	"down 5",
	"down 8",
	"down 5",
	"down 4",
	"down 2",
	"up 1",
	"forward 6",
	"forward 6",
	"forward 8",
	"up 5",
	"down 5",
	"forward 8",
	"forward 7",
	"down 9",
	"down 7",
	"down 6",
	"down 3",
	"forward 3",
	"up 7",
	"forward 2",
	"up 6",
	"forward 7",
	"forward 9",
	"down 9",
	"forward 3",
	"down 2",
	"down 2",
	"down 7",
	"down 7",
	"forward 8",
	"down 7",
	"forward 9",
	"up 7",
	"down 4",
	"down 8",
	"forward 2",
	"forward 2",
	"up 1",
	"forward 8",
	"down 5",
	"forward 8",
	"forward 4",
	"up 1",
	"forward 2",
	"forward 2",
	"forward 2",
	"down 9",
	"down 7",
	"down 9",
	"forward 9",
	"forward 4",
	"up 7",
	"down 4",
	"forward 9",
	"forward 8",
	"up 5",
	"up 1",
	"down 1",
	"down 9",
	"down 6",
	"up 8",
	"forward 2",
	"forward 7",
	"up 7",
	"forward 7",
	"forward 6",
	"down 6",
	"forward 8",
	"up 6",
	"forward 7",
	"down 7",
	"forward 4",
	"forward 9",
	"up 9",
	"up 8",
	"forward 8",
	"forward 3",
	"up 1",
	"up 4",
	"down 4",
	"up 9",
	"up 8",
	"forward 6",
	"down 2",
	"down 2",
	"up 4",
	"forward 4",
	"forward 1",
	"down 8",
	"forward 6",
	"down 5",
	"forward 6",
	"down 8",
	"up 1",
	"forward 1",
	"up 9",
	"down 8",
	"up 3",
	"up 9",
	"forward 9",
	"forward 2",
	"down 2",
	"up 6",
	"up 6",
	"forward 8",
	"up 3",
	"down 3",
	"forward 5",
	"up 2",
	"forward 1",
	"forward 1",
	"forward 8",
	"down 4",
	"forward 8",
	"forward 7",
	"down 8",
	"forward 7",
	"down 7",
	"down 2",
	"down 1",
	"down 7",
	"up 9",
	"down 5",
	"up 1",
	"forward 7",
	"down 5",
	"down 9",
	"down 2",
	"down 8",
	"down 4",
	"forward 2",
	"forward 2",
	"forward 1",
	"down 6",
	"up 7",
	"forward 2",
	"down 1",
	"down 5",
	"forward 3",
	"forward 8",
	"down 4",
	"up 2",
	"up 9",
	"up 7",
	"forward 7",
	"forward 4",
	"up 8",
	"up 3",
	"up 4",
	"forward 6",
	"down 7",
	"forward 7",
	"up 6",
	"down 9",
	"up 6",
	"forward 4",
	"up 3",
	"down 3",
	"up 6",
	"down 9",
	"down 6",
	"forward 7",
	"forward 9",
	"forward 2",
	"down 6",
	"up 3",
	"up 9",
	"forward 9",
	"forward 8",
	"up 4",
	"up 5",
	"forward 6",
	"down 5",
	"up 4",
	"up 9",
	"down 8",
	"forward 5",
	"up 5",
	"forward 7",
	"forward 6",
	"down 8",
	"down 5",
	"down 2",
	"up 5",
	"down 8",
	"forward 7",
	"forward 7",
	"up 6",
	"forward 3",
	"down 8",
	"forward 5",
	"forward 8",
	"down 7",
	"forward 6",
	"down 2",
	"forward 2",
	"forward 8",
	"down 4",
	"forward 5",
	"forward 7",
	"forward 5",
	"forward 8",
	"forward 9",
	"forward 6",
	"down 6",
	"up 7",
	"down 9",
	"forward 7",
	"forward 1",
	"up 2",
	"forward 8",
	"down 1",
	"up 9",
	"forward 7",
	"down 2",
	"up 2",
	"forward 5",
	"down 4",
	"down 3",
	"up 8",
	"up 6",
	"forward 3",
	"up 4",
	"forward 3",
	"forward 1",
	"forward 1",
	"up 9",
	"down 2",
	"down 9",
	"up 4",
	"forward 4",
	"forward 4",
	"forward 2",
	"forward 5",
	"forward 9",
	"forward 7",
	"up 4",
	"up 5",
	"down 5",
	"forward 9",
	"down 1",
	"forward 2",
	"down 6",
	"forward 9",
	"down 4",
	"down 2",
	"forward 2",
	"up 7",
	"forward 7",
	"forward 8",
	"up 9",
	"forward 3",
	"up 1",
	"down 5",
	"forward 5",
	"down 1",
	"up 5",
	"forward 4",
	"forward 5",
	"up 3",
	"down 4",
	"up 6",
	"up 1",
	"down 2",
	"forward 6",
	"down 8",
	"up 9",
	"down 7",
	"forward 9",
	"down 8",
	"forward 3",
	"forward 8",
	"down 1",
	"down 7",
	"forward 6",
	"up 6",
	"down 1",
	"down 5",
	"forward 6",
	"down 7",
	"down 2",
	"forward 6",
	"forward 2",
	"forward 8",
	"forward 7",
	"forward 2",
	"down 7",
	"up 2",
	"down 7",
	"forward 8",
	"forward 8",
	"forward 2",
	"forward 4",
	"down 1",
	"down 6",
	"down 1",
	"down 4",
	"down 4",
	"down 2",
	"down 7",
	"up 5",
	"up 5",
	"down 1",
	"forward 3",
	"up 1",
	"down 3",
	"forward 9",
	"forward 4",
	"forward 7",
	"down 4",
	"down 4",
	"down 2",
	"forward 2",
	"forward 1",
	"forward 9",
	"down 1",
	"down 4",
	"down 1",
	"forward 4",
	"up 8",
	"forward 3",
	"down 6",
	"forward 5",
	"forward 9",
	"forward 1",
	"up 8",
	"down 7",
	"down 8",
	"forward 4",
	"down 4",
	"up 5",
	"down 3",
	"forward 3",
	"down 6",
	"down 1",
	"down 9",
	"forward 8",
	"up 5",
	"down 7",
	"up 7",
	"forward 3",
	"up 5",
	"up 7",
	"down 4",
	"up 2",
	"down 2",
	"down 8",
	"up 6",
	"down 4",
	"up 1",
	"down 8",
	"down 4",
	"forward 7",
	"down 4",
	"up 3",
	"down 9",
	"forward 3",
	"up 4",
	"up 1",
	"up 5",
	"down 4",
	"forward 4",
	"forward 4",
	"down 2",
	"down 3",
	"down 4",
	"forward 2",
	"down 9",
	"down 9",
	"down 6",
	"forward 5",
	"forward 7",
	"down 3",
	"forward 2",
	"up 8",
	"down 6",
	"down 8",
	"down 2",
	"up 9",
	"down 1",
	"forward 8",
	"forward 1",
	"forward 8",
	"up 4",
	"up 1",
	"down 8",
	"up 2",
	"forward 5",
	"down 3",
	"forward 5",
	"forward 4",
	"forward 1",
	"down 9",
	"forward 2",
	"forward 5",
	"forward 4",
	"forward 3",
	"down 7",
	"up 9",
	"forward 7",
	"up 8",
	"forward 3",
	"forward 3",
	"down 5",
	"forward 7",
	"forward 5",
	"down 4",
	"forward 6",
	"down 6",
	"up 3",
	"down 3",
	"forward 7",
	"forward 8",
	"up 3",
	"up 8",
	"forward 8",
	"up 4",
	"forward 5",
	"up 4",
	"down 7",
	"forward 4",
	"down 8",
	"down 6",
	"forward 7",
	"down 4",
	"up 7",
	"down 6",
	"up 5",
	"down 2",
	"forward 7",
	"down 7",
	"up 3",
	"up 3",
	"down 3",
	"down 5",
	"forward 3",
	"forward 5",
	"forward 8",
	"down 3",
	"down 8",
	"forward 3",
	"down 1",
	"forward 4",
	"up 5",
	"forward 8",
	"up 1",
	"up 4",
	"down 4",
	"forward 1",
	"forward 6",
	"up 7",
	"up 1",
	"up 7",
	"down 5",
	"forward 8",
	"down 1",
	"forward 1",
	"forward 6",
	"up 9",
	"up 6",
	"down 4",
	"forward 5",
	"down 1",
	"forward 7",
	"down 8",
	"up 3",
	"down 9",
	"down 1",
	"up 7",
	"up 6",
	"forward 8",
	"down 9",
	"down 4",
	"forward 4",
	"up 7",
	"down 3",
	"forward 5",
	"forward 8",
	"up 5",
	"down 7",
	"up 2",
	"forward 3",
	"forward 3",
	"forward 4",
	"down 1",
	"forward 5",
	"forward 7",
	"up 7",
	"forward 7",
	"forward 1",
	"forward 3",
	"forward 9",
	"down 3",
	"forward 5",
	"down 9",
	"down 7",
	"down 4",
	"down 8",
	"up 3",
	"down 1",
	"up 5",
	"down 5",
	"forward 2",
	"down 6",
	"down 5",
	"forward 2",
	"forward 6",
	"down 6",
	"up 6",
	"down 8",
	"up 1",
	"forward 4",
	"forward 2",
	"forward 1",
	"up 1",
	"forward 5",
	"forward 3",
	"forward 7",
	"down 9",
	"forward 3",
	"down 9",
	"up 5",
	"down 7",
	"down 2",
	"down 3",
	"up 5",
	"up 3",
	"down 2",
	"forward 4",
	"forward 1",
	"down 3",
	"up 4",
	"down 8",
	"down 1",
	"forward 6",
	"down 3",
	"forward 9",
	"down 7",
	"down 3",
	"down 4",
	"down 5",
	"up 1",
	"forward 8",
	"forward 2",
	"up 3",
	"up 7",
	"up 1",
	"forward 3",
	"forward 9",
	"up 5",
	"forward 6",
	"down 8",
	"down 9",
	"down 5",
	"forward 6",
	"forward 3",
	"forward 8",
	"forward 7",
	"forward 9",
	"forward 2",
	"up 9",
	"forward 8",
	"down 2",
	"down 6",
	"down 9",
	"down 5",
	"forward 1",
	"down 7",
	"forward 2",
	"down 4",
	"down 1",
	"up 1",
	"down 2",
	"forward 5",
	"forward 8",
	"down 9",
	"up 6",
	"forward 8",
	"forward 3",
	"up 6",
	"up 2",
	"forward 8",
	"down 8",
	"up 1",
	"down 6",
	"down 2",
	"down 6",
	"forward 5",
	"forward 6",
	"down 4",
	"forward 1",
	"down 4",
	"up 7",
	"forward 4",
	"up 1",
	"forward 9",
	"forward 7",
	"up 3",
	"down 7",
	"down 4",
	"forward 3",
	"forward 8",
	"forward 6",
	"forward 6",
	"up 1",
	"up 2",
	"up 8",
	"down 7",
	"forward 8",
	"down 5",
	"forward 6",
	"down 1",
	"down 9",
	"down 2",
	"forward 6",
	"up 3",
	"down 5",
	"down 6",
	"forward 1",
	"forward 2",
	"down 3",
	"forward 8",
	"forward 1",
	"up 5",
	"down 8",
	"down 4",
	"up 9",
	"up 2",
	"forward 7",
	"forward 9",
	"up 8",
	"up 6",
	"forward 1",
	"down 7",
	"up 1",
	"down 3",
	"forward 2",
	"forward 3",
	"down 3",
	"down 2",
	"forward 3",
	"down 7",
	"forward 3",
	"forward 7",
	"forward 7",
	"down 8",
	"down 4",
	"forward 3",
	"forward 4",
	"down 7",
	"down 9",
	"down 4",
	"down 2",
	"forward 7",
	"up 8",
	"down 4",
	"down 3",
	"forward 9",
	"down 5",
	"up 6",
	"up 2",
	"down 5",
	"down 6",
	"forward 2",
	"forward 8",
	"down 1",
	"forward 6",
	"up 7",
	"down 6",
	"forward 4",
	"down 2",
	"down 5",
	"down 9",
	"forward 7",
	"up 4",
	"forward 9",
	"up 7",
	"down 4",
	"down 6",
	"up 9",
	"forward 1",
	"up 7",
	"down 5",
	"forward 3",
	"forward 3",
	"down 7",
	"down 1",
	"down 7",
	"down 7",
	"down 1",
	"forward 8",
	"forward 9",
	"forward 8",
	"down 9",
	"down 8",
	"down 5",
	"down 3",
	"forward 4",
	"forward 1",
	"down 1",
	"forward 1",
	"down 7",
	"forward 7",
	"forward 3",
	"down 8",
	"forward 3",
	"forward 9",
	"forward 8",
	"down 7",
	"forward 8",
	"down 2",
	"up 5",
	"forward 7",
	"forward 4",
	"down 2",
	"up 6",
	"up 8",
	"forward 7",
	"down 9",
	"up 3",
	"forward 4",
	"up 9",
	"up 5",
	"up 5",
	"up 6",
	"down 4",
	"down 5",
	"up 5",
	"forward 7",
	"forward 6",
	"down 4",
	"forward 5",
	"forward 4",
	"up 7",
	"forward 8",
	"down 2",
	"forward 1",
	"down 9",
	"down 8",
	"forward 7",
	"down 7",
	"down 1",
	"forward 9",
	"down 7",
	"forward 6",
	"down 2",
	"up 9",
	"forward 1",
	"up 9",
	"down 3",
	"up 9",
	"down 1",
	"forward 1",
	"down 1",
	"up 6",
	"down 8",
	"up 2",
	"down 3",
	"forward 1",
	"down 4",
	"up 5",
	"down 5",
	"down 4",
	"forward 5",
	"forward 4",
	"down 9",
	"up 7",
	"down 7",
	"forward 7",
	"forward 6",
	"forward 8",
	"down 8",
	"forward 8",
	"down 1",
	"down 1",
	"down 8",
	"down 2",
	"up 2",
	"up 1",
	"forward 5",
	"down 1",
	"up 5",
	"up 2",
	"down 6",
	"up 8",
	"forward 5",
	"down 8",
	"down 1",
	"up 5",
	"down 1",
	"forward 4",
	"down 6",
	"down 4",
	"forward 2",
	"forward 2",
	"down 1",
	"up 4",
	"up 8",
	"down 6",
	"down 2",
	"forward 5",
	"forward 8",
	"forward 7",
	"down 5",
	"down 7",
	"down 3",
	"forward 6",
	"down 3",
	"down 3",
	"forward 6",
	"forward 6",
	"forward 6",
	"up 7",
	"forward 1",
	"down 5",
	"down 2",
	"up 8",
	"forward 6",
	"down 7",
	"down 6",
	"forward 1",
	"up 5",
	"down 4",
	"up 9",
	"forward 3",
	"up 3",
	"forward 9",
	"forward 9",
	"forward 7",
	"forward 5",
	"down 9",
	"forward 1",
	"forward 6",
	"up 8",
	"down 7",
	"forward 9",
	"forward 5",
	"up 4",
	"down 8",
	"forward 8",
	"forward 4",
	"down 9",
	"up 2",
	"forward 5",
	"forward 8",
	"down 8",
	"down 9",
	"down 9",
	"forward 4",
	"forward 8",
	"down 5",
	"down 5",
	"forward 5",
	"forward 5",
	"up 9",
	"up 7",
	"forward 3",
	"up 4",
	"down 8",
	"up 6",
	"up 6",
	"down 4",
	"down 3",
	"forward 2",
	"forward 9",
	"down 4",
	"down 2",
	"forward 4",
	"up 9",
	"forward 1",
	"down 8",
	"down 9",
	"down 9",
	"down 3",
	"forward 1",
	"down 5",
	"up 9",
	"forward 6",
	"up 4",
	"forward 2",
	"forward 2",
	"forward 4",
	"down 9",
	"up 5",
	"up 1",
	"down 6",
	"forward 7",
	"down 8",
	"forward 4",
	"forward 9",
	"up 9",
	"up 4",
	"down 5",
	"down 3",
	"forward 2",
	"down 8",
	"down 6",
	"forward 3",
	"down 2",
	"forward 6",
	"up 2",
	"forward 6",
	"down 7",
	"up 4",
	"forward 1",
	"forward 4",
	"up 4",
	"forward 6",
	"forward 8",
	"down 7",
	"down 6",
	"up 7",
	"down 2",
	"down 4",
	"down 5",
	"forward 1",
	"up 4",
	"forward 8",
	"forward 6",
	"down 8",
	"up 5",
	"up 2",
	"up 9",
	"up 5",
	"forward 6",
	"down 4",
	"up 3",
	"down 8",
	"down 6",
	"down 2",
	"up 3",
	"up 5",
	"down 1",
	"forward 9",
	"up 8",
	"up 2",
	"down 3",
	"forward 6",
	"down 1",
	"forward 5",
	"down 3",
	"up 1",
	"up 2",
	"down 5",
	"down 7",
	"forward 8",
	"down 8",
	"up 9",
	"forward 3",
	"down 8",
	"down 8",
	"forward 1",
	"down 4",
	"down 4",
	"forward 3",
	"up 6",
	"down 3",
	"down 7",
	"down 7",
	"up 1",
	"forward 3",
	"forward 2",
}

// Command is the type for a parsed and validated command
type Command struct {
	Direction string
	Step      int
}

// validateDirectionString check if string part is valid
func validateDirectionString(input string) bool {
	switch input {
	case "forward", "down", "up":
		return true
	default:
		return false
	}
}

// Converts a string command into a struct
func createCommandFromString(input string) *Command {
	// Split input on the space seperator
	s := strings.Split(input, " ")
	if len(s) != 2 {
		return nil
	}

	// Check if first part is a valid direction
	if !validateDirectionString(s[0]) {
		return nil
	}

	// Parse second string to int
	stepSize, err := strconv.Atoi(s[1])
	if err != nil {
		log.Printf("error parsing int from string: %s", err)
		return nil
	}

	return &Command{Direction: s[0], Step: stepSize}
}

// calculateCommandResults caluclates the postion and depth of the submarine given a list of commands
func calculateCommandResults(commands []string) (int, int) {
	position, depth := 0, 0

	for _, commandString := range commands {
		command := createCommandFromString(commandString)
		if command == nil {
			log.Printf("Failed to parse '%s' into a valid command. command ignored", commandString)
		} else {
			switch command.Direction {
			case "forward":
				position += command.Step
			case "up":
				if depth-command.Step < 0 {
					depth = 0
				} else {
					depth -= command.Step
				}
			case "down":
				depth += command.Step
			}
		}

	}

	return position, depth
}

// calculateNewCommandResults calculates the postion and depth of the submarine using aim method for a given list of commands
func calculateNewCommandResults(commands []string) (int, int) {
	position, depth, aim := 0, 0, 0

	for _, commandString := range commands {
		command := createCommandFromString(commandString)
		if command == nil {
			log.Printf("Failed to parse '%s' into a valid command. command ignored", commandString)
		} else {
			switch command.Direction {
			case "forward":
				position += command.Step
				depth += command.Step * aim
			case "up":
				aim -= command.Step
			case "down":
				aim += command.Step
			}
		}

	}

	return position, depth
}

type Position struct {
	Horizontal int
	Depth      int
}

func Solve() {
	x, y := calculateCommandResults(input)
	xNew, yNew := calculateNewCommandResults(input)
	log.Println("----------")
	log.Println("Day 2:")
	log.Printf("day2puzzle1 - Multiplying the coordinates of the final postion gives: %d", x*y)
	log.Printf("day2puzzle2 - Multiplying the coordinates of the final postion gives: %d", xNew*yNew)
	log.Print("----------")
}
