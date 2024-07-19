package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type DistancesMm []int
type TimesMs []int

func Day6() {
	fmt.Println("day 6 p 1")
	// file_name := "example_input.txt"
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	distances, times := parse(scanner)
	result := process(distances, times)
	fmt.Println(result)
}

func parse(scanner *bufio.Scanner) (DistancesMm, TimesMs) {
	scanner.Scan()
	line1 := scanner.Text()
	scanner.Scan()
	line2 := scanner.Text()

	getNums := func(line string) []int {
		lineSplit := strings.Split(line, " ")

		var list []int = make([]int, 0)
		var sb strings.Builder
		for _, s := range lineSplit[1:] {
			trimmed := strings.Trim(s, " ")
			if len(trimmed) == 0 {
				continue
			}
			sb.WriteString(trimmed)
		}

		num, err := strconv.Atoi(sb.String())
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, num)
		return list
	}

	var times TimesMs = getNums(line1)
	var distances DistancesMm = getNums(line2)

	return distances, times
}

func process(distances DistancesMm, times TimesMs) int {
	/*
		each ms button is held -> speed increased by 1mm/ms

	*/

	prodOfCounts := 1

	fmt.Printf("times %+v\n", times)
	fmt.Printf("distances %+v\n", distances)

	limit := len(times)
	if limit != len(distances) {
		log.Fatal("times and distances do not have the same element count")
	}

	for i := 0; i < limit; i++ {
		t := times[i]
		d := distances[i]

		minBT := t
		maxBT := 0

		for tChecking := 0; tChecking <= t; tChecking++ {
			//tChecking is time held button down
			//equates to mm/s
			dAchieved := tChecking * (t - tChecking)
			if dAchieved > d {
				if tChecking > maxBT {
					maxBT = tChecking
				}
				if tChecking < minBT {
					minBT = tChecking
				}
			}
		}

		prodOfCounts *= maxBT - minBT + 1
	}

	return prodOfCounts
}
