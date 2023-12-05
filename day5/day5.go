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
	destination uint64
	source      uint64
	spread      uint64 // range
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

// func Day5P1() {
// 	fmt.Println("day 5 p 1")
// 	// file_name := "example_input.txt"
// 	file_name := "input.txt"
//
// 	file, err := os.Open(file_name)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
//
// 	scanner := bufio.NewScanner(file)
//
// 	seeds, dataMap := parseInputP1(scanner)
//
// 	// process to lowest location that corresponds with *any* of the seed numbers
// 	lowestLocation := processSeedsAndMapsP1(seeds, dataMap)
//
// 	fmt.Printf("%+v\n", lowestLocation)
// }
//
// func processSeedsAndMapsP1(seeds []int, maps map[DataSetKey][]IdMap) int {
// 	lowestLocation := 2147483647
//
// 	allDataSetKeys := []DataSetKey{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight,
// 		lightToTemperature, temperatureToHumidity, humidityToLocation}
//
// 	for _, seed := range seeds {
//
// 		currentId := seed
// 		// fmt.Println(currentId)
// 		for _, key := range allDataSetKeys {
// 			// fmt.Printf("\t%v\n", key)
// 			var foundIdMap IdMap
// 			found := false
// 			// is id within any of this data set's maps
// 			for _, idMap := range maps[key] {
// 				if idMap.source <= currentId && idMap.source+idMap.spread >= currentId {
// 					foundIdMap = idMap
// 					found = true
// 					// fmt.Printf("\t\tcurrentId %v\n", currentId)
// 					// fmt.Printf("\t\tfound %+v\n", foundIdMap)
// 					break
// 				}
// 			}
//
// 			// if !found {
// 			// 	fmt.Printf("\t\tcurrentId %v\n", currentId)
// 			// 	fmt.Printf("\t\tnot found\n")
// 			// }
//
// 			if found {
// 				// minVal := min(foundIdMap.source, foundIdMap.destination)
// 				offset := currentId - foundIdMap.source
// 				destination := foundIdMap.destination + offset
// 				currentId = destination
// 				// fmt.Printf("\t\tresultId %v\n", destination)
// 			}
//
// 			// if no map, id maps as is, no change to currentId
// 		}
//
// 		if currentId < lowestLocation {
// 			lowestLocation = currentId
// 		}
// 	}
//
// 	return lowestLocation
// }
//
// func parseInputP1(scanner *bufio.Scanner) ([]int, map[DataSetKey][]IdMap) {
// 	dataMap := make(map[DataSetKey][]IdMap)
// 	seeds := make([]int, 0)
//
// 	headingToDataSet := func(s string) DataSetKey {
// 		var d DataSetKey
// 		switch s {
// 		case "seed-to-soil":
// 			d = seedToSoil
// 		case "soil-to-fertilizer":
// 			d = soilToFertilizer
// 		case "fertilizer-to-water":
// 			d = fertilizerToWater
// 		case "water-to-light":
// 			d = waterToLight
// 		case "light-to-temperature":
// 			d = lightToTemperature
// 		case "temperature-to-humidity":
// 			d = temperatureToHumidity
// 		case "humidity-to-location":
// 			d = humidityToLocation
// 		default:
// 			d = seedToSoil
// 		}
// 		return d
// 	}
//
// 	scanner.Scan()
// 	firstLine := scanner.Text()
// 	firstLineSplit := strings.Split(firstLine, ": ")
// 	seedSplit := strings.Split(firstLineSplit[1], " ")
// 	for _, s := range seedSplit {
// 		num, err := strconv.Atoi(s)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		seeds = append(seeds, num)
// 	}
//
// 	var currentDataSet DataSetKey
// 	headerMode := false
// 	for scanner.Scan() {
// 		s := scanner.Text()
// 		if len(s) == 0 {
// 			headerMode = true
// 			continue
// 		}
//
// 		if headerMode {
// 			sSplit := strings.Split(s, " ")
// 			header := sSplit[0]
// 			currentDataSet = headingToDataSet(header)
// 			headerMode = false
// 			continue
// 		}
//
// 		//not header mode
// 		sSplit := strings.Split(s, " ")
// 		nums := make([]int, 0)
// 		for _, numStr := range sSplit {
// 			num, err := strconv.Atoi(numStr)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			nums = append(nums, num)
// 		}
//
// 		var idMap = IdMap{nums[0], nums[1], nums[2]}
//
// 		if dataMap[currentDataSet] == nil {
// 			dataMap[currentDataSet] = make([]IdMap, 0)
// 		}
// 		dataMap[currentDataSet] = append(dataMap[currentDataSet], idMap)
// 	}
//
// 	return seeds, dataMap
// }

func Day5P2() {
	fmt.Println("day 5 p 2")
	fmt.Println("note: 7873085 is incorrect (too high)")
	// file_name := "example_input.txt"
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seeds, dataMap := parseInputP2(scanner)

	// process to lowest location that corresponds with *any* of the seed numbers
	lowestLocation := processSeedsAndMapsP2(seeds, dataMap)

	fmt.Printf("%+v\n", lowestLocation)
}

func processSeedsAndMapsP2(seedData []uint64, maps map[DataSetKey][]IdMap) uint64 {
	lowestLocation := uint64(18446744073709551615) //uint64 max value

	allDataSetKeys := []DataSetKey{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight,
		lightToTemperature, temperatureToHumidity, humidityToLocation}

	tikTok := false
	seedBase := uint64(0)
	seedSpread := uint64(0)
	count := 0
	pairs := len(seedData) / 2
	for _, seedDataElement := range seedData {
		if !tikTok {
			seedBase = seedDataElement
			tikTok = true
			continue
		} else {
			seedSpread = seedDataElement
			tikTok = false
		}

		count++
		fmt.Printf("%d/%d: base %v spread %v\n", count, pairs, seedBase, seedSpread)

		for seed := seedBase; seed < seedBase+seedSpread; seed++ {
			// fmt.Printf("seed %v\n", seed)
			currentId := seed
			for _, key := range allDataSetKeys {
				// fmt.Printf("\t%v\n", key)
				var foundIdMap IdMap
				found := false
				// is id within any of this data set's maps
				for _, idMap := range maps[key] {
					if idMap.source <= currentId && idMap.source+idMap.spread > currentId {
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
	}

	return lowestLocation
}

func parseInputP2(scanner *bufio.Scanner) ([]uint64, map[DataSetKey][]IdMap) {
	dataMap := make(map[DataSetKey][]IdMap)
	seeds := make([]uint64, 0)

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
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, uint64(num))
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
		nums := make([]uint64, 0)
		for _, numStr := range sSplit {
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, uint64(num))
		}

		var idMap = IdMap{nums[0], nums[1], nums[2]}

		if dataMap[currentDataSet] == nil {
			dataMap[currentDataSet] = make([]IdMap, 0)
		}
		dataMap[currentDataSet] = append(dataMap[currentDataSet], idMap)
	}

	return seeds, dataMap
}
