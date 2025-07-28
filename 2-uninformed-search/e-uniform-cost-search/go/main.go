package main

import (
	"fmt"
	"sort"
)

func successors(node []int) [][]int {
	name := node[0]
	cost := node[1]

	switch name {
	case 1:
		return [][]int{
			{2, cost + 5},
			{3, cost + 6},
		}
	case 2:
		return [][]int{
			{1, cost + 5},
			{3, cost + 6},
			{4, cost + 3},
			{5, cost + 5},
		}
	case 3:
		return [][]int{
			{1, cost + 6},
			{2, cost + 6},
			{5, cost + 2}, 
		}
	case 4:
		return [][]int{
			{2, cost + 3},
			{5, cost + 3},
			{6, cost + 4},
		}
	case 5:
		return [][]int{
			{2, cost + 5},
			{3, cost + 2},
			{4, cost + 3},
			{6, cost + 1},
		}
	case 6:
		return [][]int{
			{4, cost + 4},
			{5, cost + 1},
		}
	}
	return [][]int{}
}

func uniformCost(begin, end int) {
	list := [][]int{{begin, 0}}

	for len(list) > 0 {
		current := list[0]
		list = list[1:]
		fmt.Println("Current Node: ", current)

		if current[0] == end {
			fmt.Println("SOLUTION")
			return
		}

		tmp := successors(current)
		fmt.Println("Successors:", tmp)

		if len(tmp) > 0 {
			list = append(list, tmp...)
			sort.Slice(list, func(i, j int) bool {
				return list[i][1] < list[j][1]
			})
			fmt.Println("New List: ", list)
		}
	}

	fmt.Println("NO-SOLUTION")
}

func main() {
	uniformCost(1, 6)
}
