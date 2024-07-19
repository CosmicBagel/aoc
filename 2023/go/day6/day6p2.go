package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day6P2() {
	fmt.Println("day 6 p 2")
	file_name := "example_input.txt"
	// file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Printf("%+v\n", scanner.Text())
	}
}
