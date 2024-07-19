package adventofcode2023golang

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day2P1() {
	red_max := 12
	green_max := 13
	blue_max := 14
	// file_name := "example_input.txt"
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	regex_id, _ := regexp.Compile("(\\d+)$")
	isGameValid := func(s string) (int, bool) {
		valid := true
		split_a := strings.Split(s, ": ")

		id_str := regex_id.FindString(split_a[0])
		id, _ := strconv.Atoi(id_str)

		split_b := strings.Split(split_a[1], "; ")
	outter:
		for _, s_a := range split_b {
			split_c := strings.Split(s_a, ", ")

			for _, s_b := range split_c {
				split_d := strings.Split(s_b, " ")
				count, _ := strconv.Atoi(split_d[0])
				color := split_d[1]

				if color == "red" && count > red_max {
					valid = false
					break outter
				}

				if color == "green" && count > green_max {
					valid = false
					break outter
				}

				if color == "blue" && count > blue_max {
					valid = false
					break outter
				}
			}
		}

		return id, valid
	}

	sum := 0
	for scanner.Scan() {
		id, ok := isGameValid(scanner.Text())
		if ok {
			sum += id
		}

	}

	println(sum)
}

func Day2P2() {
	// file_name := "example_input.txt"
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// regex_id, _ := regexp.Compile("(\\d+)$")
	minCubes := func(s string) (int, int, int) {
		split_a := strings.Split(s, ": ")

		// id_str := regex_id.FindString(split_a[0])
		// id, _ := strconv.Atoi(id_str)

		split_b := strings.Split(split_a[1], "; ")

		green_min := 0
		red_min := 0
		blue_min := 0

		for _, s_a := range split_b {
			split_c := strings.Split(s_a, ", ")

			for _, s_b := range split_c {
				split_d := strings.Split(s_b, " ")
				count, _ := strconv.Atoi(split_d[0])
				color := split_d[1]

				if color == "red" && count > red_min {
					red_min = count
					continue
				}

				if color == "green" && count > green_min {
					green_min = count
					continue
				}

				if color == "blue" && count > blue_min {
					blue_min = count
					continue
				}
			}
		}

		return green_min, red_min, blue_min
	}

	sum := 0
	for scanner.Scan() {
		green_min, red_min, blue_min := minCubes(scanner.Text())
		sum += green_min * red_min * blue_min

	}

	println(sum)
}
