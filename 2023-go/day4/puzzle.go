package day4

import (
	"fmt"
	"math"
	"os"
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
	cards := parseFile()
	total := 0

	for _, card := range cards {
		if length := len(intersection(card.WinningNumbers, card.TrialNumbers)); length > 0 {
			total = total + int(math.Exp2(float64(length-1)))
		}
	}

	fmt.Println("Result of day 4 (part 1): ", total)
}

func solvePart2() {
	cards := parseFile()
	doubledCards := make(map[int]int)
	total := 0

	for index := range cards {
		doubledCards[index] = 1
	}

	for index, card := range cards {
		matches := intersection(card.WinningNumbers, card.TrialNumbers)

		for key := range matches {
			doubledCards[index+key+1] += doubledCards[index]
		}
	}

	for _, count := range doubledCards {
		total = total + count
	}

	fmt.Println("Result of day 4 (part 2): ", total)
}

type Card struct {
	WinningNumbers []int
	TrialNumbers   []int
}

func parseFile() []Card {
	file, err := os.ReadFile("./inputs/input4.txt")
	if err != nil {
		panic(err)
	}

	cards := make([]Card, 0)

	lines := strings.Split(string(file), "Card")[1:]
	for _, line := range lines {
		tmp := strings.Split(line, ":")
		tmp = strings.Split(tmp[1], " | ")

		cards = append(cards, Card{WinningNumbers: parseToIntegers(tmp[0]), TrialNumbers: parseToIntegers(tmp[1])})
	}

	return cards
}

func parseToIntegers(str string) []int {
	tmp := strings.Fields(str)
	result := make([]int, 0)

	for _, toParse := range tmp {
		number, err := strconv.Atoi(toParse)
		if err != nil {
			panic(err)
		}
		result = append(result, number)
	}

	return result
}

func intersection(a []int, b []int) []int {
	intersection := make([]int, 0)

	for _, numA := range a {
		if slices.Contains(b, numA) {
			intersection = append(intersection, numA)
		}
	}

	return intersection
}
