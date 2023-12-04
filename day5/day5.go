package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day5P1() {
	fmt.Println("day 5 p 1")
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
