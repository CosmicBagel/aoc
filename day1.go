package adventofcode2023golang

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "strings"

	"github.com/dghubble/trie"
)

func Day1() {
	// 	example_input :=
	// 		`1abc2
	// pqr3stu8vwx
	// a1b2c3d4e5f
	// treb7uchet`

	file, err := os.Open("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		a, b := processLineD1(scanner.Text())
		sum += (a * 10) + b
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	println(sum)
}

func processLineD1(s string) (a int, b int) {
	bytes := []byte(s)
	first_num := 0
	for _, b := range bytes {
		if b <= 57 && b >= 48 {
			first_num = int(b - 48)
			break
		}
	}

	second_num := 0
	for i := len(bytes) - 1; i >= 0; i-- {
		b := bytes[i]
		if b <= 57 && b >= 48 {
			second_num = int(b - 48)
			break
		}
	}

	return first_num, second_num
}

func Day1P2() {
	tireForward := trie.NewRuneTrie()
	tireForward.Put("one", 1)
	tireForward.Put("1", 1)
	tireForward.Put("two", 2)
	tireForward.Put("2", 2)
	tireForward.Put("three", 3)
	tireForward.Put("3", 3)
	tireForward.Put("four", 4)
	tireForward.Put("4", 4)
	tireForward.Put("five", 5)
	tireForward.Put("5", 5)
	tireForward.Put("six", 6)
	tireForward.Put("6", 6)
	tireForward.Put("seven", 7)
	tireForward.Put("7", 7)
	tireForward.Put("eight", 8)
	tireForward.Put("8", 8)
	tireForward.Put("nine", 9)
	tireForward.Put("9", 9)

	tireBack := trie.NewRuneTrie()
	tireBack.Put("eno", 1)
	tireBack.Put("1", 1)
	tireBack.Put("owt", 2)
	tireBack.Put("2", 2)
	tireBack.Put("eerht", 3)
	tireBack.Put("3", 3)
	tireBack.Put("ruof", 4)
	tireBack.Put("4", 4)
	tireBack.Put("evif", 5)
	tireBack.Put("5", 5)
	tireBack.Put("xis", 6)
	tireBack.Put("6", 6)
	tireBack.Put("neves", 7)
	tireBack.Put("7", 7)
	tireBack.Put("thgie", 8)
	tireBack.Put("8", 8)
	tireBack.Put("enin", 9)
	tireBack.Put("9", 9)

	// 	exampleInput :=
	// 		`two1nine
	// eightwothree
	// abcone2threexyz
	// xtwone3four
	// 4nineeightseven2
	// zoneight234
	// 7pqrstsixteen`

	file, err := os.Open("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// scanner := bufio.NewScanner(strings.NewReader(exampleInput))
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		a := search(s, tireForward)
		b := search(reverse(s), tireBack)

		sum += (a * 10) + b
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	println(sum)
}

func search(s string, t *trie.RuneTrie) int {
	num := 0
	for i := 0; i < len(s); i++ {
		err := t.WalkPath(s[i:], func(key string, value interface{}) error {
			n, ok := value.(int)
			if !ok {
				fmt.Println("v not an int")
			}
			num = n
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}

		if num != 0 {
			break
		}
	}
	return num
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
