package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
2d grid
info for each element:
- isPath
- flow
- originalCharacter
-

replace all none path elements with .

flood fill from edge to find outside edge of loop

follow loop, note direction flow, marking open slots as
	inside or outside the loop (if its not part of the loop)

add inside marks to list as you go

flood search from each inside mark (not infinitely obviously)

report count of found inside tiles


------

create blank grid with just the loop on it (using existing found path)

find left most part of the loop
	know left of this is out of the loop, right is in the loop

follow loop with this (adjusting perspective with bends / turns)
	when there's an empty tile, flood fill search for other inner tiles

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

type Grid struct {
	height int
	width  int

	elements [][]byte
}

func makeGrid(width, height int) *Grid {
	g := Grid{
		width:    width,
		height:   height,
		elements: make([][]byte, height),
	}

	for y := 0; y < height; y++ {
		row := make([]byte, width)
		g.elements[y] = row
		for x := 0; x < width; x++ {
			row[x] = '.'
		}
	}
	return &g
}

func (g *Grid) getPoint(p Point) byte {
	return g.elements[p.Y][p.X]
}

func (g *Grid) setPoint(p Point, b byte) {
	g.elements[p.Y][p.X] = b
}

func (g *Grid) draw() {
	fmt.Printf("printing %d by %d grid\n", g.width, g.height)
	for y := 0; y < g.height; y++ {
		fmt.Printf("%s\n", string(g.elements[y]))
	}
}

func (g *Grid) boundsCheck(p Point) bool {
	if p.X < 0 || p.X >= g.width ||
		p.Y < 0 || p.Y >= g.height {
		return false
	}
	return true
}

func main() {
	fmt.Println("day 10 p 2")
	// file_name := "example_inputA.txt" // expecting 4
	// file_name := "example_inputB.txt" // expecting 4
	// file_name := "example_inputC.txt" // expecting 8
	// file_name := "example_inputD.txt" // expecting 10
	// file_name := "example_inputE.txt" // expecting 1
	// file_name := "example_inputF.txt" // expecting 1
	// file_name := "example_inputG.txt" // expecting 1
	// file_name := "example_inputH.txt" // expecting 1
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	startingNode, gridWidth, gridHeight := parse(scanner)
	//blank canvas, then add just to path to it
	grid := makeGrid(gridWidth, gridHeight)
	writePathToGrid(startingNode, grid)

	// grid.draw()
	result := findInnerTileCount(startingNode, grid)
	// grid.draw()
	// result := findFarthestDistance(startingNode)

	fmt.Printf("enclosed tiles: %d\n", result)
}

func parse(scanner *bufio.Scanner) (*Node, int, int) {
	var startingNode *Node = nil

	scanner.Scan()
	currentLine := scanner.Text()
	gridWidth := len(currentLine)

	aboveNodes := make([]*Node, 0, gridWidth)
	currentNodes := make([]*Node, 0, gridWidth)

	for i := 0; i < gridWidth; i++ {
		n := Node{
			location:       Point{i, -1},
			originalMarker: '.',
			nodeType:       Empty,
		}
		aboveNodes = append(aboveNodes, &n)
	}

	y := 0
	for {
		// fmt.Println(currentLine)
		for x, c := range currentLine {
			loc := Point{x, y}
			var aboveNode = aboveNodes[x]
			var leftNode *Node = nil
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

			leftHasEast := false
			if leftNode != nil {
				lnt := leftNode.nodeType
				leftHasEast = lnt == Horizontal || lnt == NEBend || lnt == SEBend || lnt == Start
			}

			switch c {
			case '|':
				n.nodeType = Vertical
				if aboveHasSouth {
					n.north = aboveNode
					aboveNode.south = &n
				}
			case '-':
				n.nodeType = Horizontal
				if leftHasEast {
					n.west = leftNode
					leftNode.east = &n
				}
			case 'L':
				n.nodeType = NEBend
				if aboveHasSouth {
					n.north = aboveNode
					aboveNode.south = &n
				}
			case 'J':
				n.nodeType = NWBend
				if leftHasEast {
					n.west = leftNode
					leftNode.east = &n
				}
				if aboveHasSouth {
					n.north = aboveNode
					aboveNode.south = &n
				}
			case '7':
				n.nodeType = SWBend
				if leftHasEast {
					n.west = leftNode
					leftNode.east = &n
				}
			case 'F':
				n.nodeType = SEBend
				// must be connected to by an east node or south node
			case 'S':
				n.nodeType = Start
				startingNode = &n
				if leftHasEast {
					n.west = leftNode
					leftNode.east = &n
				}
				if aboveHasSouth {
					n.north = aboveNode
					aboveNode.south = &n
				}
			}

			currentNodes = append(currentNodes, &n)
		}

		// for _, n := range currentNodes {
		// 	fmt.Printf("%+v\n", n)
		// }

		// clear and swap lines
		temp := aboveNodes[:0]
		aboveNodes = currentNodes
		currentNodes = temp

		y += 1
		if !scanner.Scan() {
			if scanner.Err() != nil {
				log.Fatal(scanner.Err())
			}
			break
		}
		currentLine = scanner.Text()
	}

	if startingNode == nil {
		log.Fatal("did not find starting node")
	}

	return startingNode, gridWidth, y
}

type Dir int

const (
	North Dir = iota
	East
	South
	West
)

func dirToName(d Dir) string {
	name := "West"
	switch d {
	case North:
		name = "North"
	case East:
		name = "East"
	case South:
		name = "South"
	}

	return name
}
func (d Dir) oppositeDir() Dir {
	return Dir((d + 2) % 4)
}

func writePathToGrid(startingNode *Node, grid *Grid) {
	traveler := makeTravelFunc(startingNode)

	grid.setPoint(startingNode.location, byte(startingNode.originalMarker))

	for {
		n, _ := traveler()

		if n.originalMarker == 'S' {
			break
		}

		grid.setPoint(n.location, byte(n.originalMarker))

	}

}

func makeTravelFunc(start *Node) func() (*Node, Dir) {
	lastDirTraveled := North
	currentNode := start

	return func() (*Node, Dir) {
		directions := []*Node{
			currentNode.north,
			currentNode.east,
			currentNode.south,
			currentNode.west,
		}

		entryPoint := lastDirTraveled.oppositeDir()
		for i, n := range directions {
			if Dir(i) != entryPoint && n != nil {
				lastDirTraveled = Dir(i)
				currentNode = n
				return currentNode, lastDirTraveled
			}
		}

		log.Fatal("Failed to find valid direction (incomplete circle?)")
		return nil, North
	}
}

func findInnerTileCount(startingNode *Node, grid *Grid) int {

	// step 1 find left most location (or one of them)
	//	left side is known outside, right is known inside
	// step 2, trace from point, noting orientation as you go

	traveler := makeTravelFunc(startingNode)
	secondNode, secondDir := traveler()
	lowestX := secondNode.location.X
	leftestNode := secondNode

	lastDirAtLeftestNode := secondDir

	// fmt.Printf("second node %+v\n", secondNode)

	// travel whole loop once to find leftestNode
	for {
		n, d := traveler()

		if n.location.X < lowestX {
			leftestNode = n
			lowestX = n.location.X
			lastDirAtLeftestNode = d
		}

		if n.originalMarker == 'S' {
			break
		}
	}

	traveler = makeTravelFunc(startingNode)
	node, lastDir := traveler()

	// flow direction is how we determine what side of point is inside
	flowClockWise := determineFlow(leftestNode, lastDirAtLeftestNode)

	// fmt.Printf("flow clockwise %v\n", flowClockWise)
	// fmt.Printf("leftest node loc %+v\n", leftestNode.location)
	// fmt.Printf("last dir %s\n", dirToName(lastDirAtLeftestNode))

	foundCount := 0
	for {
		points := getInnerPointsForNode(node, lastDir, flowClockWise)
		for _, p := range points {
			foundCount += floodSearch(grid, p)
		}

		node, lastDir = traveler()
		if node == startingNode {
			break
		}
	}

	return foundCount
}

func determineFlow(leftestNode *Node, lastDir Dir) bool {
	nodeType := leftestNode.nodeType
	if nodeType == Start {
		nodeType = detectStartNodeType(leftestNode)
	}

	// fmt.Printf("node type %d\n", nodeType)
	// fmt.Printf("original marker %s\n", string(leftestNode.originalMarker))

	switch nodeType {
	case Vertical: // |
		return lastDir == North
	case NEBend: // L
		return lastDir == West
	case SEBend: // F
		return lastDir == North
	}

	log.Fatal("impossible nodetype at leftmost position")
	return false
}

func detectStartNodeType(startingNode *Node) NodeType {
	n := startingNode.north != nil
	e := startingNode.east != nil
	s := startingNode.south != nil
	w := startingNode.west != nil

	if n && s {
		return Vertical
	} else if e && w {
		return Horizontal
	} else if n && e {
		return NEBend
	} else if n && w {
		return NWBend
	} else if s && e {
		return SEBend
	} else {
		return SWBend
	}
}

func getInnerPointsForNode(node *Node, lastDirTraveled Dir, flowClockWise bool) []Point {
	points := make([]Point, 0, 2)

	// imagine your hands are on a steering wheel, thumbs sticking out
	// left hand thumb points clockwise, right hand points counter clockwise
	// your palm always faces towards the center of the steering wheel
	// by knowing the direction we're moving we know where the center is
	// left hand -> clockwise
	// right hand -> counter clockwise
	// direction of thumb determined by lastDirTraveled
	// direction of palm indicates inner side of point

	nodeType := node.nodeType
	if nodeType == Start {
		nodeType = detectStartNodeType(node)
	}

	northLoc := Point{node.location.X, node.location.Y - 1}
	eastLoc := Point{node.location.X + 1, node.location.Y}
	southLoc := Point{node.location.X, node.location.Y + 1}
	westLoc := Point{node.location.X - 1, node.location.Y}

	// northEastLoc := Point{node.location.X + 1, node.location.Y - 1}
	// northWestLoc := Point{node.location.X - 1, node.location.Y - 1}
	// southEastLoc := Point{node.location.X + 1, node.location.Y + 1}
	// southWestLoc := Point{node.location.X - 1, node.location.Y + 1}

	switch nodeType {
	case Horizontal: // -
		if flowClockWise {
			if lastDirTraveled == East {
				points = append(points, southLoc)
			} else {
				points = append(points, northLoc)
			}
		} else {
			if lastDirTraveled == East {
				points = append(points, northLoc)
			} else {
				points = append(points, southLoc)
			}
		}
	case Vertical: // |
		if flowClockWise {
			if lastDirTraveled == North {
				points = append(points, eastLoc)
			} else {
				points = append(points, westLoc)
			}
		} else {
			if lastDirTraveled == North {
				points = append(points, westLoc)
			} else {
				points = append(points, eastLoc)
			}
		}
	case NEBend: // L
		if flowClockWise {
			if lastDirTraveled == South {
				points = append(points, southLoc)
				points = append(points, westLoc)
				// points = append(points, southWestLoc)
			}
			// no points added if moving west
		} else {
			if lastDirTraveled == West {
				points = append(points, southLoc)
				points = append(points, westLoc)
				// points = append(points, southWestLoc)
			}
			// no points added if moving south
		}
	case NWBend: // J
		if flowClockWise {
			if lastDirTraveled == East {
				points = append(points, southLoc)
				points = append(points, eastLoc)
				// points = append(points, southEastLoc)
			}
		} else {
			if lastDirTraveled == South {
				points = append(points, southLoc)
				points = append(points, eastLoc)
				// points = append(points, southEastLoc)
			}
		}
	case SEBend: // F
		if flowClockWise {
			if lastDirTraveled == West {
				points = append(points, northLoc)
				points = append(points, westLoc)
			}
		} else {
			if lastDirTraveled == North {
				points = append(points, northLoc)
				points = append(points, westLoc)
			}
		}
	case SWBend: // 7
		if flowClockWise {
			if lastDirTraveled == North {
				points = append(points, northLoc)
				points = append(points, eastLoc)
			}
		} else {
			if lastDirTraveled == East {
				points = append(points, northLoc)
				points = append(points, eastLoc)
			}
		}
	}

	return points
}

func floodSearch(grid *Grid, startingPoint Point) int {
	if !grid.boundsCheck(startingPoint) {
		return 0
	}
	if grid.getPoint(startingPoint) != '.' {
		return 0
	}

	/*Flood-fill (node):
	1. If node is not Inside return.
	2. Set the node
	3. Perform Flood-fill one step to the south of node.
	4. Perform Flood-fill one step to the north of node
	5. Perform Flood-fill one step to the west of node
	6. Perform Flood-fill one step to the east of node
	7. Return.
	*/

	found := 0
	// grid.setPoint(startingPoint, 'I')
	// found += 1

	first := true
	stack := make([]Point, 0)
	stack = append(stack, startingPoint)
	for len(stack) > 0 {
		// pop point from top of stack
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !grid.boundsCheck(p) {
			continue
		}

		if grid.getPoint(p) != '.' {
			continue
		}

		if first {
			grid.setPoint(p, 'A')
			first = false
		} else {
			grid.setPoint(p, 'I')
		}
		found += 1

		nLoc := Point{p.X, p.Y - 1}
		eLoc := Point{p.X + 1, p.Y}
		sLoc := Point{p.X, p.Y + 1}
		wLoc := Point{p.X - 1, p.Y}

		stack = append(stack, nLoc)
		stack = append(stack, eLoc)
		stack = append(stack, sLoc)
		stack = append(stack, wLoc)
	}

	return found
}
