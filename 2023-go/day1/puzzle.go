package day1

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func Solve(part int) {
	if part == 1 {
		solvePart1()
	} else {
		solvePart2()
	}
}

func solvePart1() {
	lines := parseFile()
	total := 0

	for _, line := range lines {
		first, last := getFirstLastIndexesOfDigits(line)

		numberAsString := string(line[first]) + string(line[last])
		parsedNumber, err := strconv.Atoi(numberAsString)
		if err != nil {
			panic(err)
		}

		total = total + parsedNumber
	}

	fmt.Println("Result of day 1 (part 1): ", total)
}

func solvePart2() {
	lines := parseFile()
	total := 0

	for _, line := range lines {
		firstDigit, lastDigit := getFirstLastIndexesOfDigits(line)
		var leftNumber string
		var rightNumber string

		indx, literal := getFirstIndexOfLiteralDigit(line)
		if indx > -1 && indx < firstDigit {
			leftNumber = parseLiteralDigit(literal)
		} else {
			leftNumber = string(line[firstDigit])
		}

		indx, literal = getLastIndexOfLiteralDigit(line)
		if indx > lastDigit {
			rightNumber = parseLiteralDigit(literal)
		} else {
			rightNumber = string(line[lastDigit])
		}

		parsedNumber, err := strconv.Atoi(leftNumber + rightNumber)
		if err != nil {
			panic(err)
		}

		total = total + parsedNumber
	}

	fmt.Println("Result of day 1 (part 2): ", total)
}

func parseFile() []string {
	file, err := os.ReadFile("./inputs/input1.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	return lines
}

func getFirstLastIndexesOfDigits(line string) (int, int) {
	re := regexp.MustCompile("[0-9]")
	result := make([]int, 0)

	for i := 0; i < len(line); i++ {
		if re.MatchString(string(line[i])) {
			result = append(result, i)
		}
	}

	return slices.Min(result), slices.Max(result)
}

func getFirstIndexOfLiteralDigit(line string) (int, string) {
	literals := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	indexes := make(map[int]string, 0)

	for _, literal := range literals {
		indx := strings.Index(line, literal)
		if indx >= 0 {
			indexes[indx] = literal
		}
	}

	if len(indexes) == 0 {
		return -1, ""
	}

	min := slices.Min(maps.Keys(indexes))
	return min, indexes[min]
}

func getLastIndexOfLiteralDigit(line string) (int, string) {
	literals := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	indexes := make(map[int]string, 0)

	for _, literal := range literals {
		indx := strings.LastIndex(line, literal)
		if indx >= 0 {
			indexes[indx] = literal
		}
	}

	if len(indexes) == 0 {
		return -1, ""
	}

	max := slices.Max(maps.Keys(indexes))
	return max, indexes[max]
}

func parseLiteralDigit(literal string) string {
	switch literal {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		panic("not a valid digit")
	}
}
