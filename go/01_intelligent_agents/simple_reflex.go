package main

import (
	"fmt"
	"time"
)

// Agente reactivo simple
func reflexAgent(location, state string) string {
	if state == "DIRTY" {
		return "CLEAN"
	} else if location == "A" {
		return "RIGHT"
	} else if location == "B" {
		return "LEFT"
	}
	return ""
}

// Lógica principal
func test(states []string) {
	for {
		location := states[0]
		var state string
		if location == "A" {
			state = states[1]
		} else {
			state = states[2]
		}

		action := reflexAgent(location, state)
		fmt.Printf("Location: %s | Action: %s\n", location, action)

		if action == "CLEAN" {
			if location == "A" {
				states[1] = "CLEAN"
			} else if location == "B" {
				states[2] = "CLEAN"
			}
		} else if action == "RIGHT" {
			states[0] = "B"
		} else if action == "LEFT" {
			states[0] = "A"
		}

		time.Sleep(3 * time.Second)
	}
}

func main() {
	states := []string{"A", "DIRTY", "DIRTY"} // [location, A_state, B_state]
	test(states)
}
