package day2

import (
	"fmt"
	"os"
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
	games := parseFile()

	for _, game := range games {
		if game.isPossible(12, 14, 13) {
			total = total + game.Id
		}
	}

	fmt.Println("Result of day 2 (part 1): ", total)
}

func solvePart2() {
	total := 0
	games := parseFile()

	for _, game := range games {
		requirements := game.getMinimalRequirement()
		power := requirements.Blue * requirements.Red * requirements.Green
		total = total + power
	}

	fmt.Println("Result of day 2 (part 2): ", total)
}

func parseFile() []Game {
	file, err := os.ReadFile("./inputs/input2.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	games := make([]Game, 0)
	for _, line := range lines {
		games = append(games, NewGame(line))
	}

	return games
}

type Set struct {
	Red   int
	Blue  int
	Green int
}

func NewSet(input string) Set {
	substr := strings.Split(input, ",")
	red := 0
	blue := 0
	green := 0

	for _, str := range substr {
		quantityStr := strings.Split(strings.Trim(str, " "), " ")[0]
		quantityInt, err := strconv.Atoi(quantityStr)
		if err != nil {
			panic("cannot get cube quantity")
		}

		if strings.Contains(str, "blue") {
			blue = quantityInt
		} else if strings.Contains(str, "red") {
			red = quantityInt
		} else {
			green = quantityInt
		}
	}

	return Set{Red: red, Blue: blue, Green: green}
}

type Game struct {
	Id   int
	Sets []Set
}

func NewGame(input string) Game {
	substr := strings.Split(input, ":")
	gameId := strings.Replace(substr[0], "Game ", "", 1)
	id, err := strconv.Atoi(gameId)
	if err != nil {
		panic("cannot get game id")
	}

	sets := make([]Set, 0)
	gameSets := strings.Split(substr[1], ";")
	for _, set := range gameSets {
		sets = append(sets, NewSet(set))
	}

	return Game{Id: id, Sets: sets}
}

func (g Game) isPossible(red int, blue int, green int) bool {
	for _, set := range g.Sets {
		if set.Red > red || set.Blue > blue || set.Green > green {
			return false
		}
	}

	return true
}

func (g Game) getMinimalRequirement() Set {
	red := 0
	blue := 0
	green := 0

	for _, set := range g.Sets {
		if set.Red > red {
			red = set.Red
		}

		if set.Blue > blue {
			blue = set.Blue
		}

		if set.Green > green {
			green = set.Green
		}
	}

	return Set{Red: red, Blue: blue, Green: green}
}
