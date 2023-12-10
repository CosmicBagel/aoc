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
	scanner.Scan()
	fmt.Println(scanner.Text())
}
