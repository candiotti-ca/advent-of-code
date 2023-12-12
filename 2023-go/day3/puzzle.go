package day3

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Solve(part int) {
	if part == 1 {
		solvePart1()
	} else {
		solvePart2()
	}
}

func solvePart1() {
	total := 0
	lines := parseFile()

	for index, line := range lines {
		if index == 59 {
			fmt.Println("----------------")
		}
		for key, value := range line.Numbers {
			isNumberPartOfEngine := line.isNumberPartOfEngine(key, value)

			if !isNumberPartOfEngine && index > 0 {
				previousLine := lines[index-1]
				isNumberPartOfEngine = previousLine.isNumberPartOfEngine(key, value)
			}

			if !isNumberPartOfEngine && index < len(lines)-1 {
				nextLine := lines[index+1]
				isNumberPartOfEngine = nextLine.isNumberPartOfEngine(key, value)
			}

			if isNumberPartOfEngine {
				number, err := strconv.Atoi(key)
				if err != nil {
					panic("cannot convert number " + key)
				}

				total = total + number
			} else {
				// FIXME gerer les doublons par ligne
			}
		}
	}

	fmt.Println("Result of day 3 (part 1): ", total)
}

func solvePart2() {
	// total := 0
	// games := parseFile()

	// for _, game := range games {
	// 	requirements := game.getMinimalRequirement()
	// 	power := requirements.Blue * requirements.Red * requirements.Green
	// 	total = total + power
	// }

	// fmt.Println("Result of day 3 (part 2): ", total)
}

func parseFile() []Line {
	file, err := os.ReadFile("./inputs/input3.txt")
	if err != nil {
		panic(err)
	}

	lines := make([]Line, 0)
	digitRegxp := regexp.MustCompile("[0-9]+")
	for index, line := range strings.Split(string(file), "\n") {
		numbers := make(map[string]int, 0)
		allNbs := digitRegxp.FindAllString(line, -1)
		allInx := digitRegxp.FindAllStringIndex(line, -1)

		for indx, value := range allInx {
			numbers[allNbs[indx]] = value[0]
		}

		symbols := make([]int, 0)
		for j, rune := range line {
			char := string(rune)
			if !digitRegxp.MatchString(char) && char != "." {
				symbols = append(symbols, j)
			}
		}

		lines = append(lines, Line{Id: index, Symbols: symbols, Numbers: numbers})
	}

	return lines
}

type Line struct {
	Id      int
	Symbols []int
	Numbers map[string]int
}

func (l Line) isNumberPartOfEngine(numberAsString string, index int) bool {
	min := index - 1
	max := index + len(numberAsString)

	doesMatch := func(el int) bool {
		return el >= min && el <= max
	}

	return slices.ContainsFunc(l.Symbols, doesMatch)
}
