package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
- make a new sequence from the difference at each step of the history

- if that sequence is not all zeros,
repeat this process, using the sequence you just generated as the input sequence
if that sequence is not all zeros -> repeat thisprocess using the sequence just
generated as the input sequence

- once all of the values in your latest sequence are zeros, you can extrapolate what the next
value of the original history should be

	eg
	1   3   6  10  15  21 | 28
	2   3   4   5   6	| 7
	1   1   1   1 | 1
	0   0   0	| 0
	lower tier increments upper tier (0 added for free in bottom tier)
*/
func main() {
	fmt.Println("day 9 p 1")
	file_name := "example_input.txt" // expecting 114 (18, 28, 68)
	// file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputNums := parse(scanner)

	// fmt.Printf("%+v\n", inputNums)

	sum := process(inputNums)
	fmt.Printf("sum: %d\n", sum)
}

func parse(scanner *bufio.Scanner) [][]int {
	nums := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")

		lineNums := make([]int, 0)
		for _, s := range splitLine {
			val, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			lineNums = append(lineNums, val)
		}
		nums = append(nums, lineNums)
	}

	return nums
}

func process(nums [][]int) int {
	
	return 0
}
