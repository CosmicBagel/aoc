package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card int

type Hand struct {
	cards    [5]Card
	bid      int
	handType HandType //for sorting later
}

const (
	C2 Card = iota
	C3
	C4
	C5
	C6
	C7
	C8
	C9
	C10
	Jack
	Queen
	King
	Ace
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse //3 of a kind + pair
	FourOfAKind
	FiveOfAKind
)

type Hands []Hand

func (h Hands) Len() int      { return len(h) }
func (h Hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Hands) Less(i, j int) bool {
	a := h[i]
	b := h[j]
	if a.handType == b.handType {
		for ci := 0; ci < 5; ci++ {
			// other has better card
			if a.cards[ci] < b.cards[ci] {
				return true
			}

			// other has lower value card
			if a.cards[ci] > b.cards[ci] {
				return false
			}
		}
		// hands are equal
		return false
	}

	// different hands types, can just use that
	return a.handType < b.handType
}

func Day7P1() {
	fmt.Println("day 7 p 1")
	// file_name := "example_input.txt"
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := make(Hands, 0)
	// printHands := func() {
	// 	for i, h := range hands {
	// 		fmt.Printf("hand %d: %+v\n", i, h)
	// 	}
	// }

	for scanner.Scan() {
		hands = append(hands, parseP1(scanner.Text()))
	}
	// printHands()
	sort.Sort(hands)

	sum := 0
	for i, h := range hands {
		value := h.bid * (i + 1)
		// fmt.Printf("hand %d: %+v\n", i, h)
		// fmt.Printf("\tvalue %d\n", value)
		sum += value

	}
	fmt.Printf("total value: %d\n", sum)
}

func parseP1(s string) Hand {
	hand := Hand{}

	split := strings.Split(s, " ")

	// cards
	for i, r := range split[0] {
		if r >= 50 && r <= 57 {
			hand.cards[i] = Card(r - 50)
			continue
		}

		switch r {
		case 'T':
			hand.cards[i] = C10
		case 'J':
			hand.cards[i] = Jack
		case 'Q':
			hand.cards[i] = Queen
		case 'K':
			hand.cards[i] = King
		case 'A':
			hand.cards[i] = Ace
		}
	}

	//hand type
	hand.handType = calcHandType(hand.cards)
	// fmt.Printf("\thand type result %d\n", hand.handType)

	// bid
	num, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}
	hand.bid = num

	return hand
}

func calcHandType(cards [5]Card) HandType {
	// fmt.Printf("%+v\n", cards)
	pairs := 0
	isThreeOfAKind := false

	counts := map[Card]int{}
	for _, c := range cards {
		if _, ok := counts[c]; !ok {
			counts[c] = 0
		}
		counts[c] += 1
	}
	// fmt.Printf("\tcounts: %+v\n", counts)

	for k := range counts {
		n := counts[k]
		switch n {
		case 2:
			pairs += 1
		case 3:
			isThreeOfAKind = true
		case 4:
			return FourOfAKind
		case 5:
			return FiveOfAKind
		}
	}
	// fmt.Printf("\tparis: %d\n", pairs)
	// fmt.Printf("\tisThreeOfAKind: %+v\n", isThreeOfAKind)

	if isThreeOfAKind {
		if pairs > 0 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	}

	if pairs == 2 {
		return TwoPair
	}

	if pairs == 1 {
		return OnePair
	}

	return HighCard
}
