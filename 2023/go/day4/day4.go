package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day4P1() {
	fmt.Println("day 4 p 1")
	// file_name := "example_input.txt"
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		sum += processLineP1(s)
	}

	fmt.Println(sum)
}

func processLineP1(s string) int {
	points := 0

	pointsUp := func () {
		if points == 0 {
			points = 1
			return
		}

		points *= 2
	}

	s_name_removed := strings.Split(s, ": ")[1]
	s_winning_and_scratched := strings.Split(s_name_removed, " | ")
	s_winning := s_winning_and_scratched[0]
	s_scratched := s_winning_and_scratched[1]
	
	s_winning_numbers := strings.Split(s_winning, " ")
	s_scratched_numbers := strings.Split(s_scratched, " ")

	winning_numbers := make(map[int]bool, 0)

	// fmt.Println(s_winning_numbers)
	// fmt.Println(s_scratched_numbers)

	for _, s_num := range s_winning_numbers {
		if len(s_num) == 0 {
			continue
		}
		s_num_trimmed := strings.Trim(s_num, " ")
		num, err := strconv.Atoi(s_num_trimmed)
		if err != nil {
			log.Fatalf("winning s_num_trimmed: %s %v",s_num_trimmed, err)
		}

		winning_numbers[num] = true
	}
	
	for _, s_num := range s_scratched_numbers {
		if len(s_num) == 0 {
			continue
		}
		s_num_trimmed := strings.Trim(s_num, " ")
		num, err := strconv.Atoi(s_num)
		if err != nil {
			log.Fatalf("scratched s_num_trimmed: %s %v",s_num_trimmed, err)
		}

		if winning_numbers[num] {
			pointsUp()
		}
	}

	return points
}

func Day4P2() {
	/*
	cards_processed := 0

	card_copy_count := make(map[int]int)
	always process card as we go through the list

	*/
	println("day 4 p 2")
	// file_name := "example_input.txt"
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	line_num := 0
	scanner := bufio.NewScanner(file)

	cards_processed := 0
	card_copies_won := make(map[int]int)
	for scanner.Scan() {
		s := scanner.Text()
		line_num += 1

		copies_of_card := card_copies_won[line_num]
		delete(card_copies_won, line_num)

		winning_count := processLineP2(s)
		// fmt.Printf("%d: winning_count: %d\n", line_num, winning_count)
		for i := line_num + 1; i <= line_num + winning_count; i++ {
			card_copies_won[i] += copies_of_card + 1
			// fmt.Printf("\t%d adding %d copies\n", i, copies_of_card)
		}

		cards_processed += copies_of_card + 1
	}

	fmt.Println(cards_processed)
}

// returns card copy winning count
func processLineP2(s string) int {
	winning_count := 0

	pointsUp := func () {
		winning_count += 1
	}

	s_name_removed := strings.Split(s, ": ")[1]
	s_winning_and_scratched := strings.Split(s_name_removed, " | ")
	s_winning := s_winning_and_scratched[0]
	s_scratched := s_winning_and_scratched[1]
	
	s_winning_numbers := strings.Split(s_winning, " ")
	s_scratched_numbers := strings.Split(s_scratched, " ")

	winning_numbers := make(map[int]bool, 0)

	// fmt.Println(s_winning_numbers)
	// fmt.Println(s_scratched_numbers)

	for _, s_num := range s_winning_numbers {
		if len(s_num) == 0 {
			continue
		}
		s_num_trimmed := strings.Trim(s_num, " ")
		num, err := strconv.Atoi(s_num_trimmed)
		if err != nil {
			log.Fatalf("winning s_num_trimmed: %s %v",s_num_trimmed, err)
		}

		winning_numbers[num] = true
	}
	
	for _, s_num := range s_scratched_numbers {
		if len(s_num) == 0 {
			continue
		}
		s_num_trimmed := strings.Trim(s_num, " ")
		num, err := strconv.Atoi(s_num)
		if err != nil {
			log.Fatalf("scratched s_num_trimmed: %s %v",s_num_trimmed, err)
		}

		if winning_numbers[num] {
			pointsUp()
		}
	}

	return winning_count
}
