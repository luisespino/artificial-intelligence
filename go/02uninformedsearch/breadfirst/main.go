package main

import (
	"fmt"
)

// returns the successors of a given node in a graph.
func successors(n int) []int {
	switch n {
	case 1:
		return []int{2, 4, 5}
	case 2:
		return []int{1, 3, 4, 5, 6}
	case 3:
		return []int{2, 5, 6}
	case 4:
		return []int{1, 2, 5, 7, 8}
	case 5:
		return []int{1, 2, 3, 4, 6, 7, 8, 9}
	case 6:
		return []int{2, 3, 5, 8, 9}
	case 7:
		return []int{4, 5, 8}
	case 8:
		return []int{4, 5, 6, 7, 9}
	case 9:
		return []int{5, 6, 8}
	default:
		return nil
	}
}

// Breadth-First Search (BFS) algorithm to find a path from the begin node to the end node.
func breadFirstSearch(begin, end int) {
	list := []int{begin}

	for len(list) > 0 {
		current := list[0]
		list = list[1:]

		fmt.Println(current)

		if current == end {
			fmt.Println("SOLUTION")
			return
		}

		tmp := successors(current)
		fmt.Println(tmp)

		if tmp != nil {
			list = append(list, tmp...)
			fmt.Println(list)
		}
	}
	fmt.Println("NO-SOLUTION")
}

// Main function to execute the Breadth-First Search algorithm.
func main() {
	fmt.Println("Breah-First Search:")
	breadFirstSearch(1, 9)
}
