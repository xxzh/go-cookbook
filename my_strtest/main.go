package main

import (
	"bytes"
	"fmt"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type State int

const (
	StateA State = iota
	StateB
	StateC
	StateD
	State_other
)

type Pos struct {
	i int   // the position index
	s State // the current state
}

// On returns the next state and a boolean value indicating if the target case is found.
func (s State) On(a rune) (State, bool) {
	switch a {
	case 'A':
		if s == StateB {
			return State_other, true
		}
		return StateA, false
	case 'B':
		if s == StateA {
			return State_other, true
		}
		return StateB, false
	case 'C':
		if s == StateD {
			return State_other, true
		}
		return StateC, false
	case 'D':
		if s == StateC {
			return State_other, true
		}
		return StateD, false
	default:
		return State_other, false
	}
}

// iterate the string once and removes one target case if found.
func iterate(S string, removed []bool) int {
	var pos = Pos{-1, State_other}
	var count int
	for i, c := range S {
		if removed[i] {
			// already removed
			continue
		}
		next_stat, found := pos.s.On(c)
		if !found {
			pos.i = i
			pos.s = next_stat
			continue
		}
		// found the case
		removed[i] = true
		if pos.i >= 0 {
			removed[pos.i] = true
		}
		count += 2
		// reset and start again
		pos.s = State_other
		pos.i = -1
	}
	return count
}

func Solution(S string) string {
	// Implement your solution here
	removed := make([]bool, len(S))
	for {
		count := iterate(S, removed)
		if count == 0 {
			// no more target case found
			break
		}
	}
	var buf bytes.Buffer
	for i, c := range removed {
		if !c {
			buf.WriteByte(S[i])
		}
	}
	return buf.String()
}

func main() {
	// r := Solution("CBACD")
	// r := Solution("CABABD")
	r := Solution("DAABABCDADDDA")
	fmt.Println(r)
}
