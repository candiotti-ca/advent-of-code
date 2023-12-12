package main

import (
	"adventofcode/day1"
	"adventofcode/day2"
	"adventofcode/day3"
	"adventofcode/day4"
	"os"
	"strconv"
)

func main() {
	chosenDay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("parameter 'day' not valid")
	}

	if chosenDay < 1 || chosenDay > 25 {
		panic("the day has to be between 1 and 25")
	}

	chosenPart, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("parameter 'part' not valid")
	}
	if chosenPart < 1 || chosenPart > 2 {
		panic("the part has to be 1 or 2")
	}

	switch chosenDay {
	case 1:
		day1.Solve(chosenPart)
	case 2:
		day2.Solve(chosenPart)
	case 3:
		day3.Solve(chosenPart)
	case 4:
		day4.Solve(chosenPart)
	default:
		panic("day not supported")
	}
}
