package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type DistancesMm []int
type TimesMs []int

func Day7() {
	fmt.Println("day 7 p 1")
	file_name := "example_input.txt"
	// file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// distances, times := parse(scanner)
	// result := process(distances, times)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
