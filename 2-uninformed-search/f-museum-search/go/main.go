package main

import (
	"fmt"
)

// succesor devuelve los sucesores de un nodo dado
func succesor(n []string) [][]string {
	switch n[0] {
	case "A":
		return [][]string{{"B"}, {"C"}, {"D"}}
	case "B":
		return [][]string{{"C"}, {"E"}}
	case "C":
		return [][]string{{"E"}, {"F"}, {"G"}}
	case "D":
		return [][]string{{"C"}, {"F"}}
	case "E":
		return [][]string{{"G"}}
	case "F":
		return [][]string{{"G"}}
	default:
		return nil
	}
}

// print only node names
func printNodes(path [][]string) string  {
	nodes := ""
	for i := range path {
			nodes += path[i][0] + " "
	}
	return nodes
}

// reverse invierte un slice de strings
func reverse(path []string) []string {
	reversed := make([]string, len(path))
	for i := range path {
		reversed[i] = path[len(path)-1-i]
	}
	return reversed
}

// museum realiza la búsqueda
func museum(begin, end string) {
	numsol := 0
	novisited := [][]string{{begin}}

	for len(novisited) > 0 {
		current := novisited[0]
		novisited = novisited[1:]

		fmt.Print("Current: ", current[0])

		if current[0] == end {
			numsol++
			fmt.Println(" | ** SOLUTION PATH: ", reverse(current)," **")
			continue
		}

		temp := succesor(current)
		fmt.Print(" | Succ: ", printNodes(temp))

		if temp != nil {
			for i := range temp {
				temp[i] = append(temp[i], current...)
			}
			novisited = append(novisited, temp...)
			fmt.Println("| NewList: ",printNodes(novisited))
		}
	}

	fmt.Printf("TOTAL: %d\n", numsol)
}

func main() {
	museum("A", "G")
}
