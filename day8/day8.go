package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Direction = int

const (
	L Direction = iota
	R
)

type NodeCategory = int

const (
	Beginning NodeCategory = iota
	Normal
	End
)

type Node struct {
	label    string
	category NodeCategory
	left     *Node
	right    *Node
}

type NodeGraph struct {
	root Node
}

func Day8P1() {
	fmt.Println("day 8 p 1")
	// file_name := "example_inputA.txt" //should require 2 steps
	// file_name := "example_inputB.txt" //should require 6 steps
	file_name := "input.txt"

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//first line is instructions
	scanner.Scan()
	firstLine := scanner.Text()
	instructions := parseInstructionsP1(firstLine)

	// empty line
	scanner.Scan()
	if len(scanner.Text()) != 0 {
		log.Fatal("expected second line to be empty")
	}

	//all following lines are the nodes
	graph := parseNodesP1(scanner)

	//follow instructions along graph starting from root node
	count := 0
	instructionCount := len(instructions)
	instructionPos := 0
	node := graph.root
	for node.category != End {
		ins := instructions[instructionPos]

		switch ins {
		case L:
			node = *node.left
		case R:
			node = *node.right
		}

		count += 1
		instructionPos = (instructionPos + 1) % instructionCount
	}

	fmt.Printf("step count: %d\n", count)
}

func parseInstructionsP1(firstLine string) []Direction {
	dirs := make([]Direction, 0)

	for _, r := range firstLine {
		switch r {
		case 'L':
			dirs = append(dirs, L)
		case 'R':
			dirs = append(dirs, R)
		}
	}

	return dirs
}

func parseNodesP1(scanner *bufio.Scanner) NodeGraph {
	graph := NodeGraph{}

	nodeMap := make(map[string]*Node)
	getNode := func(label string) *Node {
		node, ok := nodeMap[label]
		if !ok {
			node = &Node{
				label:    label,
				category: Normal,
				left:     nil,
				right:    nil,
			}
			nodeMap[label] = node
		}
		return node
	}

	for scanner.Scan() {
		line := scanner.Text()

		splitA := strings.Split(line, " = ")
		label := splitA[0]

		splitB := strings.Split(splitA[1], ", ")
		leftDest := strings.Trim(splitB[0], "()")
		rightDest := strings.Trim(splitB[1], "()")

		category := Normal
		if label == "AAA" {
			category = Beginning
		}
		if label == "ZZZ" {
			category = End
		}

		leftNode := getNode(leftDest)
		rightNode := getNode(rightDest)

		node := getNode(label)
		node.category = category
		node.left = leftNode
		node.right = rightNode

		if category == Beginning {
			graph.root = *node
		}
	}

	// testing
	// for key := range nodeMap {
	// 	n := *nodeMap[key]
	// 	fmt.Printf("%s: %+v left: %s, right: %s\n", key, n, n.left.label, n.right.label)
	// }

	return graph
}
