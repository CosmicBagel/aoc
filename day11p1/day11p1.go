package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type MapData struct {
	width  int
	height int

	galaxies []Point
}

func (d MapData) draw() {
	grid := make([][]byte, d.height)
	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			grid[y] = append(grid[y], '.')
		}
	}

	for _, gp := range d.galaxies {
		grid[gp.y][gp.x] = '#'
	}

	fmt.Printf("displaying galaxy map\n")
	for y := 0; y < d.height; y++ {
		fmt.Printf("%s\n", string(grid[y]))
	}
}

type Point struct {
	x int
	y int
}

type Pair struct {
	a Point
	b Point
}

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
	mapData := parse(scanner)
	mapData.draw()

	// sum of shortest path between each pair of galaxies
}

func parse(scanner *bufio.Scanner) MapData {
	// expands grid as we parse

	galaxies := make([]Point, 0)

	scanner.Scan()
	currentLine := scanner.Text()
	originalWidth := len(currentLine)
	emptyColumns := make([]bool, originalWidth)
	for i := range emptyColumns {
		emptyColumns[i] = true
	}

	verOffset := 0
	verticalExpansionOffsets := make([]int, 0)

	y := 0
	for {
		rowEmpty := true
		fmt.Printf("%s\n", scanner.Bytes())
		for x, b := range scanner.Bytes() {
			if b == '#' {
				galaxies = append(galaxies, Point{x, y})
				emptyColumns[x] = false
				rowEmpty = false
			}
		}
		if rowEmpty {
			verOffset += 1
		}
		verticalExpansionOffsets = append(verticalExpansionOffsets, verOffset)
		y += 1
		if !scanner.Scan() {
			break
		}
	}
	originalHeight := y

	horOffset := 0
	horizontalExpansionOffsets := make([]int, 0)
	for _, c := range emptyColumns {
		if c {
			horOffset += 1
		}
		horizontalExpansionOffsets = append(horizontalExpansionOffsets, horOffset)
	}

	// update galaxy points with universe expansion offsets
	for i, gp := range galaxies {
		p := Point{gp.x + horizontalExpansionOffsets[gp.x],
			gp.y + verticalExpansionOffsets[gp.y]}
		galaxies[i] = p
	}

	dat := MapData{}
	dat.galaxies = galaxies
	dat.width = originalWidth + horizontalExpansionOffsets[originalWidth-1]
	dat.height = originalHeight + verticalExpansionOffsets[originalHeight-1]

	return dat
}

func enumeratePairs(mapData MapData) []Pair {
	//    n!
	// ------- = expected count
	// k!(n-k)!
	// where n is the number of objects (galaxies) and k is how many are chosen (1 pair = 2)

	galaxyCount := len(mapData.galaxies)
	expectedCount := factorial(galaxyCount) / (2 * factorial(galaxyCount-2))

	fmt.Printf("expectedCount %d\n", expectedCount)
	pairs := make([]Pair, 0, expectedCount)

	for i, ga := range mapData.galaxies[:galaxyCount-1] {
		for _, gb := range mapData.galaxies[i+1:] {
			pairs = append(pairs, Pair{ga, gb})
		}
	}

	// 01 02 03 04 05 06 07 08
	// 12 13 14 15 16 17 18
	// 34 35 36 37 38
	// 45 46 47 48
	// 56 57 58
	// 67 68
	// 78
	//36 possible combinations

	fmt.Printf("galaxy pairs: %d\n", len(pairs))

	return pairs
}

func factorial(n int) int {
	f := n
	for f > 1 {
		f -= 1
		n *= f
	}

	return n
}
