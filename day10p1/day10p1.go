package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
build graph

	keep last line of nodes so we can connect any if needed (simple array will do)
	only next and prev pointers on each node
	node type can be pipe or empty
	also store X,Y point node is located in (may help with debug)
	also store original node character (for debug)
	return starting node

	- note connections can only be made if both points have mutual open side
		eg -| no connection
		   -7 connection

trace from start forwards and backwards, where the two traveling pointers meet
will be the furthest distance (only following prev pointers, one following next pointers)
*/

type Point struct {
	X int
	Y int
}

type NodeType int

const (
	Empty NodeType = iota
	Vertical
	Horizontal
	NEBend
	NWBend
	SWBend
	SEBend
	Start
)

type Node struct {
	north *Node
	east  *Node
	south *Node
	west  *Node

	nodeType       NodeType
	location       Point
	originalMarker rune
}

func main() {
	fmt.Println("day 10 p 1")
	file_name := "example_inputA.txt" // expecting 4 at point 3,3 (0,0 is top left)
	// file_name := "example_inputB.txt" // expecting 4 at point 3,3  (same as A, but with excess pipe)
	// file_name := "example_inputC.txt" // expecting 8 at point 4,2
	// file_name := "example_inputD.txt" // expecting 8 at point 4,2  (same as C, but with excess pipe)
	// file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scanner.Scan()
	// fmt.Println(scanner.Text())
	startingNode := parse(scanner)
	result := findFarthestDistance(startingNode)

	fmt.Printf("fartherst distance: %d", result)
}

func parse(scanner *bufio.Scanner) Node {
	startingNode := Node{nodeType: Empty}

	scanner.Scan()
	currentLine := scanner.Text()
	gridWidth := len(currentLine)

	aboveNodes := make([]Node, 0, gridWidth)
	currentNodes := make([]Node, 0, gridWidth)

	for i := 0; i < gridWidth; i++ {
		n := Node{
			location:       Point{i, -1},
			originalMarker: '.',
			nodeType:       Empty,
		}
		aboveNodes = append(aboveNodes, n)
	}

	y := 0
	for {
		for x, c := range currentLine {

			loc := Point{x, y}
			aboveNode := aboveNodes[x]
			leftNode := Node{nodeType: Empty}
			if x > 0 {
				leftNode = currentNodes[x-1]
			}

			n := Node{
				location:       loc,
				originalMarker: c,
				nodeType:       Empty,
			}

			ant := aboveNode.nodeType
			aboveHasSouth := ant == Vertical || ant == SEBend || ant == SWBend || ant == Start
			lnt := leftNode.nodeType
			leftHasEast := lnt == Horizontal || lnt == NEBend || lnt == SEBend || lnt == Start

			switch c {
			case '|':
				n.nodeType = Vertical
				if aboveHasSouth {
					n.north = &aboveNode
					aboveNode.south = &n
				}
			case '-':
				n.nodeType = Horizontal
				if leftHasEast {
					n.west = &n
					leftNode.east = &n
				}
			case 'L':
				n.nodeType = NEBend
				if aboveHasSouth {
					n.north = &aboveNode
					aboveNode.south = &n
				}
			case 'J':
				n.nodeType = NWBend
				if leftHasEast {
					n.west = &n
					leftNode.east = &n
				}
				if aboveHasSouth {
					n.north = &aboveNode
					aboveNode.south = &n
				}
			case '7':
				n.nodeType = SWBend
				if leftHasEast {
					n.west = &n
					leftNode.east = &n
				}
			case 'F':
				n.nodeType = SEBend
			case 'S':
				n.nodeType = Start
				startingNode = n
				if leftHasEast {
					n.west = &n
					leftNode.east = &n
				}
				if aboveHasSouth {
					n.north = &aboveNode
					aboveNode.south = &n
				}
			}

			currentNodes = append(currentNodes, n)
		}

		// clear and swap lines
		clear(aboveNodes)
		temp := aboveNodes
		aboveNodes = currentNodes
		currentNodes = temp

		y += 1
		if !scanner.Scan() {
			break
		}
	}

	if startingNode.nodeType == Empty {
		log.Fatal("did not find starting node")
	}
	return startingNode
}

func findFarthestDistance(startingNode Node) int {
	return 0
}
