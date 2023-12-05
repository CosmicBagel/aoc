package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
	maps defined by three numbers
	destination start; source start; range
*/

type IdMap struct {
	destination int
	source      int
	spread      int // range
}

type DataSetKey int

const (
	seedToSoil DataSetKey = iota
	soilToFertilizer
	fertilizerToWater
	waterToLight
	lightToTemperature
	temperatureToHumidity
	humidityToLocation
)

func Day5P1() {
	fmt.Println("day 5 p 1")
	// file_name := "example_input.txt"
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seeds, dataMap := parseInputP1(scanner)

	// process to lowest location that corresponds with *any* of the seed numbers
	lowestLocation := processSeedsAndMapsP1(seeds, dataMap)

	fmt.Printf("%+v\n", lowestLocation)
}

func processSeedsAndMapsP1(seeds []int, maps map[DataSetKey][]IdMap) int {
	lowestLocation := 2147483647

	allDataSetKeys := []DataSetKey{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight,
		lightToTemperature, temperatureToHumidity, humidityToLocation}

	for _, seed := range seeds {

		currentId := seed
		// fmt.Println(currentId)
		for _, key := range allDataSetKeys {
			// fmt.Printf("\t%v\n", key)
			var foundIdMap IdMap
			found := false
			// is id within any of this data set's maps
			for _, idMap := range maps[key] {
				if idMap.source <= currentId && idMap.source+idMap.spread >= currentId {
					foundIdMap = idMap
					found = true
					// fmt.Printf("\t\tcurrentId %v\n", currentId)
					// fmt.Printf("\t\tfound %+v\n", foundIdMap)
					break
				}
			}

			// if !found {
			// 	fmt.Printf("\t\tcurrentId %v\n", currentId)
			// 	fmt.Printf("\t\tnot found\n")
			// }

			if found {
				// minVal := min(foundIdMap.source, foundIdMap.destination)
				offset := currentId - foundIdMap.source
				destination := foundIdMap.destination + offset
				currentId = destination
				// fmt.Printf("\t\tresultId %v\n", destination)
			}

			// if no map, id maps as is, no change to currentId
		}

		if currentId < lowestLocation {
			lowestLocation = currentId
		}
	}

	return lowestLocation
}

func parseInputP1(scanner *bufio.Scanner) ([]int, map[DataSetKey][]IdMap) {
	dataMap := make(map[DataSetKey][]IdMap)
	seeds := make([]int, 0)

	headingToDataSet := func(s string) DataSetKey {
		var d DataSetKey
		switch s {
		case "seed-to-soil":
			d = seedToSoil
		case "soil-to-fertilizer":
			d = soilToFertilizer
		case "fertilizer-to-water":
			d = fertilizerToWater
		case "water-to-light":
			d = waterToLight
		case "light-to-temperature":
			d = lightToTemperature
		case "temperature-to-humidity":
			d = temperatureToHumidity
		case "humidity-to-location":
			d = humidityToLocation
		default:
			d = seedToSoil
		}
		return d
	}

	scanner.Scan()
	firstLine := scanner.Text()
	firstLineSplit := strings.Split(firstLine, ": ")
	seedSplit := strings.Split(firstLineSplit[1], " ")
	for _, s := range seedSplit {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, num)
	}

	var currentDataSet DataSetKey
	headerMode := false
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			headerMode = true
			continue
		}

		if headerMode {
			sSplit := strings.Split(s, " ")
			header := sSplit[0]
			currentDataSet = headingToDataSet(header)
			headerMode = false
			continue
		}

		//not header mode
		sSplit := strings.Split(s, " ")
		nums := make([]int, 0)
		for _, numStr := range sSplit {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}

		var idMap = IdMap{nums[0], nums[1], nums[2]}

		if dataMap[currentDataSet] == nil {
			dataMap[currentDataSet] = make([]IdMap, 0)
		}
		dataMap[currentDataSet] = append(dataMap[currentDataSet], idMap)
	}

	return seeds, dataMap
}

func Day5P2() {
	fmt.Println("day 5 p 2")
	file_name := "example_input.txt"
	// file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}

}
