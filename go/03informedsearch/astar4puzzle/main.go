package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var id = 1

func heuristic(start, end string) int {
	tilesOut := 0
	for i := 0; i < len(start); i++ {
		if start[i] != end[i] {
			tilesOut++
		}
	}
	return tilesOut
}

func inc() int {
	id++
	return id
}

type Node struct {
	state  string
	cost   int
	id     int
	level  int
}

func successors(n Node, end string) []Node {
	var suc []Node
	for i := 0; i < len(n.state)-1; i++ {
		level := n.level + 1
		chars := []rune(n.state)
		chars[i], chars[i+1] = chars[i+1], chars[i]
		child := string(chars)
		suc = append(suc, Node{
			state: child,
			cost:  heuristic(child, end) + level,
			id:    inc(),
			level: level,
		})
	}
	return suc
}

func aStar(start, end string) string {
	dot := "graph{"
	list := []Node{{
		state: start,
		cost:  heuristic(start, end),
		id:    id,
		level: 0,
	}}
	dot += fmt.Sprintf("\n%d [label=\"%s\"];\n", list[0].id, list[0].state)

	cont := 0
	for len(list) > 0 {
		current := list[0]
		list = list[1:]

		if current.state == end {
			dot += "}"
			return dot
		}

		temp := successors(current, end)
		for _, val := range temp {
			dot += fmt.Sprintf("%d [label=\"%s\"]; ", val.id, val.state)
			dot += fmt.Sprintf("%d--%d [label=\"%d+%d\"];\n", current.id, val.id, val.cost-val.level, val.level)
		}

		list = append(list, temp...)
		sort.Slice(list, func(i, j int) bool { return list[i].cost < list[j].cost })

		cont++
		if cont > 100 {
			fmt.Println("The search is looped!")
			dot += "}"
			return dot
		}
	}
	dot += "}"
	return dot
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter start text and end text separated by a space: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		input = "halo hola"
	}
	parts := strings.Split(input, " ")
	if len(parts) < 2 {
		fmt.Println("Invalid entry.")
		return
	}
	start := parts[0]
	end := parts[1]

	fmt.Println(aStar(start, end))
}
