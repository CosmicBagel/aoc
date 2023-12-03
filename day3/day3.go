package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var num_re *regexp.Regexp

func Day3P1() {

	fmt.Println("day 3 p 1")

	/*
		hold up to three rows simultaneously
		read two rows to start
		scan current row for symbol (anything not a number and not .)
		check for numbers adjacent to symbol on current row
		check for numbers on above row (if present) and below row (if present) for vertically
			and diagonally adjacent numbers

		use regex to find the indices of start and end of each number and each symbol
		beware of duplicate number use eg
		.123.  here 123 would be caught in diagonal top left, diagonal top right,
		..*..  and vertical up checks
		.456.  here 456 would be caught in diagonal bottom left, diagonal bottom right
			   and vertical down checks

		alternatively find numbers, then check for symbol vertically or
			diagonally (touching corner)
		probably simpler to process, don't need any kind of bookkeeping for used numbers
	*/

	var err error
	num_re, err = regexp.Compile("(\\d+)")
	if err != nil {
		log.Fatal(err)
	}

	// file_name := "example_input.txt"
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	above_line := ""
	scanner.Scan()
	current_line := scanner.Text()
	scanner.Scan()
	below_line := scanner.Text()
	//process
	sum += p1ProcessLine(above_line, current_line, below_line)

	for scanner.Scan() {
		above_line = current_line
		current_line = below_line
		below_line = scanner.Text()

		//process
		sum += p1ProcessLine(above_line, current_line, below_line)
	}

	above_line = current_line
	current_line = below_line
	below_line = ""
	//process
	sum += p1ProcessLine(above_line, current_line, below_line)

	println(sum)
}

func p1ProcessLine(above_line string, current_line string, below_line string) int {
	sum := 0
	isAboveEmpty := len(above_line) == 0
	isBelowEmpty := len(below_line) == 0

	/*
		..abbbbbc..
		...12345...
		..deeeeef..
		b is vertical up region check (n checks where n is digits in number)
		e is vertical down region check (n checks ^^^)
		a is diagonal up left check
		c is diagonal up right check
		d is diagonal down left check
		f is diagonal down right check
	*/

	num_indices := num_re.FindAllStringIndex(current_line, -1)

	for _, inds := range num_indices {
		// println("A")
		left := inds[0]
		right := inds[1]

		partCount := 0

		// println("B")
		isSymbol := func(c byte) bool {
			// not a number and not period
			b := (c < 48 || c > 57) && c != '.'
			// if b {
			// 	fmt.Printf("%s ", string(c))
			// }
			return b
		}

		// vertical up checks
		if !isAboveEmpty {
			for i := left; i < right; i++ {
				c := above_line[i]
				if isSymbol(c) {
					partCount += 1
				}
			}
		}

		// println("C")

		// vertical down checks
		if !isBelowEmpty {
			for i := left; i < right; i++ {
				c := below_line[i]
				if isSymbol(c) {
					partCount += 1
				}
			}
		}

		// horizontal left right checks
		if left > 0 {
			c := current_line[left-1]
			if isSymbol(c) {
				partCount += 1
			}
		}
		if right < len(current_line) {
			c := current_line[right]
			if isSymbol(c) {
				partCount += 1
			}
		}

		// println("D")
		// above diagonal checks
		if left > 0 && !isAboveEmpty {
			c := above_line[left-1]
			if isSymbol(c) {
				partCount += 1
			}
		}
		// println("E")

		if right < len(current_line) && !isAboveEmpty {
			c := above_line[right]
			if isSymbol(c) {
				partCount += 1
			}
		}

		// println("F")
		// below diagonal checks
		if left > 0 && !isBelowEmpty {
			c := below_line[left-1]
			if isSymbol(c) {
				partCount += 1
			}
		}

		// println("G")
		if right < len(current_line) && !isBelowEmpty {
			c := below_line[right]
			if isSymbol(c) {
				partCount += 1
			}
		}

		// println("H")
		if partCount > 0 {
			num, err := strconv.Atoi(current_line[left:right])
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Printf("%d %d\n", num, partCount)
			sum += num * partCount
		}

	}

	return sum
}

func Day3P2() {
	println("day 3 p 2")
}
