package main

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
	label    string // just in case for debugging
	category NodeCategory
	left     *Node
	right    *Node
}

func main() {
	fmt.Println("day 8 p 2")
	// file_name := "example_inputA.txt" //should require 2 steps
	// file_name := "example_inputB.txt" //should require 5 steps
	// file_name := "example_inputC.txt" //should require 6 steps
	file_name := "input.txt" // probable answer 8,811,050,362,409

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//first line is instructions
	scanner.Scan()
	firstLine := scanner.Text()
	instructions := parseInstructionsP2(firstLine)

	// empty line
	scanner.Scan()
	if len(scanner.Text()) != 0 {
		log.Fatal("expected second line to be empty")
	}

	//all following lines are the nodes
	currentNodes := parseNodesP2(scanner)
	nodeCount := len(currentNodes)
	// fmt.Printf("starting node count %d\n", nodeCount)

	//follow instructions along graph starting from root node
	var count uint64 = 0
	instructionCount := len(instructions)
	instructionPos := 0

	endReachedCounts := make([]int, nodeCount)

	isCountingCycle := make([]bool, nodeCount)
	hasCountedCycle := make([]bool, nodeCount)
	endCycleCount := make([]uint64, nodeCount)
	for i := range currentNodes {
		isCountingCycle[i] = false
		hasCountedCycle[i] = false
	}

	// fmt.Println("starting path tracing")
	for {
		// allDone := true
		allHasCounted := true

		ins := instructions[instructionPos]
		for i, node := range currentNodes {
			switch ins {
			case L:
				currentNodes[i] = node.left
			case R:
				currentNodes[i] = node.right
			}

			if isCountingCycle[i] {
				endCycleCount[i] += 1
			}
			if currentNodes[i].category != End {
				// allDone = false
			} else {
				endReachedCounts[i] += 1

				if !hasCountedCycle[i] {
					hasCountedCycle[i] = true
					endCycleCount[i] = count + 1

					// fmt.Printf("%d cycle count %d\n", i, endCycleCount[i])
					// fmt.Printf("\tend reached count %d\n", endReachedCounts[i])
					// fmt.Printf("\tstep count %d\n", count+1)
					// fmt.Printf("\tstep count - endCycleCount %d\n", (count+1)-uint64(endCycleCount[i]))
				}

				// if !isCountingCycle[i] {
				// 	if !hasCountedCycle[i] {
				// 		isCountingCycle[i] = true
				// 		endCycleCount[i] = 0
				// 		fmt.Printf("%d: starting cycle counting at steps %d\n", i, count + 1)
				// 	}
				// } else {
				// 	isCountingCycle[i] = false
				// 	hasCountedCycle[i] = true
				// 	fmt.Printf("%d cycle count %d\n", i, endCycleCount[i])
				// 	fmt.Printf("\tend reached count %d\n", endReachedCounts[i])
				// 	fmt.Printf("\tstep count %d\n", count)
				// 	fmt.Printf("\tstep count - endCycleCount %d\n", count-uint64(endCycleCount[i]))
				// }

			}
			if !hasCountedCycle[i] {
				allHasCounted = false
			}
		}

		count += 1
		instructionPos = (instructionPos + 1) % instructionCount

		// if allDone || allHasCounted {
		if allHasCounted {
			break
		}
	}

	// fmt.Printf("using cycle counts to find least common multiple\n")
	lcmResult := lowestCommonMultiple(endCycleCount)

	fmt.Printf("total steps needed: %d\n", lcmResult)
}

func parseInstructionsP2(firstLine string) []Direction {
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

func parseNodesP2(scanner *bufio.Scanner) []*Node {
	startingNodes := make([]*Node, 0)

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
		if label[len(label)-1] == 'A' {
			category = Beginning
		}
		if label[len(label)-1] == 'Z' {
			category = End
		}

		leftNode := getNode(leftDest)
		rightNode := getNode(rightDest)

		node := getNode(label)
		node.category = category
		node.left = leftNode
		node.right = rightNode

		if category == Beginning {
			startingNodes = append(startingNodes, node)
		}
	}

	// testing
	// for key := range nodeMap {
	// 	n := *nodeMap[key]
	// 	fmt.Printf("%s: %+v left: %s, right: %s\n", key, n, n.left.label, n.right.label)
	// }

	return startingNodes
}

func lowestCommonMultiple(nums []uint64) uint64 {

	// all primes below 1000
	primes := []uint64{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199,
		211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293,
		307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397,
		401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499,
		503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599,
		601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691,
		701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797,
		809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887,
		907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997,
	}

	factors := make([]map[uint64]uint64, len(nums))
	for i := 0; i < len(factors); i++ {
		factors[i] = make(map[uint64]uint64)
	}
	for i, n := range nums {
		//find factors of n
		f := factors[i]
		for n > 1 {
			noFactorFound := true
			for _, p := range primes {
				if n % p == 0 {
					_, ok := f[p]
					if !ok {
						f[p] = 0
					}
					f[p] += 1
					n /= p
					noFactorFound = false
					break
				}
			}
			if noFactorFound {
				log.Fatalf("Failed to find factor for %d\n", n)
			}
		}
	}

	// lcm factors are the factors of which have the highest count
	// https://www.calculatorsoup.com/calculators/math/lcm.php
	// LCM by prime factorization
	lcmFactors := make(map[uint64]uint64)
	for _, f := range factors {
		for prime := range f {
			_, ok := lcmFactors[prime]
			if !ok {
				lcmFactors[prime]  = 0
			}
			
			if f[prime] > lcmFactors[prime] {
				lcmFactors[prime] = f[prime]
			}
		}
	}

	var lcmResult uint64 = 1 
	for prime := range lcmFactors {
		for i := uint64(0); i < lcmFactors[prime]; i++ {
			lcmResult *= uint64(prime)
		}
	}

	return lcmResult
}
