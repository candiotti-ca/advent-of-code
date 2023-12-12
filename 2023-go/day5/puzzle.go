package day5

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
	almanach := parseFile()
	result := -1

	for _, seed := range almanach.seeds {
		location := almanach.getSeedLocation(seed)
		if result == -1 || location < result {
			result = location
		}
	}

	fmt.Println("Result of day 5 (part 1): ", result)
}

func solvePart2() {
	// almanach := parseFile()
	// fmt.Printf("len(almanach.seeds): %v\n", len(almanach.seeds))
	// // result := -1

	// min := 0
	// for i, seed := range almanach.seeds {
	// 	if i%2 == 0 {
	// 		min = seed
	// 		continue
	// 	}

	// 	for j :=min;j<
	// }

	fmt.Println("Result of day 5 (part 2): ", 0)
}

func parseFile() Almanach {
	input, err := os.ReadFile("./inputs/input5.txt")
	if err != nil {
		panic(err)
	}

	seeds := make([]int, 0)
	seedToSoil := NewRecords()
	soilToFertilizer := NewRecords()
	fertilizerToWater := NewRecords()
	waterToLight := NewRecords()
	lightToTemperature := NewRecords()
	temperatureToHumidity := NewRecords()
	humidityToLocation := NewRecords()

	lines := strings.Split(string(input), "\n")
	cursor := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, " map:") {
			cursor += 1
			continue
		}

		if cursor == 0 {
			seeds = extractSeeds(line)
			continue
		}

		record := NewRecord(line)
		switch cursor {
		case 1:
			seedToSoil.addRecord(record)
		case 2:
			soilToFertilizer.addRecord(record)
		case 3:
			fertilizerToWater.addRecord(record)
		case 4:
			waterToLight.addRecord(record)
		case 5:
			lightToTemperature.addRecord(record)
		case 6:
			temperatureToHumidity.addRecord(record)
		case 7:
			humidityToLocation.addRecord(record)
		}
	}

	return Almanach{
		seeds,
		seedToSoil,
		soilToFertilizer,
		fertilizerToWater,
		waterToLight,
		lightToTemperature,
		temperatureToHumidity,
		humidityToLocation,
	}
}

func parseInt(str string) int {
	converted, err := strconv.Atoi(str)
	if err != nil {
		panic("cannot convert <" + str + "> to integer")
	}

	return converted
}

func extractSeeds(line string) []int {
	numbers := strings.Split(line, "seeds:")[1]
	numbers = strings.Trim(numbers, " ")

	seeds := make([]int, 0)
	for _, number := range strings.Split(numbers, " ") {
		seeds = append(seeds, parseInt(number))
	}

	return seeds
}

type Record struct {
	SourceMin   int
	SourceMax   int
	Destination int
}

func (r Record) String() string {
	return fmt.Sprintf("From seed %v to %v, destination begins at %v", r.SourceMin, r.SourceMax, r.Destination)
}

func NewRecord(line string) Record {
	columns := strings.Split(line, " ")
	destination := parseInt(columns[0])
	min := parseInt(columns[1])
	max := min + parseInt(columns[2])

	return Record{
		SourceMin:   min,
		SourceMax:   max,
		Destination: destination,
	}
}

type Records struct {
	List []Record
}

func (records *Records) addRecord(r Record) {
	records.List = append(records.List, r)
}

func (r Record) isIncluded(i int) bool {
	return i >= r.SourceMin && i <= r.SourceMax
}

func NewRecords() Records {
	return Records{List: make([]Record, 0)}
}

func (records Records) getDestination(i int) int {
	for _, r := range records.List {
		if r.isIncluded(i) {
			return r.Destination + i - r.SourceMin
		}
	}

	return i
}

type Almanach struct {
	seeds                 []int
	seedToSoil            Records
	soilToFertilizer      Records
	fertilizerToWater     Records
	waterToLight          Records
	lightToTemperature    Records
	temperatureToHumidity Records
	humidityToLocation    Records
}

func (a Almanach) getSeedLocation(seed int) int {
	soil := a.seedToSoil.getDestination(seed)
	fertilizer := a.soilToFertilizer.getDestination(soil)
	water := a.fertilizerToWater.getDestination(fertilizer)
	light := a.waterToLight.getDestination(water)
	temperature := a.lightToTemperature.getDestination(light)
	humidity := a.temperatureToHumidity.getDestination(temperature)
	location := a.humidityToLocation.getDestination(humidity)

	return location
}

// func e() {
// 	file, err := os.Open("")
// 	if err != nil {
// 		panic("cannot open file")
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	lines := make([]string, 0)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}

// 	if scanner.Err() != nil {
// 		panic("cannot read file")
// 	}
// }
