package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
)

type MapData struct {
	width  int64
	height int64

	galaxies []Point
}

func (d MapData) draw() {
	grid := make([][]byte, d.height)
	for y := int64(0); y < d.height; y++ {
		for x := int64(0); x < d.width; x++ {
			grid[y] = append(grid[y], '.')
		}
	}

	for _, gp := range d.galaxies {
		grid[gp.y][gp.x] = '#'
	}

	fmt.Printf("displaying  galaxy map\n")
	for y := int64(0); y < d.height; y++ {
		fmt.Printf("%s\n", string(grid[y]))
	}
}

type Point struct {
	x int64
	y int64
}

type Pair struct {
	a Point
	b Point
}

func main() {
	fmt.Println("day 11 p 2")
	// file_name := "example_input.txt" // expecting 1030
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mapData := parse(scanner)
	// mapData.draw()
	pairs := enumeratePairs(mapData)

	// sum of shortest path between each pair of galaxies
	sum := findShortestPathForAll(pairs)
	fmt.Printf("sum of shortest paths: %d\n", sum)
}

func parse(scanner *bufio.Scanner) MapData {
	// expands grid as we parse

	galaxies := make([]Point, 0)

	expansionRate := int64(999999)

	scanner.Scan()
	currentLine := scanner.Text()
	originalWidth := int64(len(currentLine))
	emptyColumns := make([]bool, originalWidth)
	for i := range emptyColumns {
		emptyColumns[i] = true
	}

	verOffset := int64(0)
	verticalExpansionOffsets := make([]int64, 0)

	y := 0
	for {
		rowEmpty := true
		// fmt.Printf("%s\n", scanner.Bytes())
		for x, b := range scanner.Bytes() {
			if b == '#' {
				galaxies = append(galaxies, Point{int64(x), int64(y)})
				emptyColumns[x] = false
				rowEmpty = false
			}
		}
		if rowEmpty {
			verOffset += expansionRate
		}
		verticalExpansionOffsets = append(verticalExpansionOffsets, verOffset)
		y += 1
		if !scanner.Scan() {
			break
		}
	}
	originalHeight := int64(y)

	horOffset := int64(0)
	horizontalExpansionOffsets := make([]int64, 0)
	for _, c := range emptyColumns {
		if c {
			horOffset += expansionRate
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

	galaxyCount := int64(len(mapData.galaxies))
	// fmt.Printf("galaxyCount %d\n", galaxyCount)

	// expectedCount := calcExpectedCount(galaxyCount)
	// fmt.Printf("expectedCount %d\n", expectedCount)
	pairs := make([]Pair, 0)

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

func calcExpectedCount(n int64) int64 {
	two := big.NewInt(2)
	numerator := factorial(n)
	denominator := factorial(n - 2)
	denominator = denominator.Mul(denominator, two)
	expectedCount := numerator.Div(numerator, denominator)

	return expectedCount.Int64()
}

func factorial(in int64) *big.Int {
	n := big.NewInt(in)
	f := big.NewInt(in)
	one := big.NewInt(1)
	negOne := big.NewInt(-1)
	for f.Cmp(one) == 1 {
		f = f.Sub(f, negOne)
		n = n.Mul(f, n)
	}

	return n
}

func findShortestPathForAll(pairs []Pair) int64 {
	sum := int64(0)
	for _, pair := range pairs {
		diff := Point{
			int64(math.Abs(float64(pair.b.x - pair.a.x))),
			int64(math.Abs(float64(pair.b.y - pair.a.y))),
		}

		sum += diff.x + diff.y
	}

	return sum
}
