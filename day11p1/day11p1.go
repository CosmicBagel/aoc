package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("day 11 p 1")
	file_name := "example_input.txt" // expecting 374
	// file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//blank canvas, then add just to path to it

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}
